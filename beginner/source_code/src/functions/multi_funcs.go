package main

import (
    "fmt"
)

func badge1(a string) {
    fmt.Printf("Bắp %s đã được bỏ vào\n", a)
}

func badge2(a string, b string) {
	fmt.Printf("Bắp %s đã được bỏ vào\n", a)
	fmt.Printf("Bắp %s đã được bỏ vào\n", b)
}

func badge5(a string, b string, c string, d string, e string) {
	fmt.Printf("Bắp %s đã được bỏ vào\n", a)
	fmt.Printf("Bắp %s đã được bỏ vào\n", b)
	fmt.Printf("Bắp %s đã được bỏ vào\n", b)
	fmt.Printf("Bắp %s đã được bỏ vào\n", d)
	fmt.Printf("Bắp %s đã được bỏ vào\n", e)
}


func main() {
	badge1("1")
	badge2("2", "3")
	badge5("4", "5", "6", "7", "8")
}