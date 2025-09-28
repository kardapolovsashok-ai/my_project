package main
import( "fmt"; "sync"; "time")
var (tasks []string
	 mu sync.Mutex)

	 func addTask(task string) {
		mu.Lock()
		tasks = append(tasks,task)
		fmt.Println("Задача добавлена", task)
		mu.Unlock() }

	func getTask()string {
		mu.Lock()
		defer mu.Unlock()
		if len(tasks) == 0 {
			return "проектом" 
		}
		task := tasks[0]
		tasks = tasks[1:]
		fmt.Println("Взята задача", task)
		return task }

	func producer(id int, count int) {
		for i := 1; i <= count; i ++ {
			taskName := fmt.Sprintf("Добавленая задача %d-%d", id,i)
			addTask(taskName)
			time.Sleep(time.Millisecond * 200)
		}		
	}
	
	func consumer(id int) {
		for {
				task := getTask()
				if task == "" {
					fmt.Printf("Задача нет%d \n", id)
					time.Sleep(time.Second)
					continue
				}
				fmt.Printf("СМИ %d работает над: %s\n", id, task)
				time.Sleep(time.Millisecond * 500) 
		}
	}

	func main() {
		go producer(1,3)
		go producer(2,2)
		go consumer(1)
		go consumer(2)
		fmt.Println("СМИ работает")
		time.Sleep(10 * time.Second)
		fmt.Println("\n Завершено")
	}
