package concurrency

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(4)
	callDone(&wg, 4)
	wg.Wait()
}

// try "go vet"
func someFunction(wg sync.WaitGroup) {
	wg.Done()
}

func callDone(wg *sync.WaitGroup, times int) {
	for range times {
		wg.Done()
	}
}

func TestMutex(t *testing.T) {
	var mu sync.Mutex
	mu.Lock()

	var rwMutex sync.RWMutex
	rwMutex.Lock()
	rwMutex.Unlock()
	rwMutex.TryLock()
	rwMutex.RLock()
	rwMutex.RUnlock()
	rwMutex.TryRLock()
	rwMutex.RLocker()
}

func TestRuntime(t *testing.T) {
	println(runtime.NumCPU())
	fmt.Println(runtime.NumCPU())
}

func TestPool(t *testing.T) {
	type A struct {
		name string
	}
	var pool = sync.Pool{
		New: func() interface{} {
			return A{name: "a0"}
		},
	}
	o0 := pool.Get().(A)
	pool.Put(A{name: "a1"})
	pool.Put(A{name: "a2"})
	o1 := pool.Get().(A)
	o2 := pool.Get().(A)
	fmt.Println(o0)
	fmt.Println(o1)
	fmt.Println(o2)
}
