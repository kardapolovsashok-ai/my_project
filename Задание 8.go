package main
import ( "fmt"; "math/rand"; "sync"; "time")
type Task struct {
id int 
data string}

type Result struct { 
taskID int 
output string}

func main() {
 rand.Seed(time.Now().UnixNano())
 
 tasks := make(chan Task, 10)
 results := make(chan Result, 10)
 
 var stats struct {
sync.Mutex
processed int
failed  int}
 var wgWorkers sync.WaitGroup
 var wgProducers sync.WaitGroup

 numWorkers := 3
 for i := 1; i <= numWorkers; i++ {
  wgWorkers.Add(1)
  go worker(i, tasks, results, &stats, &wgWorkers)
 }

 numProducers := 2
 for i := 1; i <= numProducers; i++ {
  wgProducers.Add(1)
  go producer(i, tasks, &wgProducers)
 }

 go func() {
  for result := range results {
   fmt.Printf("Результат: %s\n", result.output)
  }
 }()
 
 wgProducers.Wait()
 close(tasks)
 
 wgWorkers.Wait()
 close(results)
 
 stats.Lock()
 fmt.Printf("\nСтатистика:\nОбработано: %d\nОшибок: %d\n", 
  stats.processed, stats.failed)
 stats.Unlock()
}

func producer(id int, tasks chan<- Task, wg *sync.WaitGroup) {
 defer wg.Done()
 
 for i := 1; i <= 5; i++ {
  task := Task{
   id:   id*100 + i,
   data: fmt.Sprintf("задача от продюсера %d", id),
  }
  tasks <- task
  fmt.Printf("Продюсер %d отправил задачу %d\n", id, task.id)
  time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
 }
}

func worker(id int, tasks <-chan Task, results chan<- Result, 
 stats *struct {
  sync.Mutex
  processed int
  failed    int
 }, wg *sync.WaitGroup) {
 defer wg.Done()
 
 for task := range tasks {
  fmt.Printf("Воркер %d начал задачу %d\n", id, task.id)
 
  time.Sleep(time.Millisecond * time.Duration(rand.Intn(200)))
  
  success := rand.Intn(10) > 2 

  stats.Lock()
  if success {
   stats.processed++
   results <- Result{
    taskID: task.id,
    output: fmt.Sprintf("Задача %d выполнена воркером %d", task.id, id),
   }
  } else {
   stats.failed++
   results <- Result{
    taskID: task.id,
    output: fmt.Sprintf("Задача %d провалилась воркером %d", task.id, id),
   }
  }
  stats.Unlock()
  
  fmt.Printf("Воркер %d закончил задачу %d\n", id, task.id)
 }
 }
