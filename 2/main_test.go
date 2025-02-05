package main

import "testing"

// Unit тесты
func TestSliceExample(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	expected := []int{2, 4, 6}
	result := sliceExample(input)

	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Expected %d at index %d, got %d", expected[i], i, v)
		}
	}
}

func TestAddElements(t *testing.T) {
	input := []int{1, 2, 3}
	expected := []int{1, 2, 3, 4}
	result := addElements(input, 4)

	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Expected %d at index %d, got %d", expected[i], i, v)
		}
	}
}

func TestCopySlice(t *testing.T) {
	input := []int{1, 2, 3}
	copyList := copySlice(input)

	if len(copyList) != len(input) {
		t.Errorf("Expected length %d, got %d", len(input), len(copyList))
	}
	for i, v := range copyList {
		if v != input[i] {
			t.Errorf("Expected %d at index %d, got %d", input[i], i, v)
		}
	}
	input[0] = 99
	if copyList[0] == 99 {
		t.Errorf("Copy should not be affected by changes in original slice")
	}
}

func TestRemoveElement(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := []int{1, 2, 4, 5}
	result := removeElement(input, 2)

	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Expected %d at index %d, got %d", expected[i], i, v)
		}
	}
}
