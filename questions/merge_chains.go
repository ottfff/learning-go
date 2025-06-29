package questions

import (
	"sync"
	"time"
)

func gen(data []int, delay time.Duration) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, i := range data {
			time.Sleep(delay)
			ch <- i
		}
	}()
	return ch
}

func merge(cs ...<-chan int) <-chan int {
	result := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(len(cs))
	go func() {
		wg.Wait()
		close(result)
	}()
	for _, ch := range cs {
		go func() {
			defer wg.Done()
			for i := range ch {
				result <- i
			}
		}()
	}
	return result
}
