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
	"github.com/pt-main/lc/parsing/stringParsing"
	"github.com/pt-main/lc/public"
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
			Type:    "skip",
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
		WithColors().
		Build()
	if err != nil {
		return "", err
	}
	engine.GetUEP().Scope["config"] = make(map[string]map[string]interface{})
	engine.GetUEP().Scope["current_section"] = "global"
	_ = engine.NewCommandString("section", func(se *enginepkg.StringEngine, node *stringParsing.ParsedNode) error {
		name := node.Metadata["name"].(string)
		se.UEP.Scope["current_section"] = name
		cfg := se.UEP.Scope["config"].(map[string]map[string]interface{})
		if _, ok := cfg[name]; !ok {
			cfg[name] = make(map[string]interface{})
		}
		return nil
	}, "Set current section")
	_ = engine.NewCommandString("keyval", func(se *enginepkg.StringEngine, node *stringParsing.ParsedNode) error {
		key := strings.TrimSpace(node.Metadata["key"].(string))
		rawVal := strings.TrimSpace(node.Metadata["value"].(string))
		val := parseValue(rawVal)
		sectionName, _ := se.UEP.Scope["current_section"].(string)
		cfg := se.UEP.Scope["config"].(map[string]map[string]interface{})
		if _, ok := cfg[sectionName]; !ok {
			cfg[sectionName] = make(map[string]interface{})
		}
		cfg[sectionName][key] = val
		return nil
	}, "Store key-value pair")
	if err := engine.ProcessString(config); err != nil {
		return "", err
	}
	cfg := engine.GetUEP().Scope["config"].(map[string]map[string]interface{})
	jsonData, err := json.MarshalIndent(cfg, "", "  ")
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
features = [caching, logging, metrics]`))
}
