package main

import (
    "fmt"
    "errors"
)

type People struct {
    name string
    age int
}

func find(people People, peoples []People) (error, People) {
    for _, p := range peoples {
        if people.name == p.name {
            return nil, p
        }
    }
    return errors.New("Không tìm thấy người có tên là " + people.name), People{}
}

func main() {

    peoples := []People{
        People{ name: "Quan", age: 25},
        People{ name: "Quy", age: 23 },
        People{ name: "Que", age: 18 },
    }


    err, people := find(People{ name: "Quan" }, peoples)

    if err == nil {
        fmt.Println("Thông tin về  người cần tìm là", people)
    }
}