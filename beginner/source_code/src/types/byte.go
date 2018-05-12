package main

import (  
    "fmt"
    "unsafe"
)

func main() {  
    var a byte = 89
    fmt.Println("giá trị của a là:", a, )
    fmt.Printf("kiểu dữ liệu của a là %T, kích thước của a là %d byte\n", a, unsafe.Sizeof(a)) 
}