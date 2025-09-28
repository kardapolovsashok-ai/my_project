package main

import (
 "fmt"
 "sync"
 "time"
)

type SafeCache struct {
 data map[string]string
 mu   sync.RWMutex
}

func (c *SafeCache) Get(k string) string {
 c.mu.RLock()
 defer c.mu.RUnlock()
 return c.data[k]
}

func (c *SafeCache) Set(k, v string) {
 c.mu.Lock()
 defer c.mu.Unlock()
 c.data[k] = v
}

func main() {
 cache := &SafeCache{data: make(map[string]string)}

 cache.Set("user", "Aleksandr")

 go func() {
  for i := 0; i < 5; i++ {
   fmt.Println("Читаю:", cache.Get("user"))
   time.Sleep(100 * time.Millisecond)
  }
 }()

 time.Sleep(300 * time.Millisecond)
 cache.Set("user", "Timur")
 fmt.Println("Обновил на Timur")

 time.Sleep(1 * time.Second)
}
