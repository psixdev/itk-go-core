package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getDefaultMap() *StringIntMap {
	m := NewStringIntMap()

	m.Add("key1", 1)
	m.Add("key2", 2)
	m.Add("key3", 3)

	return m
}

func TestMain(m *testing.M) {
	code := m.Run()

	os.Exit(code)
}

func TestNew(t *testing.T) {
	m := NewStringIntMap()

	assert.Equal(t, m.inner, map[string]int{})
}

func TestAdd(t *testing.T) {
	m := NewStringIntMap()

	m.Add("key1", 1)
	m.Add("key2", 2)
	m.Add("key3", 3)

	assert.Equal(t, m.inner, map[string]int{
		"key1": 1,
		"key2": 2,
		"key3": 3,
	})
}

func TestRemove(t *testing.T) {
	t.Run("existing key", func(t *testing.T) {
		m := getDefaultMap()

		m.Remove("key2")

		assert.Equal(t, m.inner, map[string]int{
			"key1": 1,
			"key3": 3,
		})
	})

	t.Run("non-existent key", func(t *testing.T) {
		m := getDefaultMap()

		m.Remove("key7")

		assert.Equal(t, m.inner, map[string]int{
			"key1": 1,
			"key2": 2,
			"key3": 3,
		})
	})
}

func TestCopy(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		m := NewStringIntMap()

		nm := m.Copy()

		assert.Equal(t, m.inner, nm.inner)
		assert.NotSame(t, &m.inner, &nm.inner)
	})

	t.Run("non-empty", func(t *testing.T) {
		m := getDefaultMap()

		nm := m.Copy()

		assert.Equal(t, m.inner, nm.inner)
		assert.NotSame(t, &m.inner, &nm.inner)
	})
}

func TestExists(t *testing.T) {
	m := getDefaultMap()

	t.Run("existing key", func(t *testing.T) {
		assert.True(t, m.Exists("key2"))
	})

	t.Run("non-existent key", func(t *testing.T) {
		assert.False(t, m.Exists("key7"))
	})
}

func TestGet(t *testing.T) {
	m := getDefaultMap()

	t.Run("existing key", func(t *testing.T) {
		v, ok := m.Get("key2")

		require.True(t, ok)
		assert.Equal(t, v, 2)
	})

	t.Run("non-existent key", func(t *testing.T) {
		_, ok := m.Get("key7")

		assert.False(t, ok)
	})
}
