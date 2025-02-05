package main

import "sync"

// StringIntMap - потокобезопасная структура для хранения пар "строка - целое число"
type StringIntMap struct {
	mu   sync.RWMutex
	data map[string]int
}

// NewStringIntMap создаёт новый экземпляр StringIntMap
func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		data: make(map[string]int),
	}
}

// Add добавляет элемент в карту
func (m *StringIntMap) Add(key string, value int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value
}

// Remove удаляет элемент по ключу
func (m *StringIntMap) Remove(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.data, key)
}

// Copy возвращает копию карты
func (m *StringIntMap) Copy() map[string]int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	copyMap := make(map[string]int, len(m.data))
	for k, v := range m.data {
		copyMap[k] = v
	}
	return copyMap
}

// Exists проверяет наличие ключа в карте
func (m *StringIntMap) Exists(key string) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	_, exists := m.data[key]
	return exists
}

// Get возвращает значение по ключу и флаг успешности операции
func (m *StringIntMap) Get(key string) (int, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	val, exists := m.data[key]
	return val, exists
}
