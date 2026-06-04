package byteParsing

type ParserInterface interface {
	Parse(code []byte) ([]ParsedBytes, error)
}
