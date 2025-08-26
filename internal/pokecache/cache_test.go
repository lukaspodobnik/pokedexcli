package pokecache

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}

func TestConcurrentAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cache := NewCache(interval)
	var wg sync.WaitGroup
	numOperations := 1000

	// Test concurrent Adds
	for i := 0; i < numOperations; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)
			val := []byte(fmt.Sprintf("value%d", i))
			cache.Add(key, val)
		}(i)
	}
	wg.Wait()

	// Test concurrent Gets
	for i := 0; i < numOperations; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)
			val, ok := cache.Get(key)
			if !ok {
				t.Errorf("expected to find key %s", key)
				return
			}
			expectedVal := []byte(fmt.Sprintf("value%d", i))
			if string(val) != string(expectedVal) {
				t.Errorf("expected value %s, got %s for key %s", expectedVal, val, key)
			}
		}(i)
	}
	wg.Wait()
}

func TestGetNonExistentKey(t *testing.T) {
	const interval = 5 * time.Second
	cache := NewCache(interval)

	// Try to get a key that hasn't been added
	_, ok := cache.Get("non_existent_key")
	if ok {
		t.Errorf("expected to not find non-existent key")
	}
}
