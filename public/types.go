package public

type ResType int

const (
	ByteResType ResType = iota
	StringResType
)

type EndianType int

const (
	BigEndian EndianType = iota
	LittleEndian
)

type EngineType int

const (
	ByteEngineType = iota
	StringEngineType
)
