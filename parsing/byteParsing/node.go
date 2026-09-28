package byteParsing

// ParsedBytes is a single parsed instruction in binary mode.
type ParsedBytes struct {
	// Switch is the command or opcode part of the instruction.
	Switch []byte
	// Raw is the complete original byte slice this node was parsed from.
	Raw []byte
	// Args holds the raw byte slice of every argument.
	Args [][]byte
	// Metadata is free-form storage for anything the parser wants to keep.
	Metadata map[string]interface{}
}
