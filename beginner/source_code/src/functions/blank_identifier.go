package main

import (
    "fmt"
)

func head(a int, b int) (error, int) {
    return nil, a
}

func main() {
    _, value := head(1,2)
    fmt.Println(value)
}