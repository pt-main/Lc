package stringParsing

import "github.com/pt-main/lc/v2/engine/core"

// ParsedNode is a single token or syntactic unit in text mode.
type ParsedNode struct {
	// Raw is the exact substring matched, or the full line or block.
	Raw string
	// Switch is the token type, such as "NUMBER" or "IDENT", or the command name.
	Switch string
	// Metadata holds the regexp named groups, "__raw" (the full original
	// text), "__value" (the matched value) and the "__prev" / "__next" links
	// added after parsing, which point to the neighbouring nodes or nil.
	Metadata core.ScopeType
}
