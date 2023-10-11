package handlers

import (
	. "github.com/davixiao/go-db/db/resp"
)

func set(args []Value) Value {
	if len(args) != 2 {
		return Value{Type: "error", Str: "ERR wrong number of arguments for 'set' command"}
	}

	key := args[0].Bulk
	value := args[1].Bulk

	SETsMu.Lock()
	SETs[key] = value
	SETsMu.Unlock()
	return Value{Type: "string", Str: "OK"}
}
