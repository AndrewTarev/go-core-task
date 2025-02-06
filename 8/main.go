package main

import (
	"fmt"
)

// SemaphoreWaitGroup - кастомная версия WaitGroup, использующая канал в качестве семафора.
type SemaphoreWaitGroup struct {
	sem chan struct{}
}

// NewSemaphoreWaitGroup создает новый SemaphoreWaitGroup с заданным количеством разрешений.
func NewSemaphoreWaitGroup(count int) *SemaphoreWaitGroup {
	return &SemaphoreWaitGroup{
		sem: make(chan struct{}, count), // Канал с емкостью count
	}
}

// Add добавляет новое разрешение.
func (swg *SemaphoreWaitGroup) Add(n int) {
	// Инициализируем канал если он nil
	if swg.sem == nil {
		swg.sem = make(chan struct{}, n)
	}

	// Добавляем n разрешений в канал
	for i := 0; i < n; i++ {
		swg.sem <- struct{}{} // Записываем в канал для блокировки
	}
}

// Done уменьшает счетчик и освобождает одно разрешение.
func (swg *SemaphoreWaitGroup) Done() {
	<-swg.sem // Извлекаем из канала одно разрешение
}

// Wait ожидает, пока все операции завершатся.
func (swg *SemaphoreWaitGroup) Wait() {
	// Ожидаем пока все разрешения будут освобождены
	for len(swg.sem) > 0 {
		<-swg.sem
	}
}

func main() {
	var wg SemaphoreWaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		fmt.Println("Task 1")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Task 2")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Task 3")
	}()

	wg.Wait()
	fmt.Println("All tasks completed")
}
