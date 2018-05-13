package main

import (
    "fmt"
    _ "package/cmath"
)

func init() {  
    fmt.Println("Hàm main đang được gọi")
}

func main() {
    a := 6
    b := 2

    add := cmath.Add(a, b)
    sub := cmath.Sub(a, b)
    mul := cmath.Mul(a, b) 
    div := cmath.Div(a, b)

    fmt.Printf("Tổng của hai số  %d và %d là : %d\n", a, b, add)
    fmt.Printf("Hiệu của hai số  %d và %d là : %d\n", a, b, sub)
    fmt.Printf("Tích của hai số  %d và %d là : %d\n", a, b, mul)
    fmt.Printf("Thương của hai số  %d và %d là : %d\n", a, b, div)
}