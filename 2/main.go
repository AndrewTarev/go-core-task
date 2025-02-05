package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Генерация случайного слайса
func generateRandomSlice(n int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano())) // Используем локальный генератор
	result := make([]int, n)
	for i := range result {
		result[i] = r.Intn(100) // Числа от 0 до 99
	}
	return result
}

// Возвращает слайс только с четными числами
func sliceExample(slice []int) []int {
	var result []int
	for _, v := range slice {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	return result
}

// Добавляет элемент в конец слайса
func addElements(slice []int, num int) []int {
	return append(slice, num)
}

// Создает копию слайса
func copySlice(slice []int) []int {
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	return newSlice
}

// Удаляет элемент по индексу
func removeElement(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	originalSlice := generateRandomSlice(10)
	fmt.Println("Original Slice:", originalSlice)

	evenSlice := sliceExample(originalSlice)
	fmt.Println("Even Numbers Slice:", evenSlice)

	newSlice := addElements(originalSlice, 99)
	fmt.Println("After Adding Element:", newSlice)

	copiedSlice := copySlice(originalSlice)
	fmt.Println("Copied Slice:", copiedSlice)

	removedSlice := removeElement(originalSlice, 2)
	fmt.Println("After Removing Element at Index 2:", removedSlice)
}
