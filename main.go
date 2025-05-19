package main

import (
    "start.com/start/greeting" // Import your module's package
    "fmt"
)

func main() {
    message := greeting.Hello("Alice")
    fmt.Println(message)
}
