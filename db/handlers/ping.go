package handlers

import (
	. "github.com/davixiao/go-db/db/resp"
)

func ping(args []Value) Value {
	if len(args) == 0 {
		return Value{Type: "string", Str: "PONG"}
	}

	return Value{Type: "string", Str: args[0].Bulk}
}
