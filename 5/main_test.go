package main

import (
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {
	tests := []struct {
		a        []int
		b        []int
		expected bool
		values   []int
	}{
		{
			a:        []int{65, 3, 58, 678, 64},
			b:        []int{64, 2, 3, 43},
			expected: true,
			values:   []int{64, 3},
		},
		{
			a:        []int{1, 2, 3},
			b:        []int{4, 5, 6},
			expected: false,
			values:   []int{},
		},
		{
			a:        []int{},
			b:        []int{1, 2, 3},
			expected: false,
			values:   []int{},
		},
		{
			a:        []int{1, 2, 3},
			b:        []int{},
			expected: false,
			values:   []int{},
		},
		{
			a:        []int{10, 20, 30},
			b:        []int{20, 10, 40},
			expected: true,
			values:   []int{20, 10},
		},
	}

	for _, test := range tests {
		result, values := Intersections(test.a, test.b)
		if result != test.expected || !reflect.DeepEqual(values, test.values) {
			t.Errorf("For input a=%v, b=%v, expected (%v, %v) but got (%v, %v)", test.a, test.b, test.expected, test.values, result, values)
		}
	}
}
