package main

import (
    "fmt"
)

func badge(b ...string) {
    fmt.Printf("Các Bắp đựng là %s\n", b)
    for _, value := range b {
        fmt.Printf("Bắp %s đã được bỏ vào\n", value)
    }
}


func main() {
	badge("1")
	badge("2", "3")
	badge("4", "5", "6", "7", "8")
}