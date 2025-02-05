package main

import "testing"

func TestStringIntMap(t *testing.T) {
	m := NewStringIntMap()

	// Test Add and Get
	m.Add("key1", 100)
	if val, exists := m.Get("key1"); !exists || val != 100 {
		t.Errorf("Expected key1 to have value 100, got %d", val)
	}

	// Test Exists
	if !m.Exists("key1") {
		t.Errorf("Expected key1 to exist")
	}
	if m.Exists("key2") {
		t.Errorf("Expected key2 to not exist")
	}

	// Test Remove
	m.Remove("key1")
	if m.Exists("key1") {
		t.Errorf("Expected key1 to be removed")
	}

	// Test Copy
	m.Add("key2", 200)
	m.Add("key3", 300)
	copiedMap := m.Copy()
	if len(copiedMap) != 2 || copiedMap["key2"] != 200 || copiedMap["key3"] != 300 {
		t.Errorf("Copy() did not return correct data")
	}
}
