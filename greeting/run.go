package greeting

import (
    "fmt"
)

func Run() {
		fmt.Println()
    message := Hello("Alice") // from greeting.go
    fmt.Println(message)
}
