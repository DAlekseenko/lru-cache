package storage

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var testValues = map[string]string{"First": "First Value", "Second": "Second Value", "Third": "Third Value"}

func TestLRU_Add(t *testing.T) {

	req := require.New(t)

	t.Run("should add three value in cache", func(t *testing.T) {
		lru := NewLRUCache(3)
		for key, val := range testValues {
			req.True(lru.Add(key, val))
		}
	})

	t.Run("should return false if key in cache", func(t *testing.T) {
		lru := NewLRUCache(3)
		for key, val := range testValues {
			req.True(lru.Add(key, val))
		}
		req.False(lru.Add("First", "Test"))
	})

	t.Run("should add value and delete less used", func(t *testing.T) {
		lru := NewLRUCache(3)

		lru.Add("First", "Test")
		lru.Add("Second", "Test")
		lru.Add("Third", "Test")

		req.True(lru.Add("Fourth", "Test"))
		value, ok := lru.Get("First")
		req.False(ok)
		req.Empty(value)
	})
}

func TestLRU_Get(t *testing.T) {

	req := require.New(t)

	t.Run("should get element from cache", func(t *testing.T) {
		lru := NewLRUCache(3)
		k := "First"
		for key, val := range testValues {
			req.True(lru.Add(key, val))
		}
		value, ok := lru.Get(k)
		req.True(ok)
		req.Equal(testValues[k], value)
	})

	t.Run("should answer correctly if value not exist", func(t *testing.T) {
		lru := NewLRUCache(3)

		for key, val := range testValues {
			req.True(lru.Add(key, val))
		}

		value, ok := lru.Get("Test")
		req.False(ok)
		req.Empty(value)
	})

}

func TestLRU_Remove(t *testing.T) {

	req := require.New(t)

	t.Run("should remove element from cache", func(t *testing.T) {
		lru := NewLRUCache(3)
		k := "First"
		for key, val := range testValues {
			req.True(lru.Add(key, val))
		}
		ok := lru.Remove(k)
		req.True(ok)

		value, exist := lru.Get(k)

		req.Empty(value)
		req.False(exist)
	})

	t.Run("should answer correctly if value not exist", func(t *testing.T) {
		lru := NewLRUCache(3)

		for key, val := range testValues {
			req.True(lru.Add(key, val))
		}

		ok := lru.Remove("Test")
		req.False(ok)
	})

}
