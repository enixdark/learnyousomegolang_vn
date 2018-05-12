package main

import (  
    "fmt"
    "unsafe"
)

func main() {  
    var a float32 = 89.1
    var b float64 = -909.2


    fmt.Println("giá trị của a là:", a, 
        "\ngiá trị của b là:", b)

    fmt.Printf("kiểu dữ liệu của a là %T, kích thước của a là %d byte\n", a, unsafe.Sizeof(a)) 
    fmt.Printf("kiểu dữ liệu của b là %T, kích thước của b là %d bytes\n", b, unsafe.Sizeof(b)) 
}