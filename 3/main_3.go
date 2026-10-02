package main

import "fmt"

type StringIntMap struct {
	inner map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		inner: make(map[string]int),
	}
}

func (m *StringIntMap) Add(key string, value int) {
	m.inner[key] = value
}

func (m *StringIntMap) Remove(key string) {
	delete(m.inner, key)
}

func (m *StringIntMap) Copy() *StringIntMap {
	newMap := NewStringIntMap()

	for key, value := range m.inner {
		newMap.Add(key, value)
	}

	return newMap
}

func (m *StringIntMap) Exists(key string) bool {
	_, ok := m.inner[key]

	return ok
}

func (m *StringIntMap) Get(key string) (int, bool) {
	val, ok := m.inner[key]

	return val, ok
}

func main() {
	siMap := NewStringIntMap()

	siMap.Add("key1", 1)
	siMap.Add("key2", 2)
	siMap.Add("key3", 3)
	fmt.Println(siMap)

	siMap.Remove("key2")
	siMap.Remove("key7")
	fmt.Println(siMap)

	copySiMap := siMap.Copy()
	copySiMap.Add("key4", 4)
	fmt.Println(copySiMap)
	fmt.Println(siMap)

	fmt.Println(siMap.Exists("key1"), siMap.Exists("key2"))

	fmt.Println(siMap.Get("key3"))
	fmt.Println(siMap.Get("key7"))
}
