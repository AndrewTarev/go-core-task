package main

import "testing"

func TestConvertToString(t *testing.T) {
	result := convertToString(42, 052, 0x2A, 3.14, "Golang", true, 1+2i)
	expected := "42 42 42 3.14 Golang true (1+2i)"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestHashWithSalt(t *testing.T) {
	input := "teststring"
	hashed := hashWithSalt(input)

	if len(hashed) != 64 {
		t.Errorf("Expected hash length of 64, got %d", len(hashed))
	}
}
