package main

import (
    "fmt"
)

type Car struct {
	color string
}

type People struct {
    name string
	age int
	Car
}



func (c Car) Show() {
    fmt.Println("Màu của xe là mau", c.color)
}

func main() {

	p := People{ name: "Quan", age: 25 }
	p.Car = Car{ color: "blue"}
	     
    p.Show()
}