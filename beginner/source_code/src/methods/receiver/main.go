package main

import (  
    "fmt"
)

type People struct {
    name string
    age int
}

func (p People) changeName1(name string) {  
    p.name = name
}

func (p *People) changeName2(name string) {  
    p.name = name
}

func main() {  
    p := People{
        name: "Quan",
        age:  25,
    }
    fmt.Println("Thông tin của một người: ", p)
    
    p.changeName1("Cong")

    fmt.Println("Thông tin sau khi gọi một phương thức với tham trị để  thay đổi tên là: ", p)

    (&p).changeName2("Cong")
    fmt.Println("Thông tin sau khi gọi một phương thức với tham chiếu để  thay đổi tên là: ", p)
}