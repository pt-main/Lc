package parsing

import "github.com/pt-main/lc/engine/core"

type ParseOption struct {
	UEP   *core.UniversalEngineParams
	Flags []string
	Other map[string]interface{}
}

type ParserInterface[I any, P any] interface {
	Parse(I, ...*ParseOption) ([]P, error)
	String() string
}
