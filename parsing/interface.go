package parsing

type ParserInterface interface {
	Parse(code string) ([]ParsedNode, error)
}
