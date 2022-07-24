package lru

import (
	"testing"
)

func TestPut(t *testing.T) {
	cases := []struct {
		Name     string
		TestFunc func()
	}{
		{
			Name: "can add to empty lru",
			TestFunc: func() {
				cache := NewLRU(5)
				cache.Put("first", 3)

				if cache.length != 1 {
					t.Errorf("cache length wasn't updated")
				}

				if cache.list.first.value != 3 {
					t.Errorf("first value wasn't updated")
				}

				n := cache.searchMap["first"]
				if n == nil {
					t.Errorf("couldn't find cached value from search map")
				}

				if cache.list.first != n || cache.list.last != n {
					t.Errorf("first or last cache value doesn't equal the newly created node")
				}

				if cache.list.last != n {
					t.Errorf("node was not updated as last one")
				}
			}},
		{
			Name: "can update existing value",
			TestFunc: func() {
				cache := NewLRU(5)
				cache.Put("first", 3)
				cache.Put("second", 7)
				cache.Put("third", 5)

				cache.Put("second", 1)

				n := cache.searchMap["second"]

				if n.value != 1 {
					t.Errorf("cache value wasn't updated, was: %d", n.value)
				}

				if cache.list.first != n {
					t.Errorf("node wasn't updated to be first in cache")
				}

				if last := cache.searchMap["first"]; cache.list.last != last {
					t.Errorf("last node not correct, had value: %+v", last)
				}
			},
		},
		{
			Name: "updated correctly when capacity exceeded",
			TestFunc: func() {
				cache := NewLRU(3)
				cache.Put("first", 3)
				cache.Put("second", 7)
				cache.Put("third", 5)
				cache.Put("fourth", 1)

				if cache.list.first.value != 1 {
					t.Errorf("first value not correct: %d", cache.list.first.value)
				}

				if cache.list.last.value != 7 {
					t.Errorf("last value not correct: %d", cache.list.last.value)
				}
			},
		},
	}

	for _, c := range cases {
		c.TestFunc()
	}
}

func TestGet(t *testing.T) {
	cases := []struct {
		Name     string
		TestFunc func()
	}{
		{
			Name: "can get values",
			TestFunc: func() {
				cache := NewLRU(3)

				notExist := cache.Get("nonexisting")
				if notExist != -1 {
					t.Fatalf("incorrect value for non existing item: %d", notExist)
				}

				cache.Put("first", 3)
				verifyValue(cache, t, "first", 3, 1)

				cache.Put("second", 7)
				verifyValue(cache, t, "second", 7, 2)

				cache.Put("third", 5)
				verifyValue(cache, t, "third", 5, 3)

				cache.Put("fourth", 1)
				verifyValue(cache, t, "fourth", 1, 3)
			},
		},
	}

	for _, c := range cases {
		c.TestFunc()
	}
}

func verifyValue(cache *LRU, t *testing.T, key string, value int, len int) {
	num := cache.Get(key)

	if num != value {
		t.Errorf("incorrect value: %d", num)
	}

	if f := cache.list.first.value; f != value {
		t.Errorf("first value in list not correct: %d", f)
	}

	if cache.length != len {
		t.Errorf("incorrect cache length: %d", len)
	}
}
