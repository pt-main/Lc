// Work with engine.StringEngine: register commands, read arguments from the
// node metadata the default events put there, and collect the output.

package main

import (
	"fmt"
	"strings"

	"github.com/pt-main/lc/v2"
	enginepkg "github.com/pt-main/lc/v2/engine"
	"github.com/pt-main/lc/v2/engine/core"
	"github.com/pt-main/lc/v2/parsing/stringParsing"
	"github.com/pt-main/lc/v2/public"
)

func main() {
	parser := &stringParsing.Parser2{}

	engine, err := lc.NewEngineBuilder(public.StringEngineType, public.StringResType).
		WithPipeline([]string{"main"}).
		WithStringParser(parser).
		WithDefaultEvents(true).
		Build()
	if err != nil {
		panic(err)
	}

	err = engine.NewCommandString("say", func(se enginepkg.StringEngineInterface, node *stringParsing.ParsedNode) core.ErrorInterface {
		args, _ := node.Metadata["args"].(string)
		return se.GetUep().Generator.AddString("say: "+strings.TrimSpace(args), "main")
	}, "echo the arguments back")
	if err != nil {
		panic(err)
	}

	err = engine.NewCommandString("count", func(se enginepkg.StringEngineInterface, node *stringParsing.ParsedNode) core.ErrorInterface {
		args, _ := node.Metadata["args"].(string)
		return se.GetUep().Generator.AddString(fmt.Sprintf("count: %d word(s)", len(strings.Fields(args))), "main")
	}, "count the words in the arguments")
	if err != nil {
		panic(err)
	}

	input := strings.Join([]string{
		"say hello from the string engine",
		"count one two three",
	}, "\n")

	if err := engine.ProcessString(input); err != nil {
		panic(err)
	}

	uep, _ := engine.GetUEP()
	out, err := core.GetStringRes(uep.Generator, "\n")
	if err != nil {
		panic(err)
	}
	fmt.Println(out)

	if err := engine.End(); err != nil {
		panic(err)
	}
	if err := engine.CheckEnded(); err == nil {
		fmt.Println("engine lifecycle: still running")
		return
	}
	fmt.Println("engine lifecycle: ended, further calls are refused")
}
