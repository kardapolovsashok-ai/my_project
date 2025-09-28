package main
import("fmt";"sync";"time")
var views int = 0
var mu sync.Mutex
func visitPage(id int) {
	mu.Lock()
	views = views + 1
	fmt.Printf("Горутина %d увеличила счётчик. Сейчас: %d\n", id, views)
	mu.Unlock()
	time.Sleep(time.Millisecond * 10)
}
func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			visitPage(id)
		}(i)
	}
	wg.Wait()
	fmt.Println("=")
	fmt.Printf("Итого: %d\n", views)
}

