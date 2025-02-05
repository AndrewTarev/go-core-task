package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
)

// convertToString преобразует переменные в строку и объединяет их
func convertToString(numDecimal, numOctal, numHexadecimal int, pi float64, name string, isActive bool, complexNum complex64) string {
	return fmt.Sprintf("%d %d %d %.2f %s %t %v", numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
}

// hashWithSalt принимает строку, вставляет в середину соль "go-2024", преобразует в SHA256 и возвращает хэш
func hashWithSalt(input string) string {
	runes := []rune(input)
	mid := len(runes) / 2
	runes = append(runes[:mid], append([]rune("go-2024"), runes[mid:]...)...)

	hash := sha256.Sum256([]byte(string(runes)))
	return hex.EncodeToString(hash[:])
}

func main() {
	var (
		numDecimal     int       = 42
		numOctal       int       = 052
		numHexadecimal int       = 0x2A
		pi             float64   = 3.14
		name           string    = "Golang"
		isActive       bool      = true
		complexNum     complex64 = 1 + 2i
	)

	fmt.Println("Типы переменных:")
	fmt.Printf("numDecimal: %s\n", reflect.TypeOf(numDecimal))
	fmt.Printf("numOctal: %s\n", reflect.TypeOf(numOctal))
	fmt.Printf("numHexadecimal: %s\n", reflect.TypeOf(numHexadecimal))
	fmt.Printf("pi: %s\n", reflect.TypeOf(pi))
	fmt.Printf("name: %s\n", reflect.TypeOf(name))
	fmt.Printf("isActive: %s\n", reflect.TypeOf(isActive))
	fmt.Printf("complexNum: %s\n", reflect.TypeOf(complexNum))

	convertedStr := convertToString(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	fmt.Println("Преобразованная строка:", convertedStr)

	hashedValue := hashWithSalt(convertedStr)
	fmt.Println("SHA256 хэш с солью:", hashedValue)
}
