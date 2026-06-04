package byteParsing

type ParsedBytes struct {
	Switch   []byte
	Raw      []byte
	Args     [][]byte
	Metadata map[string]interface{}
}
