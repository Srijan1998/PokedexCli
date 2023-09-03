package pokecache

import "time"
import "sync"

type CacheEntry struct {
  createdAt time.Time
  val []byte
}

type PokeCache struct {
  interval int
  cache map[string]CacheEntry
  mu sync.Mutex
}

func (pokecache PokeCache) Add(key string, val []byte) {
  pokecache.mu.Lock()
  pokecache.cache[key] = CacheEntry{time.Now(), val}
  pokecache.mu.Unlock()
}

func (pokecache PokeCache) Get(key string)([]byte, bool) {
  pokecache.mu.Lock()
  val, ok := pokecache.cache[key]
  pokecache.mu.Unlock()
  if !ok {
    return nil, ok
  }
  return val.val, ok
}

func (pokecache PokeCache) removeFromCacheOlderThanIntervalItems() {
  pokecache.mu.Lock()
  for key, element := range pokecache.cache {
    if time.Now().Sub(element.createdAt).Seconds() > (time.Duration(pokecache.interval) * time.Second).Seconds() {
      delete(pokecache.cache, key)
    }
  }
  pokecache.mu.Unlock()
}

func (pokecache PokeCache) reapLoop() {
  ticker := time.NewTicker(time.Duration(pokecache.interval) * time.Second)
  quit := make(chan struct{})
  go func(pokeCache PokeCache) {
      for {
         select {
          case <- ticker.C:
              pokeCache.removeFromCacheOlderThanIntervalItems()
          case <- quit:
              ticker.Stop()
              return
          }
      }
   }(pokecache)
}

func NewCache(interval int)(PokeCache) {
  cache := PokeCache{}
  cache.interval = interval
  cache.cache = make(map[string]CacheEntry)
  cache.reapLoop()
  return cache
}
