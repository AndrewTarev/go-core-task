package main

import (
	"testing"
	"time"
)

func TestGenerator_WriteInChannel(t *testing.T) {
	generator := NewGenerator()

	// Запускаем горутину для записи в канал
	go generator.writeInChannel()

	// Даем время для записи в канал
	time.Sleep(10 * time.Millisecond)

	select {
	case val := <-generator.ch:
		// Проверяем, что значение в канале - это число в пределах от 0 до 99
		if val < 0 || val >= 100 {
			t.Errorf("expected value between 0 and 99, got %d", val)
		}
	default:
		t.Error("expected value to be written to the channel, but it was not")
	}
}

func TestGenerator_ReadFromChannel(t *testing.T) {
	generator := NewGenerator()

	// Генерируем случайное число и записываем в канал
	go generator.writeInChannel()

	// Читаем из канала
	val := generator.readFromChannel()

	// Проверяем, что значение в канале - это число в пределах от 0 до 99
	if val < 0 || val >= 100 {
		t.Errorf("expected value between 0 and 99, got %d", val)
	}
}
