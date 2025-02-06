package main

import (
	"testing"
	"time"
)

func TestSemaphoreWaitGroup(t *testing.T) {
	var wg SemaphoreWaitGroup
	wg.Add(2)

	// Создаем два задания в горутинах
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
	}()

	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
	}()

	// Ожидаем завершения всех горутин
	wg.Wait()
}

func TestSemaphoreWaitGroup_ImmediateDone(t *testing.T) {
	var wg SemaphoreWaitGroup
	wg.Add(1)

	// Завершаем сразу задачу
	wg.Done()

	// Ожидаем, что программа завершится без задержки
	wg.Wait()
}

func TestSemaphoreWaitGroup_ParallelTasks(t *testing.T) {
	var wg SemaphoreWaitGroup
	wg.Add(3)

	// Три задания, выполняющиеся параллельно
	go func() {
		defer wg.Done()
		time.Sleep(500 * time.Millisecond)
	}()

	go func() {
		defer wg.Done()
		time.Sleep(300 * time.Millisecond)
	}()

	go func() {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond)
	}()

	// Ожидаем завершения всех задач
	wg.Wait()
}

func TestSemaphoreWaitGroup_NoTasks(t *testing.T) {
	var wg SemaphoreWaitGroup
	// Без задач, сразу вызываем Wait
	wg.Wait()
}
