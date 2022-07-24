package main

import "github.com/markusryoti/go-lru/lru"

func main() {
	l := lru.NewLRU(5)
	l.Put("1", 5)
}
