package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/pt-main/lc"
	enginepkg "github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing/stringParsing"
	"github.com/pt-main/lc/public"
)

func main() {
	parser := &stringParsing.Parser2{} // parsing format: 'command arg1, arg2...'

	engine, err := lc.NewEngineBuilder(public.StringEngineType, public.StringResType).
		WithPipeline([]string{"main"}).
		WithStringParser(parser).
		WithDefaultEvents(true).
		Build()
	if err != nil {
		panic(err)
	}

	err = engine.NewCommandString("log", func(se enginepkg.StringEngineInterface, node *stringParsing.ParsedNode) core.ErrorInterface {
		args, _ := node.Metadata["args"].(string)
		return se.GetUep().Generator.AddString(fmt.Sprintf("Log [%v]: %v",
			time.Now().Format(time.Stamp), args), "main")
	}, "append log with timestamp")
	if err != nil {
		panic(err)
	}

	err = engine.ProcessString(strings.Join([]string{
		"log service_start",
		"log service_ready",
	}, "\n"))
	if err != nil {
		panic(err)
	}

	uep, _ := engine.GetUEP()
	out, err := core.GetStringRes(uep.Generator, "\n")
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}
