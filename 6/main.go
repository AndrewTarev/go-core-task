package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Generator struct {
	ch chan int
}

func NewGenerator() *Generator {
	return &Generator{
		ch: make(chan int),
	}
}

func (g *Generator) writeInChannel() {
	r := rand.New(rand.NewSource(time.Now().UnixNano())) // Используем локальный генератор
	number := r.Intn(100)                                // Числа от 0 до 99
	g.ch <- number
}

func (g *Generator) readFromChannel() int {
	return <-g.ch
}

func main() {
	generator := NewGenerator()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		generator.writeInChannel()
	}()

	go func() {
		defer wg.Done()
		fmt.Println(generator.readFromChannel())
	}()

	wg.Wait()
}
