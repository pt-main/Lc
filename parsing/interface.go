package parsing

type ParserInterface[I any, P any] interface {
	Parse(I, ...interface{}) ([]P, error)
}
