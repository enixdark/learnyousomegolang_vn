package main

import (
    "fmt"
    "errors"
)

type People struct {
    name string
    age int
}

func (p People) find(peoples ...People) (error, People) {
    for _, people := range peoples {
        if p.name == people.name {
            return nil, p
        }
    }
    return errors.New("Không tìm thấy người có tên là " + p.name), People{}
}

func main() {

    peoples := []People{
        People{ name: "Quan", age: 25},
        People{ name: "Quy", age: 23 },
        People{ name: "Que", age: 18 },
    }

    p := People{ name: "Quan" }
	err, people := find(p, peoples...)
	
    if err == nil {
        fmt.Println("Thông tin về  người cần tìm là", people)
    }
}