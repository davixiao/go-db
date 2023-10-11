package resp

// Supported RESP Datatype symbols
const (
	STRING 	= '+'
	ERROR 	= '-'
	INTEGER = ':'
	BULK 	= '$'
	ARRAY 	= '*'
)

type Value struct {
	Type string
	Str string
	Num int
	Bulk string
	Array []Value
}
