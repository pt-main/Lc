// A complete interpreter for a small language, written on the Lc StringEngine.
//
// A lexer and a parser3 grammar build the syntax tree, one engine command per
// statement form runs it, the engine generator collects the output and the
// profiler plugin reports where the time went.
//
// Language:
//
//	literals    42  "text"  true  false  [1, 2, 3]
//	variables   x := 1     declare
//	            x = 2      assign
//	functions   fn add(a, b) { return a + b }
//	            fn (x) { x * 2 }      anonymous
//	control     if c { } else { }    while c { }    for x in list { }
//	                break    continue    return value
//	builtins    print(x)  len(x)  str(x)  int(x)  map(list, fn)
//	            filter(list, fn)  reduce(list, fn, start)
//
// Functions are closures: a body sees the variables that were visible where
// the function was declared, and a named function can call itself.

package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dlclark/regexp2"
	"github.com/pt-main/lc"
	enginepkg "github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing/stringParsing"
	"github.com/pt-main/lc/parsing/stringParsing/parser3"
	"github.com/pt-main/lc/public"
	lce "github.com/pt-main/lc/public/errors"
	"github.com/pt-main/lc/tooling/astools"
	"github.com/pt-main/lc/tooling/debugging/extensiblePlugin"
	"github.com/pt-main/lc/tooling/debugging/profiler"
)

func main() {
	fmt.Println("Lc version -", lc.Version)

	source, err := readSource()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	interp, err := newInterpreter()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer interp.Engine.End()

	if perr := interp.Engine.ProcessString(source); perr != nil {
		reportError(perr)
		os.Exit(1)
	}

	uep, uerr := interp.Engine.GetUEP()
	if uerr != nil {
		fmt.Fprintln(os.Stderr, uerr)
		os.Exit(1)
	}
	out, gerr := core.GetStringRes(uep.Generator, "\n")
	if gerr != nil {
		fmt.Fprintln(os.Stderr, gerr.Format())
		os.Exit(1)
	}
	if out != "" {
		fmt.Println(out)
	}

	report, perr := interp.Engine.Plugins.CallPluginMethod(profiler.Name, "report")
	if perr == nil {
		fmt.Println(report)
	}
}

// readSource takes the program from the command line, so a one line program
// does not have to be piped in.
func readSource() (string, error) {
	if len(os.Args) > 1 {
		return strings.Join(os.Args[1:], " "), nil
	}
	data, err := os.ReadFile(os.Stdin.Name())
	if err != nil {
		return "", fmt.Errorf("read the program from stdin: %w", err)
	}
	return string(data), nil
}

// reportError prints the innermost cause, which is the one written by the
// interpreter, instead of the whole chain of engine wrappers around it.
func reportError(err core.ErrorInterface) {
	if perr, ok := parser3.AsParseError(err); ok {
		fmt.Fprintln(os.Stderr, "parse error:", parser3.FormatError(perr, false))
		return
	}
	cause := err
	for {
		next, ok := cause.Unwrap().(core.ErrorInterface)
		if !ok {
			break
		}
		cause = next
	}
	fmt.Fprintln(os.Stderr, "runtime error: "+cause.GetCode()+": "+cause.GetMsg())
}

// Values.

// value is what an expression evaluates to: an integer, a string, a bool, a
// list or a function.
type value interface{}

type function struct {
	name    string
	params  []string
	body    *stringParsing.ParsedNode
	closure *env
}

func (f *function) String() string {
	return "<fn " + f.name + "(" + strings.Join(f.params, ", ") + ")>"
}

// describe renders a value the way print and str do.
func describe(v value) string {
	switch t := v.(type) {
	case nil:
		return "nil"
	case bool:
		return strconv.FormatBool(t)
	case int:
		return strconv.Itoa(t)
	case string:
		return t
	case []value:
		parts := make([]string, 0, len(t))
		for _, item := range t {
			parts = append(parts, describe(item))
		}
		return "[" + strings.Join(parts, ", ") + "]"
	case *function:
		return t.String()
	default:
		return fmt.Sprintf("%v", t)
	}
}

func truthy(v value) bool {
	b, ok := v.(bool)
	return ok && b
}

// Scopes.

// env is one lexical scope. Lookup walks the parent chain, which is what makes
// a function body see the variables of its declaration site.
type env struct {
	vars   map[string]value
	parent *env
}

func newEnv(parent *env) *env {
	return &env{vars: make(map[string]value), parent: parent}
}

func (e *env) lookup(name string) (value, bool) {
	for s := e; s != nil; s = s.parent {
		if v, ok := s.vars[name]; ok {
			return v, true
		}
	}
	return nil, false
}

// declare creates a binding in this scope, which is what := does.
func (e *env) declare(name string, v value) {
	e.vars[name] = v
}

// assign overwrites an existing binding, which is what = does. It reports
// whether the name was visible at all, so a typo is an error and not a silent
// creation of a new variable.
func (e *env) assign(name string, v value) bool {
	for s := e; s != nil; s = s.parent {
		if _, ok := s.vars[name]; ok {
			s.vars[name] = v
			return true
		}
	}
	return false
}

// control is how a block reports how it finished. Loops need it to tell a
// continue from a break, and a function body turns return into a value.
type controlKind int

const (
	controlNormal controlKind = iota
	controlBreak
	controlContinue
	controlReturn
)

type control struct {
	kind  controlKind
	value value
}

// Interpreter.

type interpreter struct {
	Engine *lc.EngineUniversal
	root   *env
}

func newInterpreter() (*interpreter, error) {
	parser := parser3.NewParser(newLexer(), newGrammar(), "program", []string{"WHITESPACE", "COMMENT"})
	engine, err := lc.NewEngineBuilder(public.StringEngineType, public.StringResType).
		WithPipeline([]string{"main"}).
		WithStringParser(&parser3.Adapter{Parser: parser}).
		WithDefaultEvents(true).
		WithContext(context.Background()).
		Build()
	if err != nil {
		return nil, err
	}
	if err := engine.Plugins.AddPlugin(extensiblePlugin.New(engine)); err != nil {
		return nil, err
	}
	if err := engine.Plugins.AddPlugin(profiler.New()); err != nil {
		return nil, err
	}

	interp := &interpreter{Engine: engine, root: newEnv(nil)}
	if err := interp.registerCommands(); err != nil {
		return nil, err
	}
	return interp, nil
}

// registerCommands binds one engine command per statement form. The engine
// calls the handler of every top level node in order, which is exactly the
// statement list of the program.
func (in *interpreter) registerCommands() error {
	forms := []struct {
		name string
		doc  string
		run  func(*env, *stringParsing.ParsedNode) (control, core.ErrorInterface)
	}{
		{"fndecl", "declare a function", in.runFnDecl},
		{"decl", "declare a variable", in.runDecl},
		{"assign", "assign to a declared variable", in.runAssign},
		{"if", "conditional statement", in.runIf},
		{"while", "loop while a condition holds", in.runWhile},
		{"for", "iterate over a list", in.runFor},
		{"block", "a sequence of statements", in.runBlock},
		{"return", "return from a function", in.runReturn},
		{"break", "leave the nearest loop", func(*env, *stringParsing.ParsedNode) (control, core.ErrorInterface) {
			return control{kind: controlBreak}, nil
		}},
		{"continue", "start the next iteration", func(*env, *stringParsing.ParsedNode) (control, core.ErrorInterface) {
			return control{kind: controlContinue}, nil
		}},
		{"exprstmt", "an expression evaluated for its result", in.runExprStmt},
	}
	for _, form := range forms {
		run := form.run
		err := in.Engine.NewCommandString(form.name,
			func(_ enginepkg.StringEngineInterface, n *stringParsing.ParsedNode) core.ErrorInterface {
				ctrl, rerr := run(in.root, n)
				if rerr != nil {
					return rerr
				}
				// A loop signal that reached the top level has no loop left to
				// leave, so it is reported instead of dropped.
				switch ctrl.kind {
				case controlBreak:
					return core.Err(lce.CorePackageSystemError, "break outside of a loop")
				case controlContinue:
					return core.Err(lce.CorePackageSystemError, "continue outside of a loop")
				}
				return nil
			}, form.doc)
		if err != nil {
			return err
		}
	}
	return nil
}

// Lexer.

func newLexer() *stringParsing.Lexer {
	keyword := func(name, word string) stringParsing.LexerRule {
		return stringParsing.LexerRule{Type: name, Pattern: regexp2.MustCompile(word+`\b`, 0)}
	}
	rules := []stringParsing.LexerRule{
		{Type: "COMMENT", Pattern: regexp2.MustCompile(`//[^\n]*`, 0)},
		{Type: "NUMBER", Pattern: regexp2.MustCompile(`\d+`, 0)},
		{Type: "STRING", Pattern: regexp2.MustCompile(`"(\\.|[^"\\])*"`, 0)},
		keyword("FN", "fn"),
		keyword("IF", "if"),
		keyword("ELSE", "else"),
		keyword("WHILE", "while"),
		keyword("FOR", "for"),
		keyword("IN", "in"),
		keyword("RETURN", "return"),
		keyword("BREAK", "break"),
		keyword("CONTINUE", "continue"),
		keyword("TRUE", "true"),
		keyword("FALSE", "false"),
		{Type: "IDENT", Pattern: regexp2.MustCompile(`[a-zA-Z_][a-zA-Z0-9_]*`, 0)},
		{Type: "DECLARE", Pattern: regexp2.MustCompile(`:=`, 0)},
		{Type: "EQ", Pattern: regexp2.MustCompile(`==`, 0)},
		{Type: "NE", Pattern: regexp2.MustCompile(`!=`, 0)},
		{Type: "LE", Pattern: regexp2.MustCompile(`<=`, 0)},
		{Type: "GE", Pattern: regexp2.MustCompile(`>=`, 0)},
		{Type: "ASSIGN", Pattern: regexp2.MustCompile(`=`, 0)},
		{Type: "LT", Pattern: regexp2.MustCompile(`<`, 0)},
		{Type: "GT", Pattern: regexp2.MustCompile(`>`, 0)},
		{Type: "PLUS", Pattern: regexp2.MustCompile(`\+`, 0)},
		{Type: "MINUS", Pattern: regexp2.MustCompile(`-`, 0)},
		{Type: "STAR", Pattern: regexp2.MustCompile(`\*`, 0)},
		{Type: "SLASH", Pattern: regexp2.MustCompile(`/`, 0)},
		{Type: "PERCENT", Pattern: regexp2.MustCompile(`%`, 0)},
		{Type: "BANG", Pattern: regexp2.MustCompile(`!`, 0)},
		{Type: "AND", Pattern: regexp2.MustCompile(`&&`, 0)},
		{Type: "OR", Pattern: regexp2.MustCompile(`\|\|`, 0)},
		{Type: "LBRACKET", Pattern: regexp2.MustCompile(`\[`, 0)},
		{Type: "RBRACKET", Pattern: regexp2.MustCompile(`\]`, 0)},
		{Type: "LPAREN", Pattern: regexp2.MustCompile(`\(`, 0)},
		{Type: "RPAREN", Pattern: regexp2.MustCompile(`\)`, 0)},
		{Type: "LBRACE", Pattern: regexp2.MustCompile(`\{`, 0)},
		{Type: "RBRACE", Pattern: regexp2.MustCompile(`\}`, 0)},
		{Type: "COMMA", Pattern: regexp2.MustCompile(`,`, 0)},
		{Type: "SEMICOLON", Pattern: regexp2.MustCompile(`;`, 0)},
		{Type: "WHITESPACE", Pattern: regexp2.MustCompile(`\s+`, 0)},
	}
	return stringParsing.NewLexer(rules, &stringParsing.LexerConfig{UseBracketBalance: true})
}

// Grammar.

// newGrammar describes the language as a set of productions. PrattExpr handles
// the operators, so precedence and associativity live in one table instead of
// one rule per level.
func newGrammar() parser3.Grammar {
	expression := &parser3.PrattExpr{
		Atom: parser3.NodeExpr{
			NodeType: "atom",
			// The call alternative has to be tried before a bare identifier,
			// otherwise the identifier matches alone and the parentheses are
			// left over.
			Expr: parser3.ChoiceExpr{Alternatives: []parser3.Expr{
				parser3.TokenExpr{TokenType: "NUMBER"},
				parser3.TokenExpr{TokenType: "STRING"},
				parser3.TokenExpr{TokenType: "TRUE"},
				parser3.TokenExpr{TokenType: "FALSE"},
				listLiteral(),
				funcLiteral(),
				callExpr(),
				parenthesized(),
				parser3.TokenExpr{TokenType: "IDENT"},
			}},
		},
		Prefixes: map[string]parser3.Expr{
			"MINUS": parser3.TokenExpr{TokenType: "MINUS"},
			"BANG":  parser3.TokenExpr{TokenType: "BANG"},
		},
		Infixes: map[string]parser3.InfixInfo{
			"OR":      {Precedence: 1, Assoc: parser3.LeftAssoc},
			"AND":     {Precedence: 2, Assoc: parser3.LeftAssoc},
			"EQ":      {Precedence: 3, Assoc: parser3.LeftAssoc},
			"NE":      {Precedence: 3, Assoc: parser3.LeftAssoc},
			"LT":      {Precedence: 4, Assoc: parser3.LeftAssoc},
			"LE":      {Precedence: 4, Assoc: parser3.LeftAssoc},
			"GT":      {Precedence: 4, Assoc: parser3.LeftAssoc},
			"GE":      {Precedence: 4, Assoc: parser3.LeftAssoc},
			"PLUS":    {Precedence: 5, Assoc: parser3.LeftAssoc},
			"MINUS":   {Precedence: 5, Assoc: parser3.LeftAssoc},
			"STAR":    {Precedence: 6, Assoc: parser3.LeftAssoc},
			"SLASH":   {Precedence: 6, Assoc: parser3.LeftAssoc},
			"PERCENT": {Precedence: 6, Assoc: parser3.LeftAssoc},
		},
	}
	return parser3.Grammar{
		// The start rule is not wrapped in a node: the engine adapter hands
		// the engine the children of the root, so an extra wrapper would show
		// up as one more top level node.
		"program": {
			Name: "program",
			Expr: parser3.RepeatExpr{Expr: parser3.NamedExpr{RuleName: "statement"}},
		},
		"statement": {
			Name: "statement",
			Expr: parser3.ChoiceExpr{Alternatives: []parser3.Expr{
				parser3.NamedExpr{RuleName: "fndecl"},
				parser3.NamedExpr{RuleName: "decl"},
				parser3.NamedExpr{RuleName: "assign"},
				parser3.NamedExpr{RuleName: "if"},
				parser3.NamedExpr{RuleName: "while"},
				parser3.NamedExpr{RuleName: "for"},
				parser3.NamedExpr{RuleName: "block"},
				parser3.NamedExpr{RuleName: "return"},
				parser3.NamedExpr{RuleName: "break"},
				parser3.NamedExpr{RuleName: "continue"},
				parser3.NamedExpr{RuleName: "exprstmt"},
			}},
		},
		"fndecl": {
			Name: "fndecl",
			Expr: parser3.NodeExpr{
				NodeType: "fndecl",
				Expr: parser3.SequenceExpr{Exprs: []parser3.Expr{
					parser3.TokenExpr{TokenType: "FN"},
					parser3.OptionalExpr{Expr: parser3.TokenExpr{TokenType: "IDENT"}},
					parser3.TokenExpr{TokenType: "LPAREN"},
					parser3.OptionalExpr{Expr: parser3.SeparatedRepeatExpr{
						Element: parser3.TokenExpr{TokenType: "IDENT"},
						Sep:     "COMMA",
					}},
					parser3.TokenExpr{TokenType: "RPAREN"},
					parser3.NamedExpr{RuleName: "block"},
				}},
			},
		},
		"decl": {
			Name: "decl",
			Expr: parser3.NodeExpr{
				NodeType: "decl",
				Expr: parser3.SequenceExpr{Exprs: []parser3.Expr{
					parser3.TokenExpr{TokenType: "IDENT"},
					parser3.TokenExpr{TokenType: "DECLARE"},
					parser3.NamedExpr{RuleName: "expression"},
				}},
			},
		},
		"assign": {
			Name: "assign",
			Expr: parser3.NodeExpr{
				NodeType: "assign",
				Expr: parser3.SequenceExpr{Exprs: []parser3.Expr{
					parser3.TokenExpr{TokenType: "IDENT"},
					parser3.TokenExpr{TokenType: "ASSIGN"},
					parser3.NamedExpr{RuleName: "expression"},
				}},
			},
		},
		"exprstmt": {
			Name: "exprstmt",
			Expr: parser3.NodeExpr{
				NodeType: "exprstmt",
				Expr: parser3.SequenceExpr{Exprs: []parser3.Expr{
					parser3.NamedExpr{RuleName: "expression"},
					parser3.OptionalExpr{Expr: parser3.TokenExpr{TokenType: "SEMICOLON"}},
				}},
			},
		},
		"if": {
			Name: "if",
			Expr: parser3.NodeExpr{
				NodeType: "if",
				Expr: parser3.SequenceExpr{Exprs: []parser3.Expr{
					parser3.TokenExpr{TokenType: "IF"},
					parser3.NamedExpr{RuleName: "expression"},
					parser3.NamedExpr{RuleName: "block"},
					parser3.OptionalExpr{Expr: parser3.SequenceExpr{Exprs: []parser3.Expr{
						parser3.TokenExpr{TokenType: "ELSE"},
						parser3.NamedExpr{RuleName: "statement"},
					}}},
				}},
			},
		},
		"while": {
			Name: "while",
			Expr: parser3.NodeExpr{
				NodeType: "while",
				Expr: parser3.SequenceExpr{Exprs: []parser3.Expr{
					parser3.TokenExpr{TokenType: "WHILE"},
					parser3.NamedExpr{RuleName: "expression"},
					parser3.NamedExpr{RuleName: "block"},
				}},
			},
		},
		"for": {
			Name: "for",
			Expr: parser3.NodeExpr{
				NodeType: "for",
				Expr: parser3.SequenceExpr{Exprs: []parser3.Expr{
					parser3.TokenExpr{TokenType: "FOR"},
					parser3.TokenExpr{TokenType: "IDENT"},
					parser3.TokenExpr{TokenType: "IN"},
					parser3.NamedExpr{RuleName: "expression"},
					parser3.NamedExpr{RuleName: "block"},
				}},
			},
		},
		"block": {
			Name: "block",
			Expr: parser3.NodeExpr{
				NodeType: "block",
				Expr: parser3.SequenceExpr{Exprs: []parser3.Expr{
					parser3.TokenExpr{TokenType: "LBRACE"},
					parser3.RepeatExpr{Expr: parser3.NamedExpr{RuleName: "statement"}},
					parser3.TokenExpr{TokenType: "RBRACE"},
				}},
			},
		},
		"return": {
			Name: "return",
			Expr: parser3.NodeExpr{
				NodeType: "return",
				Expr: parser3.SequenceExpr{Exprs: []parser3.Expr{
					parser3.TokenExpr{TokenType: "RETURN"},
					parser3.OptionalExpr{Expr: parser3.NamedExpr{RuleName: "expression"}},
				}},
			},
		},
		"break": {
			Name: "break",
			Expr: parser3.NodeExpr{NodeType: "break", Expr: parser3.TokenExpr{TokenType: "BREAK"}},
		},
		"continue": {
			Name: "continue",
			Expr: parser3.NodeExpr{NodeType: "continue", Expr: parser3.TokenExpr{TokenType: "CONTINUE"}},
		},
		"expression": {Name: "expression", Expr: expression},
	}
}

func listLiteral() parser3.Expr {
	return parser3.NodeExpr{
		NodeType: "list",
		Expr: parser3.SequenceExpr{Exprs: []parser3.Expr{
			parser3.TokenExpr{TokenType: "LBRACKET"},
			parser3.OptionalExpr{Expr: parser3.SeparatedRepeatExpr{
				Element: parser3.NamedExpr{RuleName: "expression"},
				Sep:     "COMMA",
			}},
			parser3.TokenExpr{TokenType: "RBRACKET"},
		}},
	}
}

func callExpr() parser3.Expr {
	return parser3.NodeExpr{
		NodeType: "call",
		Expr: parser3.SequenceExpr{Exprs: []parser3.Expr{
			parser3.TokenExpr{TokenType: "IDENT"},
			parser3.TokenExpr{TokenType: "LPAREN"},
			parser3.OptionalExpr{Expr: parser3.SeparatedRepeatExpr{
				Element: parser3.NamedExpr{RuleName: "expression"},
				Sep:     "COMMA",
			}},
			parser3.TokenExpr{TokenType: "RPAREN"},
		}},
	}
}

// parenthesized groups an expression, which is what lets a comparison sit
// under a prefix operator, as in !(1 == 2). The node type marks the group, so
// eval can read the expression inside it.
func parenthesized() parser3.Expr {
	return parser3.NodeExpr{
		NodeType: "group",
		Expr: parser3.SequenceExpr{Exprs: []parser3.Expr{
			parser3.TokenExpr{TokenType: "LPAREN"},
			parser3.NamedExpr{RuleName: "expression"},
			parser3.TokenExpr{TokenType: "RPAREN"},
		}},
	}
}

func funcLiteral() parser3.Expr {
	return parser3.NodeExpr{
		NodeType: "fndecl",
		Expr: parser3.SequenceExpr{Exprs: []parser3.Expr{
			parser3.TokenExpr{TokenType: "FN"},
			parser3.TokenExpr{TokenType: "LPAREN"},
			parser3.OptionalExpr{Expr: parser3.SeparatedRepeatExpr{
				Element: parser3.TokenExpr{TokenType: "IDENT"},
				Sep:     "COMMA",
			}},
			parser3.TokenExpr{TokenType: "RPAREN"},
			parser3.NamedExpr{RuleName: "block"},
		}},
	}
}

// Statements.

// runBlock runs the statements of a block node. The braces are part of the
// node, so they are skipped instead of being read as statements.
func (in *interpreter) runBlock(e *env, n *stringParsing.ParsedNode) (control, core.ErrorInterface) {
	statements := make([]stringParsing.ParsedNode, 0)
	for _, child := range astools.GetChildren(n) {
		if child.Switch == "LBRACE" || child.Switch == "RBRACE" {
			continue
		}
		statements = append(statements, child)
	}
	return in.runStatements(e, statements)
}

// runStatements executes a node list in order and stops at the first control
// signal, so break, continue and return unwind the block they appear in.
func (in *interpreter) runStatements(e *env, nodes []stringParsing.ParsedNode) (control, core.ErrorInterface) {
	for i := range nodes {
		ctrl, err := in.runStatement(e, &nodes[i])
		if err != nil {
			return control{}, err
		}
		if ctrl.kind != controlNormal {
			return ctrl, nil
		}
	}
	return control{}, nil
}

// runStatement dispatches a node to the handler of its form. The node types
// are the names the engine commands are registered under, so a statement at
// the top level and one nested in a body go through the same code.
func (in *interpreter) runStatement(e *env, n *stringParsing.ParsedNode) (control, core.ErrorInterface) {
	switch n.Switch {
	case "fndecl":
		return in.runFnDecl(e, n)
	case "decl":
		return in.runDecl(e, n)
	case "assign":
		return in.runAssign(e, n)
	case "if":
		return in.runIf(e, n)
	case "while":
		return in.runWhile(e, n)
	case "for":
		return in.runFor(e, n)
	case "block":
		return in.runBlock(e, n)
	case "return":
		return in.runReturn(e, n)
	case "break":
		return control{kind: controlBreak}, nil
	case "continue":
		return control{kind: controlContinue}, nil
	case "exprstmt":
		return in.runExprStmt(e, n)
	default:
		return control{}, core.Err(lce.CorePackageSystemError, "unsupported statement: %s", n.Switch)
	}
}

func (in *interpreter) runDecl(e *env, n *stringParsing.ParsedNode) (control, core.ErrorInterface) {
	name, expr, err := namedTarget(n)
	if err != nil {
		return control{}, err
	}
	val, verr := in.eval(*expr, e)
	if verr != nil {
		return control{}, verr
	}
	e.declare(name, val)
	return control{}, nil
}

func (in *interpreter) runAssign(e *env, n *stringParsing.ParsedNode) (control, core.ErrorInterface) {
	name, expr, err := namedTarget(n)
	if err != nil {
		return control{}, err
	}
	val, verr := in.eval(*expr, e)
	if verr != nil {
		return control{}, verr
	}
	if !e.assign(name, val) {
		return control{}, core.Err(lce.CorePackageSystemError, "undefined variable: %s", name)
	}
	return control{}, nil
}

func (in *interpreter) runExprStmt(e *env, n *stringParsing.ParsedNode) (control, core.ErrorInterface) {
	_, err := in.eval(*n, e)
	return control{}, err
}

func (in *interpreter) runFnDecl(e *env, n *stringParsing.ParsedNode) (control, core.ErrorInterface) {
	fn, err := makeFunction(n, e)
	if err != nil {
		return control{}, err
	}
	if fn.name == "" {
		return control{}, core.Err(lce.CorePackageSystemError, "function declaration needs a name")
	}
	e.declare(fn.name, fn)
	return control{}, nil
}

func (in *interpreter) runIf(e *env, n *stringParsing.ParsedNode) (control, core.ErrorInterface) {
	condNode := childAt(n, 1)
	if condNode == nil {
		return control{}, core.Err(lce.CorePackageSystemError, "malformed if")
	}
	cond, err := in.eval(*condNode, e)
	if err != nil {
		return control{}, err
	}
	if truthy(cond) {
		return in.runBlock(e, childAt(n, 2))
	}
	if elseNode := childAfter(n, "ELSE"); elseNode != nil {
		return in.runStatement(e, elseNode)
	}
	return control{}, nil
}

func (in *interpreter) runWhile(e *env, n *stringParsing.ParsedNode) (control, core.ErrorInterface) {
	condNode := childAt(n, 1)
	body := childAt(n, 2)
	if condNode == nil || body == nil {
		return control{}, core.Err(lce.CorePackageSystemError, "malformed while")
	}
	for {
		cond, err := in.eval(*condNode, e)
		if err != nil {
			return control{}, err
		}
		if !truthy(cond) {
			return control{}, nil
		}
		ctrl, berr := in.runBlock(e, body)
		if berr != nil {
			return control{}, berr
		}
		if ctrl.kind == controlReturn {
			return ctrl, nil
		}
		if ctrl.kind == controlBreak {
			return control{}, nil
		}
	}
}

func (in *interpreter) runFor(e *env, n *stringParsing.ParsedNode) (control, core.ErrorInterface) {
	listNode := childAfter(n, "IN")
	body := astools.FindChild(n, "block")
	nameNode := childAt(n, 1)
	if listNode == nil || body == nil || nameNode == nil {
		return control{}, core.Err(lce.CorePackageSystemError, "malformed for")
	}
	listVal, err := in.eval(*listNode, e)
	if err != nil {
		return control{}, err
	}
	items, ok := listVal.([]value)
	if !ok {
		return control{}, core.Err(lce.CorePackageSystemError,
			"for expects a list, got %s", describe(listVal))
	}
	// A fresh child scope per iteration, so a variable declared in the body
	// does not survive into the next round.
	for _, item := range items {
		iteration := newEnv(e)
		iteration.declare(nameNode.Raw, item)
		ctrl, berr := in.runBlock(iteration, body)
		if berr != nil {
			return control{}, berr
		}
		// A return carries its value out of the whole function, a break only
		// ends this loop, so the two cannot share one branch.
		if ctrl.kind == controlReturn {
			return ctrl, nil
		}
		if ctrl.kind == controlBreak {
			return control{}, nil
		}
	}
	return control{}, nil
}

func (in *interpreter) runReturn(e *env, n *stringParsing.ParsedNode) (control, core.ErrorInterface) {
	expr := childAt(n, 1)
	if expr == nil {
		return control{kind: controlReturn, value: nil}, nil
	}
	val, err := in.eval(*expr, e)
	if err != nil {
		return control{}, err
	}
	return control{kind: controlReturn, value: val}, nil
}

// Expressions.

func (in *interpreter) eval(n stringParsing.ParsedNode, e *env) (value, core.ErrorInterface) {
	switch n.Switch {
	case "NUMBER":
		num, cerr := strconv.Atoi(n.Raw)
		if cerr != nil {
			return nil, core.Wrap(lce.ParsingError, cerr, "not a number: %s", n.Raw)
		}
		return num, nil
	case "STRING":
		return unquote(n.Raw), nil
	case "TRUE":
		return true, nil
	case "FALSE":
		return false, nil
	case "IDENT":
		if v, ok := e.lookup(n.Raw); ok {
			return v, nil
		}
		return nil, core.Err(lce.CorePackageSystemError, "undefined variable: %s", n.Raw)
	case "list":
		return in.evalList(&n, e)
	case "call":
		return in.evalCall(&n, e)
	case "fndecl":
		return makeFunction(&n, e)
	case "PrefixOp":
		return in.evalPrefix(&n, e)
	case "BinaryOp":
		return in.evalBinary(&n, e)
	case "atom":
		inner := childAt(&n, 0)
		if inner == nil {
			return nil, core.Err(lce.CorePackageSystemError, "empty expression")
		}
		return in.eval(*inner, e)
	case "group":
		inner := childAt(&n, 1)
		if inner == nil {
			return nil, core.Err(lce.CorePackageSystemError, "empty group")
		}
		return in.eval(*inner, e)
	case "exprstmt":
		inner := childAt(&n, 0)
		if inner == nil {
			return nil, core.Err(lce.CorePackageSystemError, "empty statement")
		}
		return in.eval(*inner, e)
	default:
		return nil, core.Err(lce.CorePackageSystemError, "unsupported expression: %s", n.Switch)
	}
}

func (in *interpreter) evalList(n *stringParsing.ParsedNode, e *env) (value, core.ErrorInterface) {
	items := make([]value, 0)
	for _, child := range astools.GetChildren(n) {
		if isPunctuation(child.Switch) {
			continue
		}
		val, err := in.eval(child, e)
		if err != nil {
			return nil, err
		}
		items = append(items, val)
	}
	return items, nil
}

func (in *interpreter) evalCall(n *stringParsing.ParsedNode, e *env) (value, core.ErrorInterface) {
	callee := childAt(n, 0)
	if callee == nil {
		return nil, core.Err(lce.CorePackageSystemError, "malformed call")
	}
	args := make([]value, 0)
	// The first child is the callee, the last one the closing parenthesis,
	// so only what is between them is an argument.
	for i, child := range astools.GetChildren(n) {
		if i == 0 || isPunctuation(child.Switch) {
			continue
		}
		val, err := in.eval(child, e)
		if err != nil {
			return nil, err
		}
		args = append(args, val)
	}
	if isBuiltin(callee.Raw) {
		return in.callBuiltin(callee.Raw, args)
	}
	calleeVal, ok := e.lookup(callee.Raw)
	if !ok {
		return nil, core.Err(lce.CorePackageSystemError, "unknown function: %s", callee.Raw)
	}
	fn, ok := calleeVal.(*function)
	if !ok {
		return nil, core.Err(lce.CorePackageSystemError, "%s is not a function", callee.Raw)
	}
	return in.callFunction(fn, args)
}

// callFunction binds the arguments in a scope of the closure and runs the
// body. A return unwinds the body into the value of the call.
func (in *interpreter) callFunction(fn *function, args []value) (value, core.ErrorInterface) {
	if len(args) != len(fn.params) {
		return nil, core.Err(lce.CorePackageSystemError,
			"%s takes %d argument(s), got %d", fn.name, len(fn.params), len(args))
	}
	scope := newEnv(fn.closure)
	for i, param := range fn.params {
		scope.declare(param, args[i])
	}
	// A named function can call itself, so the name is bound before the body.
	if fn.name != "" {
		scope.declare(fn.name, fn)
	}
	ctrl, err := in.runBlock(scope, fn.body)
	if err != nil {
		return nil, err
	}
	return ctrl.value, nil
}

func (in *interpreter) evalPrefix(n *stringParsing.ParsedNode, e *env) (value, core.ErrorInterface) {
	op, _ := metaString(n, parser3.MetaOperator)
	operand := metaNode(n, parser3.MetaOperand)
	if operand == nil {
		return nil, core.Err(lce.CorePackageSystemError, "prefix operator without operand")
	}
	val, err := in.eval(*operand, e)
	if err != nil {
		return nil, err
	}
	switch op {
	case "MINUS":
		num, ok := val.(int)
		if !ok {
			return nil, core.Err(lce.CorePackageSystemError, "cannot negate %s", describe(val))
		}
		return -num, nil
	case "BANG":
		return !truthy(val), nil
	default:
		return nil, core.Err(lce.CorePackageSystemError, "unknown prefix operator: %s", op)
	}
}

// evalBinary evaluates the left side first, and && and || stop before the
// right side when the answer is already known.
func (in *interpreter) evalBinary(n *stringParsing.ParsedNode, e *env) (value, core.ErrorInterface) {
	op, _ := metaString(n, parser3.MetaOperator)
	leftNode := metaNode(n, parser3.MetaLeft)
	rightNode := metaNode(n, parser3.MetaRight)
	if leftNode == nil || rightNode == nil {
		return nil, core.Err(lce.CorePackageSystemError, "malformed binary operation")
	}

	left, err := in.eval(*leftNode, e)
	if err != nil {
		return nil, err
	}
	if op == "AND" && !truthy(left) {
		return false, nil
	}
	if op == "OR" && truthy(left) {
		return true, nil
	}
	right, rerr := in.eval(*rightNode, e)
	if rerr != nil {
		return nil, rerr
	}
	return applyBinary(op, left, right)
}

func applyBinary(op string, left, right value) (value, core.ErrorInterface) {
	switch op {
	case "PLUS":
		return addValues(left, right)
	case "MINUS", "STAR", "SLASH", "PERCENT":
		return arithmetic(op, left, right)
	case "AND":
		return truthy(left) && truthy(right), nil
	case "OR":
		return truthy(left) || truthy(right), nil
	case "EQ":
		return describe(left) == describe(right), nil
	case "NE":
		return describe(left) != describe(right), nil
	case "LT", "LE", "GT", "GE":
		return compare(op, left, right)
	default:
		return nil, core.Err(lce.CorePackageSystemError, "unknown operator: %s", op)
	}
}

// addValues lets one operator serve the three types it can work with, the way
// a dynamically typed language is expected to behave.
func addValues(left, right value) (value, core.ErrorInterface) {
	switch l := left.(type) {
	case int:
		r, ok := right.(int)
		if !ok {
			return nil, typeErr("+", left, right)
		}
		return l + r, nil
	case string:
		return l + describe(right), nil
	case []value:
		r, ok := right.([]value)
		if !ok {
			return nil, typeErr("+", left, right)
		}
		return append(append(make([]value, 0, len(l)+len(r)), l...), r...), nil
	default:
		return nil, typeErr("+", left, right)
	}
}

func arithmetic(op string, left, right value) (value, core.ErrorInterface) {
	l, lok := left.(int)
	r, rok := right.(int)
	if !lok || !rok {
		return nil, typeErr(op, left, right)
	}
	switch op {
	case "MINUS":
		return l - r, nil
	case "STAR":
		return l * r, nil
	case "SLASH":
		if r == 0 {
			return nil, core.Err(lce.CorePackageSystemError, "division by zero")
		}
		return l / r, nil
	default:
		if r == 0 {
			return nil, core.Err(lce.CorePackageSystemError, "division by zero")
		}
		return l % r, nil
	}
}

func compare(op string, left, right value) (value, core.ErrorInterface) {
	var cmp int
	switch l := left.(type) {
	case int:
		r, ok := right.(int)
		if !ok {
			return nil, typeErr(op, left, right)
		}
		cmp = l - r
	case string:
		r, ok := right.(string)
		if !ok {
			return nil, typeErr(op, left, right)
		}
		cmp = strings.Compare(l, r)
	default:
		return nil, typeErr(op, left, right)
	}
	switch op {
	case "LT":
		return cmp < 0, nil
	case "LE":
		return cmp <= 0, nil
	case "GT":
		return cmp > 0, nil
	default:
		return cmp >= 0, nil
	}
}

func typeErr(op string, left, right value) core.ErrorInterface {
	return core.Err(lce.CorePackageSystemError,
		"operator %s does not accept %s and %s", op, describe(left), describe(right))
}

// Builtins.

func isBuiltin(name string) bool {
	switch name {
	case "print", "len", "str", "int", "map", "filter", "reduce":
		return true
	default:
		return false
	}
}

func (in *interpreter) callBuiltin(name string, args []value) (value, core.ErrorInterface) {
	switch name {
	case "print":
		if err := arityErr(name, args, 1); err != nil {
			return nil, err
		}
		return nil, in.emit(args[0])
	case "len":
		if err := arityErr(name, args, 1); err != nil {
			return nil, err
		}
		switch t := args[0].(type) {
		case string:
			return len(t), nil
		case []value:
			return len(t), nil
		default:
			return nil, core.Err(lce.CorePackageSystemError, "len does not accept %s", describe(args[0]))
		}
	case "str":
		if err := arityErr(name, args, 1); err != nil {
			return nil, err
		}
		return describe(args[0]), nil
	case "int":
		if err := arityErr(name, args, 1); err != nil {
			return nil, err
		}
		return toInt(args[0])
	case "map":
		return in.applyToList("map", args)
	case "filter":
		return in.applyToList("filter", args)
	case "reduce":
		return in.reduceList(args)
	default:
		return nil, core.Err(lce.CorePackageSystemError, "unknown function: %s", name)
	}
}

// emit appends one line to the engine output, at the pipeline point the
// engine was built with.
func (in *interpreter) emit(v value) core.ErrorInterface {
	uep, err := in.Engine.GetUEP()
	if err != nil {
		return core.Wrap(lce.CorePackageSystemError, err, "generator is not available")
	}
	return uep.Generator.AddString(describe(v), "main")
}

func toInt(v value) (value, core.ErrorInterface) {
	switch t := v.(type) {
	case int:
		return t, nil
	case string:
		num, cerr := strconv.Atoi(strings.TrimSpace(t))
		if cerr != nil {
			return nil, core.Wrap(lce.ParsingError, cerr, "not a number: %s", t)
		}
		return num, nil
	default:
		return nil, core.Err(lce.CorePackageSystemError, "int does not accept %s", describe(v))
	}
}

// applyToList backs map and filter, which differ only in what they do with
// the value the function returned.
func (in *interpreter) applyToList(name string, args []value) (value, core.ErrorInterface) {
	if err := arityErr(name, args, 2); err != nil {
		return nil, err
	}
	fn, ok := args[1].(*function)
	if !ok {
		return nil, core.Err(lce.CorePackageSystemError, "%s needs a function", name)
	}
	items, ok := args[0].([]value)
	if !ok {
		return nil, core.Err(lce.CorePackageSystemError, "%s needs a list", name)
	}
	result := make([]value, 0, len(items))
	for _, item := range items {
		res, err := in.callFunction(fn, []value{item})
		if err != nil {
			return nil, err
		}
		if name == "map" {
			result = append(result, res)
			continue
		}
		if truthy(res) {
			result = append(result, item)
		}
	}
	return result, nil
}

func (in *interpreter) reduceList(args []value) (value, core.ErrorInterface) {
	if err := arityErr("reduce", args, 3); err != nil {
		return nil, err
	}
	fn, ok := args[1].(*function)
	if !ok {
		return nil, core.Err(lce.CorePackageSystemError, "reduce needs a function")
	}
	items, ok := args[0].([]value)
	if !ok {
		return nil, core.Err(lce.CorePackageSystemError, "reduce needs a list")
	}
	acc := args[2]
	for _, item := range items {
		res, err := in.callFunction(fn, []value{acc, item})
		if err != nil {
			return nil, err
		}
		acc = res
	}
	return acc, nil
}

func arityErr(name string, args []value, want int) core.ErrorInterface {
	if len(args) != want {
		return core.Err(lce.CorePackageSystemError,
			"%s takes %d argument(s), got %d", name, want, len(args))
	}
	return nil
}

// Utilities.

// makeFunction reads the head and the body of a function node. The name is
// optional, which is what separates a declaration from a literal.
func makeFunction(n *stringParsing.ParsedNode, closure *env) (*function, core.ErrorInterface) {
	fn := &function{closure: closure}
	inParams := false
	for i := range astools.GetChildren(n) {
		child := astools.GetChildAt(n, i)
		switch child.Switch {
		case "IDENT":
			switch {
			case inParams:
				fn.params = append(fn.params, child.Raw)
			case fn.name == "":
				fn.name = child.Raw
			}
		case "LPAREN":
			inParams = true
		case "RPAREN":
			inParams = false
		}
	}
	body := astools.FindChild(n, "block")
	if body == nil {
		return nil, core.Err(lce.CorePackageSystemError, "function has no body")
	}
	fn.body = body
	return fn, nil
}

// namedTarget reads the name and the value of a declaration or an assignment,
// whose children are the name, the operator and the expression.
func namedTarget(n *stringParsing.ParsedNode) (string, *stringParsing.ParsedNode, core.ErrorInterface) {
	children := astools.GetChildren(n)
	if len(children) < 3 {
		return "", nil, core.Err(lce.CorePackageSystemError, "malformed statement: %s", n.Switch)
	}
	return children[0].Raw, &children[2], nil
}

// childAt returns the child at index i, or nil when the node is shorter.
func childAt(n *stringParsing.ParsedNode, i int) *stringParsing.ParsedNode {
	children := astools.GetChildren(n)
	if i < 0 || i >= len(children) {
		return nil
	}
	return &children[i]
}

// childAfter returns the node that follows the first token of the given type,
// which is how the branch of an if and the list of a for are found.
func childAfter(n *stringParsing.ParsedNode, tokenType string) *stringParsing.ParsedNode {
	children := astools.GetChildren(n)
	for i := range children {
		if children[i].Switch == tokenType && i+1 < len(children) {
			return &children[i+1]
		}
	}
	return nil
}

func metaString(n *stringParsing.ParsedNode, key string) (string, bool) {
	if n.Metadata == nil {
		return "", false
	}
	v, ok := n.Metadata[key].(string)
	return v, ok
}

// metaNode reads a sub-tree the Pratt parser stored under a metadata key.
func metaNode(n *stringParsing.ParsedNode, key string) *stringParsing.ParsedNode {
	if n.Metadata == nil {
		return nil
	}
	v, ok := n.Metadata[key].(stringParsing.ParsedNode)
	if !ok {
		return nil
	}
	return &v
}

// isPunctuation marks the tokens the grammar matched literally, which carry no
// value of their own when a node list is read as a list of operands.
func isPunctuation(sw string) bool {
	switch sw {
	case "LPAREN", "RPAREN", "LBRACKET", "RBRACKET", "COMMA", "SEMICOLON":
		return true
	default:
		return false
	}
}

func unquote(s string) string {
	if len(s) < 2 || s[0] != '"' || s[len(s)-1] != '"' {
		return s
	}
	if unquoted, err := strconv.Unquote(s); err == nil {
		return unquoted
	}
	return s[1 : len(s)-1]
}
