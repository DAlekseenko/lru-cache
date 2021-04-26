package main

import (
	"fmt"
	"lru/pkg/storage"
)

func main() {
	cache := storage.NewLRUCache(3)

	cache.Add("HELLO", "WORLD")
	cache.Add("HELLO-2", "WORLD-2")
	cache.Add("HELLO-3", "WORLD-3")

	fmt.Println(cache.Get("HELLO"))

	cache.Add("HELLO-4", "WORLD-4")
	cache.Remove("HELLO-3")

	for _, v := range []string{"HELLO", "HELLO-4", "HELLO-2", "HELLO-3", "HELLO-4"} {
		fmt.Println(cache.Get(v))
	}

}
