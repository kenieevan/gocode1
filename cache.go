package main

import (
	"fmt"
	"sync"
)

type Cache struct {
	sync.Mutex
	mapping map[string]string //map can't be embeded... refer type can't be embed.
	//struct, int copy type can be embeded
	slice []int
	int
}

var mapping map[string]string

var cache = Cache{mapping: make(map[string]string)}

func Lookup(key string) string {
	cache.Lock()
	defer cache.Unlock()
	cache.Lock()
	v := cache.mapping[key]
	return v
}

func main() {
	fmt.Println("vim-go")
}
