package handlers

import (
	. "github.com/davixiao/go-db/db/resp"
)

func get(args []Value) Value {
	if len(args) != 1 {
		return Value{Type: "error", Str: "ERR wrong number of arguments for 'get' command"}
	}

	key := args[0].Bulk

	SETsMu.RLock()
	value, ok := SETs[key]
	SETsMu.RUnlock()

	if !ok {
		return Value{Type: "null"}
	}

	return Value{Type: "bulk", Bulk: value}
}
