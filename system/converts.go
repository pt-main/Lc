package system

import (
	"errors"
	"strconv"

	"github.com/pt-main/lc/stringParsing"
)

type Converts struct{}

func (c Converts) ConvertStringCommandTypeArgs(args []interface{},
) (*StringEngine, stringParsing.ParsedNode, error) {
	if len(args) != 2 {
		return nil, stringParsing.ParsedNode{}, errors.New(
			"Can't convert args: len of args is " + strconv.Itoa(len(args)) +
				" (must be 2)",
		)
	}
	stringEngine, ok1 := args[0].(*StringEngine)
	if !ok1 {
		return nil, stringParsing.ParsedNode{}, errors.New(
			"Can't get a *StringEngine from invalid interface.",
		)
	}
	parsedNode, ok2 := args[1].(stringParsing.ParsedNode)
	if !ok2 {
		return nil, stringParsing.ParsedNode{}, errors.New(
			"Can't get a ParsedNode from invalid interface.",
		)
	}
	return stringEngine, parsedNode, nil
}
