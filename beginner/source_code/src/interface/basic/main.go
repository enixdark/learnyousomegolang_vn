package main

import (
    "fmt"
)

type People struct {
    name string
    age int
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

// func (p People) stop() {
//     fmt.Printf("Anh %s đang đứng yên\n", p.name)
// }

func main() {
    p := People{ name: "Quan", age: 25 }
    var a Action = p
    a.walk()
    a.run()
    a.stop()
}