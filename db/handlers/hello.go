package handlers

import (
	. "github.com/davixiao/go-db/db/resp"
)

func hello(args []Value) Value {
	return Value{Type: "array", Array: []Value{}}
}
