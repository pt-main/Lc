package main

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pt-main/lc"
	enginepkg "github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing/stringParsing"
	"github.com/pt-main/lc/public"
	"github.com/pt-main/lc/tooling/debugging/extensiblePlugin"
	"github.com/pt-main/lc/tooling/debugging/profiler"
)

func Process(config string) (string, error) {
	grammar := []stringParsing.GrammarRule{
		{
			Type:    "section",
			Pattern: regexp.MustCompile(`^\[(?P<name>[^\]]+)\]$`),
		},
		{
			Type:    "keyval",
			Pattern: regexp.MustCompile(`^(?P<key>[^=]+?)\s*=\s*(?P<value>.*)$`),
		},
		{
			Type:    "comment",
			Pattern: regexp.MustCompile(`^\s*#[^\n]*$`),
		},
		{
			Type:    "unknown_token",
			Pattern: regexp.MustCompile(`.*`),
		},
	}
	cfgp := stringParsing.Parser1Config{
		UseLineContinuation: true,
		SkipEmptyLines:      true,
		UseBracketBalance:   false,
		TrimBlocksSpace:     true,
	}
	parser := stringParsing.NewParser1(grammar, cfgp)
	engine, err := lc.NewEngineBuilder(public.StringEngineType, public.StringResType).
		WithPipeline([]string{"main"}).
		WithStringParser(parser).
		WithDefaultEvents(true).
		WithContext(context.Background()).
		Build()
	if err != nil {
		return "", err
	}
	err = engine.Plugins.AddPlugin(extensiblePlugin.New(engine))
	if err != nil {
		return "", err
	}
	err = engine.Plugins.AddPlugin(profiler.New())
	if err != nil {
		return "", err
	}
	uep, _ := engine.GetUEP()
	uep.Scope[public.StringEngineScopeCanBeUnknown] = false
	uep.Scope["config"] = make(map[string]map[string]interface{})
	uep.Scope["current_section"] = "global"
	engine.NewCommandString("comment", func(se enginepkg.StringEngineInterface, node *stringParsing.ParsedNode) core.ErrorInterface {
		return nil
	}, "")
	engine.NewCommandString("section", func(se enginepkg.StringEngineInterface, node *stringParsing.ParsedNode) core.ErrorInterface {
		name := node.Metadata["name"].(string)
		if strings.HasPrefix(name, "!") {
			switch name[1:] {
			case "exit":
				return core.ErrExit
			default:
				return core.Err("PROCESS", "Invalid command")
			}
		}
		se.GetUep().Scope["current_section"] = name
		cfg := se.GetUep().Scope["config"].(map[string]map[string]interface{})
		if _, ok := cfg[name]; !ok {
			cfg[name] = make(map[string]interface{})
		}
		return nil
	}, "Set current section")
	engine.NewCommandString("keyval", func(se enginepkg.StringEngineInterface, node *stringParsing.ParsedNode) core.ErrorInterface {
		key := strings.TrimSpace(node.Metadata["key"].(string))
		rawVal := strings.TrimSpace(node.Metadata["value"].(string))
		val := parseValue(rawVal)
		sectionName, _ := se.GetUep().Scope["current_section"].(string)
		cfg := se.GetUep().Scope["config"].(map[string]map[string]interface{})
		if _, ok := cfg[sectionName]; !ok {
			cfg[sectionName] = make(map[string]interface{})
		}
		cfg[sectionName][key] = val
		return nil
	}, "Store key-value pair")
	if err := engine.ProcessString(config); err != nil {
		return "", err
	}
	cfg := uep.Scope["config"].(map[string]map[string]interface{})
	jsonData, err := json.MarshalIndent(cfg, "", "  ")
	report, err := engine.Plugins.CallPluginMethod("profiler", "report")
	fmt.Println(report, err)
	return string(jsonData), err
}

func parseValue(raw string) interface{} {
	raw = strings.TrimSpace(raw)
	if raw == "true" {
		return true
	}
	if raw == "false" {
		return false
	}
	if strings.HasPrefix(raw, "[") && strings.HasSuffix(raw, "]") {
		inner := strings.TrimSpace(raw[1 : len(raw)-1])
		if inner == "" {
			return []interface{}{}
		}
		parts := strings.Split(inner, ",")
		result := make([]interface{}, 0, len(parts))
		for _, p := range parts {
			result = append(result, parseValue(strings.TrimSpace(p)))
		}
		return result
	}
	if i, err := strconv.Atoi(raw); err == nil {
		return i
	}
	if f, err := strconv.ParseFloat(raw, 64); err == nil {
		return f
	}
	if len(raw) >= 2 {
		if (raw[0] == '"' && raw[len(raw)-1] == '"') || (raw[0] == '\'' && raw[len(raw)-1] == '\'') {
			return raw[1 : len(raw)-1]
		}
	}
	return raw
}

func main() {
	fmt.Println("Lc version -", lc.Version)

	fmt.Println(Process(`
# Global params
app_name = MyApp
version = 1.2.3

[server]
host = localhost
port = 8080
debug = true
timeout = 30.5

[database]
url = postgres://user:pass@localhost/db
pool_size = 10
features = [caching, logging, metrics]
`))
}

/*
Lc version - 1.5.1
Profiler report (0.00 sec total):

  String commands:
  String calls: 13 (total time: 718.071µs)
    keyval: count=9, total=484.819µs, avg=53.868µs, min=45.808µs, max=68.911µs
    section: count=2, total=92.308µs, avg=46.154µs, min=45.346µs, max=46.962µs

{
  "database": {
    "features": [
      "caching",
      "logging",
      "metrics"
    ],
    "pool_size": 10,
    "url": "postgres://user:pass@localhost/db"
  },
  "global": {
    "app_name": "MyApp",
    "version": "1.2.3"
  },
  "server": {
    "debug": true,
    "host": "localhost",
    "port": 8080,
    "timeout": 30.5
  }
} <nil>
*/
