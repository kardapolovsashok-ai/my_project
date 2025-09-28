package main

import (
    "fmt"
    "sync"
)

const totalSeats = 38

type Cinema struct {
    seats    [totalSeats]bool 
    mutex    sync.Mutex       
    bookedCount int
}

func NewCinema() *Cinema {
    return &Cinema{bookedCount: 0}
}

func (c *Cinema) BookSeat(seatNumber int) bool {
    c.mutex.Lock()         
    defer c.mutex.Unlock() 

    if seatNumber < 1 || seatNumber > totalSeats {
        return false 
    }

    if c.seats[seatNumber-1] {
        return false 
    }

    c.seats[seatNumber-1] = true 
    c.bookedCount++
    return true
}

func (c *Cinema) AvailableSeats() int {
    return totalSeats - c.bookedCount
}

func main() {
    cinema := NewCinema()

    if cinema.BookSeat(1) {
        fmt.Println("Место 1 успешно забронировано.")
    }

    if !cinema.BookSeat(1) {
        fmt.Println("Место 1 уже забронировано.")
    }

    fmt.Println("Доступно мест:", cinema.AvailableSeats())
}
