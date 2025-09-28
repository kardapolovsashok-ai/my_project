package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var votes = make(map[string]int)
var mu sync.Mutex

func vote(name string) {
	mu.Lock()
	votes[name]++
	mu.Unlock()
}

func voter(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	candidates := []string{"Путин", "Медведев", "Стрельцов"}
	for i := 0; i < 5; i++ {
		c := candidates[rand.Intn(3)]
		vote(c)
		fmt.Printf("Горутина %d: %s\n", id, c)
		time.Sleep(time.Millisecond * 50)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go voter(i, &wg)
	}

	wg.Wait()

	fmt.Println("\nИтоги:")
	mu.Lock()
	for k, v := range votes {
		fmt.Printf("%s: %d\n", k, v)
	}
	mu.Unlock()
}
