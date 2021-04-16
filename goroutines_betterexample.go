package main
import (
    "fmt"
    "time"
)

func main() {
    go start()
    fmt.Println("Started")
    time.Sleep(2 * time.Second)
    fmt.Println("Finished")
}

func start() {
    fmt.Println("In Goroutine")
    time.Sleep(3 * time.Second)
    fmt.Println("Finished Goroutine")
}
