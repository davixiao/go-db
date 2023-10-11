package handlers

import (
	. "github.com/davixiao/go-db/db/resp"
)

func client(args []Value) Value {
	return Value{Type: "array", Array: []Value{}}
}
