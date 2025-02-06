package main

import (
	"reflect"
	"testing"
)

func TestDifference(t *testing.T) {
	tests := []struct {
		slice1   []string
		slice2   []string
		expected []string
	}{
		{
			slice1:   []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			slice2:   []string{"banana", "date", "fig"},
			expected: []string{"apple", "cherry", "43", "lead", "gno1"},
		},
		{
			slice1:   []string{"one", "two", "three"},
			slice2:   []string{"three", "four"},
			expected: []string{"one", "two"},
		},
		{
			slice1:   []string{"a", "b", "c"},
			slice2:   []string{"a", "b", "c"},
			expected: []string{},
		},
		{
			slice1:   []string{},
			slice2:   []string{"x", "y", "z"},
			expected: []string{},
		},
		{
			slice1:   []string{"unique"},
			slice2:   []string{},
			expected: []string{"unique"},
		},
	}

	for _, test := range tests {
		result := Difference(test.slice1, test.slice2)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input slice1=%v, slice2=%v, expected %v but got %v", test.slice1, test.slice2, test.expected, result)
		}
	}
}
