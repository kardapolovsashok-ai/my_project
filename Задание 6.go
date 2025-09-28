package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

var mu sync.Mutex

func log(msg string) {
	mu.Lock()
	fmt.Printf("[%s] %s\n", time.Now().Format("15:04:05"), msg)
	mu.Unlock()
}

func worker(id int) {
	for i := 0; i < 5; i++ {
		log(fmt.Sprintf("рабочий %d: действие %d", id, i))
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id)
		}(i)
	}
	wg.Wait()
	fmt.Println("все действия выполнены")
}
