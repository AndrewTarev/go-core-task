package main

import (
	"math"
	"sync"
	"testing"
)

// Тест обработки чисел в конвейере
func TestProcessNumbers(t *testing.T) {
	input := make(chan uint8, 5)
	output := make(chan float64, 5)
	var wg sync.WaitGroup

	wg.Add(1)
	go processNumbers(input, output, &wg)

	// Вставляем тестовые данные
	numbers := []uint8{1, 2, 3, 4, 5}
	expected := []float64{1, 8, 27, 64, 125}

	for _, num := range numbers {
		input <- num
	}
	close(input)

	// Ждем завершения обработки
	wg.Wait()
	close(output)

	// Проверяем результаты
	i := 0
	for result := range output {
		if math.Abs(result-expected[i]) > 1e-9 { // Допускаем небольшую погрешность
			t.Errorf("Expected %v, but got %v at index %d", expected[i], result, i)
		}
		i++
	}

	// Проверяем, что получили все элементы
	if i != len(expected) {
		t.Errorf("Expected %d elements, but got %d", len(expected), i)
	}
}

// Тест на пустой вводной канал
func TestEmptyInput(t *testing.T) {
	input := make(chan uint8)
	output := make(chan float64)
	var wg sync.WaitGroup

	wg.Add(1)
	go processNumbers(input, output, &wg)

	close(input)
	wg.Wait()
	close(output)

	// Проверяем, что в выходном канале ничего нет
	select {
	case result, ok := <-output:
		if ok {
			t.Errorf("Expected no values, but got %v", result)
		}
	default:
		// Ожидаемое поведение — канал пуст
	}
}

// Тест на максимальное значение uint8
func TestMaxUint8(t *testing.T) {
	input := make(chan uint8, 1)
	output := make(chan float64, 1)
	var wg sync.WaitGroup

	wg.Add(1)
	go processNumbers(input, output, &wg)

	input <- 255
	close(input)
	wg.Wait()
	close(output)

	expected := math.Pow(255, 3)
	result := <-output

	if math.Abs(result-expected) > 1e-9 {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
