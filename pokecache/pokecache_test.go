package pokecache

import "testing"
import "time"

func TestAddGet(t *testing.T) {
  cache := NewCache(5)
  url := "https://example.com"
  response := []byte("test string")
  cache.Add(url, response)
  cacheResponse, ok := cache.Get(url)
  if !ok || string(cacheResponse) != string(response) {
    t.Fatalf("Response Does not match")
  }
}

func TestReapLoop(t *testing.T) {
  cache := NewCache(1)
  url := "https://example.com"
  response := []byte("test string")
  cache.Add(url, response)
  time.Sleep(2*time.Second)
  _, ok := cache.Get(url)
  if ok {
    t.Fatalf("Reap loop not working")
  }
}
