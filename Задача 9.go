package main

import (
    "fmt"
    "math/rand" 
    "sync"
    "time"
)

type Metrics struct {
    sync.Mutex           
    TotalRequests   int  
    Successful     int  
    Failures       int  
    ResponseTimeMs float64 
}

func (m *Metrics) IncrementTotal() {
    m.Lock()
    defer m.Unlock()
    m.TotalRequests++
}

func (m *Metrics) IncrementSuccess() {
    m.Lock()
    defer m.Unlock()
    m.Successful++
}

func (m *Metrics) IncrementFailure() {
    m.Lock()
    defer m.Unlock()
    m.Failures++
}

func (m *Metrics) UpdateResponseTime(timeMs float64) {
    m.Lock()
    defer m.Unlock()
    total := float64(m.TotalRequests)
    if total > 0 {
        m.ResponseTimeMs = ((m.ResponseTimeMs*(total-1)) + timeMs)/total
    }
}

func (m *Metrics) Report() string {
    return fmt.Sprintf("Метрики:\n"+
        "\tОбщее число запросов: %d\n"+
        "\tУспешные запросы: %d\n"+
        "\tНеудачные запросы: %d\n"+
        "\tСреднее время ответа: %.2fms",
        m.TotalRequests,
        m.Successful,
        m.Failures,
        m.ResponseTimeMs)
}

func worker(id int, metrics *Metrics) {
    rand.Seed(time.Now().UnixNano()) 
    for i := 0; i < 10; i++ {
        start := time.Now()
        time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond) 
        end := time.Since(start)
        metrics.IncrementTotal()
        if rand.Intn(10)%2 == 0 { 
            metrics.IncrementSuccess()
        } else {
            metrics.IncrementFailure()
        }
        metrics.UpdateResponseTime(float64(end.Nanoseconds()/1e6))
        fmt.Printf("Воркет #%d завершил обработку запроса №%d.\n", id, i+1)
    }
}

func main() {
    var wg sync.WaitGroup
    metrics := &Metrics{}

    wg.Add(3)
    go func() { defer wg.Done(); worker(1, metrics) }()
    go func() { defer wg.Done(); worker(2, metrics) }()
    go func() { defer wg.Done(); worker(3, metrics) }()

    wg.Wait()
    fmt.Println("\nОтчет:")
    fmt.Println(metrics.Report())
}
