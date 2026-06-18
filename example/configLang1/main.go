package main

import (
	"context"
	"log"

	"github.com/dlclark/regexp2"
	"github.com/pt-main/lc"
	engines "github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing/stringParsing"
)

func newLexer() *stringParsing.Lexer {
	rules := []stringParsing.LexerRule{
		{Type: "COMMENT", Pattern: regexp2.MustCompile(`(?:#|//)[^\n]*`, 0)},
		{Type: "WHITESPACE", Pattern: regexp2.MustCompile(`\s+`, 0)},
		{Type: "LBRACE", Pattern: regexp2.MustCompile(`\{`, 0)},
		{Type: "RBRACE", Pattern: regexp2.MustCompile(`\}`, 0)},
		{Type: "ASSIGN", Pattern: regexp2.MustCompile(`=`, 0)},
		{Type: "STRING", Pattern: regexp2.MustCompile(`"([^"]*)"`, 0)},
		{Type: "NUMBER", Pattern: regexp2.MustCompile(`-?\d+(\.\d+)?`, 0)},
		{Type: "BOOL", Pattern: regexp2.MustCompile(`true|false`, 0)},
		{Type: "IDENT", Pattern: regexp2.MustCompile(`[a-zA-Z_][a-zA-Z0-9_]*`, 0)},
	}
	return stringParsing.NewLexer(rules)
}

type DataBase struct {
	Scope core.ScopeType
}

func process(config string) (*DataBase, error) {

	parser := &configParser{}

	engine, err := lc.NewEngineBuilder(lc.StringEngineType).
		WithPipeline([]string{"main"}).
		WithStringParser(parser).
		WithDefaultEvents(true).
		WithContext(context.Background()).
		Build()
	if err != nil {
		log.Fatal(err)
	}

	engine.GetUEP().Scope["DB"] = &DataBase{Scope: make(core.ScopeType)}

	err = engine.NewCommandString("assign", func(se *engines.StringEngine, node stringParsing.ParsedNode) error {
		key, _ := node.Metadata["key"].(string)
		value := node.Metadata["value"]
		se.UEP.Scope["DB"].(*DataBase).Scope[key] = value
		return nil
	}, "save value to scope")
	if err != nil {
		return nil, err
	}

	if err := engine.ProcessString(config); err != nil {
		return nil, err
	}

	db := engine.GetUEP().Scope["DB"].(*DataBase)
	return db, nil
}
