package main

import (
	"fmt"
	"math"
	"sync"
)

// Конвейер для преобразования и возведения чисел в куб
func processNumbers(input <-chan uint8, output chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range input {
		result := math.Pow(float64(num), 3)
		output <- result
	}
}

func main() {
	// Каналы
	input := make(chan uint8, 5)
	output := make(chan float64, 5)

	var wg sync.WaitGroup

	// Запускаем конвейер в отдельной горутине
	wg.Add(1)
	go processNumbers(input, output, &wg)

	// Запись данных в первый канал
	for i := uint8(1); i <= 5; i++ {
		input <- i
	}
	close(input)

	// Ожидаем завершения обработки
	wg.Wait()
	close(output)

	// Чтение и вывод результатов из второго канала
	for result := range output {
		fmt.Println(result)
	}
}
