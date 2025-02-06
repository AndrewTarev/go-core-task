package main

import (
	"testing"
)

func TestMergeChannels(t *testing.T) {
	// Создаем несколько каналов для теста
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Заполнение каналов данными в горутинах
	go func() {
		for i := 0; i < 3; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	go func() {
		for i := 3; i < 6; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	// Слияние каналов
	merged := mergeChannels([]chan int{ch1, ch2})

	// Ожидаемые значения
	expected := []int{0, 1, 2, 3, 4, 5}

	// Чтение из merged и проверка соответствия
	var result []int
	for val := range merged {
		result = append(result, val)
	}

	if len(result) != len(expected) {
		t.Errorf("expected %d elements, got %d", len(expected), len(result))
	}

	// Проверка значений
	for i, val := range expected {
		if result[i] != val {
			t.Errorf("expected %d, got %d", val, result[i])
		}
	}
}

func TestMergeChannelsWithEmpty(t *testing.T) {
	// Тест с пустыми каналами
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Заполнение каналов данными в горутинах
	go func() {
		close(ch1)
	}()
	go func() {
		close(ch2)
	}()

	// Слияние каналов
	merged := mergeChannels([]chan int{ch1, ch2})

	// Ожидаемые значения
	expected := []int{}

	// Чтение из merged и проверка соответствия
	var result []int
	for val := range merged {
		result = append(result, val)
	}

	if len(result) != len(expected) {
		t.Errorf("expected %d elements, got %d", len(expected), len(result))
	}
}

func TestMergeChannelsSingleChannel(t *testing.T) {
	// Тест с одним каналом
	ch1 := make(chan int)

	// Заполнение канала данными
	go func() {
		for i := 0; i < 5; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	// Слияние каналов
	merged := mergeChannels([]chan int{ch1})

	// Ожидаемые значения
	expected := []int{0, 1, 2, 3, 4}

	// Чтение из merged и проверка соответствия
	var result []int
	for val := range merged {
		result = append(result, val)
	}

	if len(result) != len(expected) {
		t.Errorf("expected %d elements, got %d", len(expected), len(result))
	}

	// Проверка значений
	for i, val := range expected {
		if result[i] != val {
			t.Errorf("expected %d, got %d", val, result[i])
		}
	}
}
