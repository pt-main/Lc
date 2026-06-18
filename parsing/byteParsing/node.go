package byteParsing

// ParsedBytes represents a single parsed instruction or block in binary mode.
// It holds the raw byte slice of the entire instruction, the command identifier
// (Switch) as a byte slice, the list of argument byte slices (Args), and
// optional Metadata for additional context (e.g., line numbers, offsets).
type ParsedBytes struct {
	// Switch []byte – the command/opcode portion of the instruction.
	Switch []byte
	// Raw []byte – the complete original byte slice that produced this node.
	Raw []byte
	// Args [][]byte – each element is a raw byte slice of an argument.
	Args [][]byte
	// Metadata map[string]interface{} – extensible storage for extra info
	//   (e.g., "offset": 42, "line": 5, "source_file": "main.asm").
	Metadata map[string]interface{}
}
