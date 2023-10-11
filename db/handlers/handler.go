package handlers

import (
	"sync"
	. "github.com/davixiao/go-db/db/resp"
)

var Handlers = map[string]func([]Value) Value {
	"HELLO": 	hello,
	"CLIENT": 	client,
	"PING": 	ping,
	"SET": 		set,
	"GET": 		get,
	"HSET":    	hset,
	"HGET":    	hget,
}

// hashmap to store in-memory data of strings
var SETs = map[string]string{}
var SETsMu = sync.RWMutex{}

// hashmap to store in-memory data of hashes
var HSETs = map[string]map[string]string{}
var HSETsMu = sync.RWMutex{}



