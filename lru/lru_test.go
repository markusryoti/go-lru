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
			Name: "can put values correctly",
			TestFunc: func() {
				cache := NewLRU(3)

				cache.Put("first", 3)
				verifyPut(cache, t, "first", 3, 1)

				cache.Put("second", 7)
				verifyPut(cache, t, "second", 7, 2)

				cache.Put("third", 5)
				verifyPut(cache, t, "third", 5, 3)

				cache.Put("fourth", 1)
				verifyPut(cache, t, "fourth", 1, 3)
				verifyLast(cache, t, 7)

				cache.Put("second", 2)
				verifyPut(cache, t, "second", 2, 3)
			},
		},
	}

	for _, c := range cases {
		c.TestFunc()
	}
}

func verifyPut(cache *LRU, t *testing.T, key string, value, length int) {
	if cache.length != length {
		t.Errorf("cache length wasn't updated: %d", cache.length)
	}

	if cache.list.first.value != value {
		t.Errorf("first value wasn't updated: %d", value)
	}

	n := cache.searchMap[key]
	if n == nil {
		t.Errorf("couldn't find cached value from search map")
	}
}

func verifyLast(cache *LRU, t *testing.T, value int) {
	if cache.list.last.value != value {
		t.Errorf("last value is not correct: %d", value)
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

				checkNonExisting(cache, t)

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

func checkNonExisting(cache *LRU, t *testing.T) {
	notExist := cache.Get("nonexisting")
	if notExist != -1 {
		t.Fatalf("incorrect value for non existing item: %d", notExist)
	}
}

func verifyValue(cache *LRU, t *testing.T, key string, value int, length int) {
	num := cache.Get(key)

	if num != value {
		t.Errorf("incorrect value: %d", num)
	}

	if f := cache.list.first.value; f != value {
		t.Errorf("first value in list not correct: %d", f)
	}

	if cache.length != length {
		t.Errorf("incorrect cache length: %d", length)
	}
}
