package main
import ( "fmt"; "sync"; "time")

type Store struct {
 items map[string]int
 mutex sync.Mutex}

func main() {
 store := &Store{items: make(map[string]int)}
 var wg sync.WaitGroup
 
 wg.Add(1)
 go func() {
  defer wg.Done()
  for i := 0; i < 3; i++ {
   store.mutex.Lock()
   store.items["мороженое"] += 10
   fmt.Printf("привезли мороженое: %d\n", store.items["мороженое"])
   store.mutex.Unlock()
   time.Sleep(100 * time.Millisecond)
   }
  }()

 wg.Add(1)
 go func() {
  defer wg.Done()
  for i := 0; i < 3; i++ {
   store.mutex.Lock()
   if store.items["мороженое"] > 0 {
    store.items["мороженое"] -= 5
    fmt.Printf("Продали мороженое: %d\n", store.items["мороженое"])
   } else {
    fmt.Println("мороженого нет")
   }
   store.mutex.Unlock()
   time.Sleep(150 * time.Millisecond)
  }
 }()
 
 wg.Wait()
 fmt.Printf("Осталось мороженого: %d\n", store.items["мороженое"])
}
