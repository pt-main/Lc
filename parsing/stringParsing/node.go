package stringParsing

import "github.com/pt-main/lc/engine/core"

// ParsedNode represents a single token or syntactic unit in text mode.
// It contains the original raw string, the token type or command name (Switch),
// and a metadata map for named groups or additional attributes.
// After parsing, nodes are automatically enriched with __prev and __next links.
type ParsedNode struct {
	// Raw string – the exact substring matched (or the full line/block).
	Raw string
	// Switch string – the token type (e.g., "NUMBER", "IDENT") or command name.
	Switch string
	// Metadata core.ScopeType – holds regexp named groups, "__raw" (full
	//   original text, in all basic Parsers), "__value" (matched value, in basic Lexer),
	//   and optionally "__prev"/"__next" which point to neighboring ParsedNode (or nil).
	Metadata core.ScopeType
}
