package main

import (
    "fmt"
)

func sum(a int, b int) int {
    return a + b
}

func main() {
    fmt.Println(sum(1, 2))
    fmt.Println(sum(2, 4))
}