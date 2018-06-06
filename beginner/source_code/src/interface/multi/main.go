package main

import (
    "fmt"
)

type People struct {
    name string
    age int
}

type Cat struct {
    name string
}

type Dog struct {
    name string
}

type Action interface {
    run() 
    stop() 
    walk()
}

func (p People) run() {
    fmt.Printf("Anh %s đang chạy\n", p.name)
}

func (p People) walk() {
    fmt.Printf("Anh %s đang đi bộ\n", p.name)
}

func (p People) stop() {
    fmt.Printf("Anh %s đang đứng yên\n", p.name)
}

func (c Cat) run() {
    fmt.Printf("Con mèo %s đang đứng yên\n", c.name)
}

func (c Cat) walk() {
    fmt.Printf("Con mèo %s đang đi bộ\n", c.name)
}

func (c Cat) stop() {
    fmt.Printf("Con mèo %s đang đứng yên\n", c.name)
}

func (c* Dog) run() {
    fmt.Printf("Con chó %s đang đứng yên\n", c.name)
}

func (c* Dog) walk() {
    fmt.Printf("Con chó %s đang đi bộ\n", c.name)
}

func (c* Dog) stop() {
    fmt.Printf("Con chó %s đang đứng yên\n", c.name)
}

func main() {
    p := People{ name: "Quan", age: 25 }
    c := Cat{ name: "Meww" }
    d := &Dog{ name: "Nuu" }
    a := []Action{p, c, d}
    for _, d := range a{
        d.walk()
        d.run()
        d.stop()
    }
}