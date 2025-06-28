package _map

import (
	"fmt"
	"sync"
	"testing"
)

func TestMap(t *testing.T) {
	var m1 map[string]int // nil
	m2 := make(map[string]int)
	m3 := make(map[string]int, 100)
	m4 := map[string]int{
		"one": 1,
		"two": 2,
	}

	i, ok := m1["a"]
	fmt.Println(i, ok)
	// m1["a"] = 1   panic

	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3)
	fmt.Println(m4)
}

func TestSyncMap(t *testing.T) {
	m := sync.Map{} // any to any
	m.Store("a", "1")
	m.Store(1, "a")

	if v, ok := m.Load("a"); ok {
		fmt.Println("a =", v)
	}

	m.Delete(1)
	m.Delete(2)
	m.Delete(true)

	m.Range(func(key, value interface{}) bool {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
		return true // продолжить обход
	})

	var i interface{} = 10
	var j interface{} = "10"
	v := i.(int) // ok
	//v2 := j.(int) // panic
}
