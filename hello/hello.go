package main

import (
    "fmt"

    "example.com/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("GlaDos")
    fmt.Println(message)
}

// Continue tutorial here: https://go.dev/doc/tutorial/handle-errors