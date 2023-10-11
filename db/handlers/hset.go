package handlers

import (
	. "github.com/davixiao/go-db/db/resp"
)

func hset(args []Value) Value {
	if len(args) != 3 {
		return Value{Type: "error", Str: "ERR wrong number of arguments for 'hset' command"}
	}

	hash := args[0].Bulk
	key := args[1].Bulk
	value := args[2].Bulk

	HSETsMu.Lock()
	if _, ok := HSETs[hash]; !ok {
		HSETs[hash] = map[string]string{}
	}
	HSETs[hash][key] = value
	HSETsMu.Unlock()

	return Value{Type: "string", Str: "OK"}
}
