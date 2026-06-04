package stringParsing

type ParserInterface interface {
	Parse(code string) ([]ParsedNode, error)
}
