// Work with engine.AstEngine: walk an AST, dispatch known nodes, and let
// unknown ones be skipped instead of aborting the whole traversal.

package main

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/pt-main/lc/v2"
	enginepkg "github.com/pt-main/lc/v2/engine"
	"github.com/pt-main/lc/v2/engine/core"
	"github.com/pt-main/lc/v2/parsing/stringParsing"
	"github.com/pt-main/lc/v2/parsing/stringParsing/parser3"
	"github.com/pt-main/lc/v2/public"
)

func main() {
	lexer, lerr := stringParsing.NewLexer([]stringParsing.LexerRule{
		{Type: "IDENT", Pattern: `[a-z]+`},
		{Type: "WS", Pattern: `\s+`},
	}, nil)
	if lerr != nil {
		fmt.Println(lerr.Error())
		return
	}

	// root is the entry node Work() starts from; the "word" nodes are nested
	// inside it and get visited by the depth-first walk.
	grammar := parser3.Grammar{
		"root": {Name: "root", Expr: parser3.NodeExpr{NodeType: "root", Expr: parser3.SequenceExpr{Exprs: []parser3.Expr{
			parser3.RepeatExpr{Expr: parser3.NamedExpr{RuleName: "word"}},
		}}}},
		"word": {Name: "word", Expr: parser3.NodeExpr{NodeType: "word", Expr: parser3.TokenExpr{TokenType: "IDENT"}}},
	}

	adapter := &parser3.Adapter{
		Parser: parser3.NewParser(lexer, grammar, "root", []string{"WS"}),
	}

	engine := lc.NewAstEngine(
		public.StringResType,
		[]string{"main"},
		true, // add default events
		adapter,
		context.Background(),
		true, // canNodeBeUnknown: unknown nodes are skipped, not fatal
		true, // canMainNodeBeUnknown
	)

	engine.NewCommandFull("root", func(enginepkg.AstEngineInterface, *stringParsing.ParsedNode) core.ErrorInterface {
		return nil
	}, "entry point of the walk")

	engine.NewCommandFull("word", func(e enginepkg.AstEngineInterface, node *stringParsing.ParsedNode) core.ErrorInterface {
		return e.GetUep().Generator.AddString("word: "+node.Raw, "main")
	}, "collect every known word")

	if err := engine.Process("alpha beta gamma"); err != nil {
		panic(err)
	}

	uep := engine.GetUep()
	out, err := core.GetStringRes(uep.Generator, "\n")
	if err != nil {
		panic(err)
	}
	fmt.Println("visited nodes:")
	fmt.Println(strings.TrimSpace(out))

	names := make([]string, 0, len(engine.GetCommands()))
	for name := range engine.GetCommands() {
		names = append(names, name)
	}
	sort.Strings(names)
	fmt.Println("\nregistered commands:", strings.Join(names, ", "))

	fmt.Println("\nunknown node policy (canNodeBeUnknown = true):")
	fmt.Println("  HasCommand(\"nope\") ->", engine.HasCommand(false, "nope"))
}
