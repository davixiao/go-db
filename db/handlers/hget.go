package handlers

import (
	. "github.com/davixiao/go-db/db/resp"
)

func hget(args []Value) Value {
	if len(args) != 2 {
		return Value{Type: "error", Str: "ERR wrong number of arguments for 'hget' command"}
	}

	hash := args[0].Bulk
	key := args[1].Bulk

	HSETsMu.RLock()
	value, ok := HSETs[hash][key]
	HSETsMu.RUnlock()

	if !ok {
		return Value{Type: "null"}
	}

	return Value{Type: "bulk", Bulk: value}
}
