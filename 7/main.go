package main

import (
	"fmt"
	"sync"
)

func mergeChannels(channels []chan int) chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	for _, ch := range channels {
		wg.Add(1)
		go func(ch chan int) {
			defer wg.Done()
			for v := range ch {
				out <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	channel1 := make(chan int)
	channel2 := make(chan int)
	channel3 := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			channel1 <- i
		}
		close(channel1)
	}()

	go func() {
		for i := 5; i < 10; i++ {
			channel2 <- i
		}
		close(channel2)
	}()

	go func() {
		for i := 10; i < 15; i++ {
			channel3 <- i
		}
		close(channel3)
	}()

	merge := mergeChannels([]chan int{channel1, channel2, channel3})

	for v := range merge {
		fmt.Println(v)
	}

}
