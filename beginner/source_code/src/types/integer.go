package main

import (  
    "fmt"
    "unsafe"
)

func main() {  
    var a int8 = 89
    var a1 uint = 90
    var b int16 = 1111
    var b1 uint16 = 1112
    var c int32 = 1111111
    var c1 uint32 = 1111112
    var d int64 = 1111111111111
    var d1 uint64 = 1111111111112
  
	e := 95

    fmt.Println("giá trị của a là:", a, 
        "\ngiá trị của a1 là:", a1, 
        "\ngiá trị của b là:", b, 
        "\ngía trị của b1 là:", b1, 
        "\ngía trị của c là:", c, 
        "\ngiá trị của c1 là:", c1, 
		"\ngiá trị của d là:", d, 
        "\ngía trị của d1 là:", d1,
        "\ngía trị của e là:", e)

    fmt.Printf("kiểu dữ liệu của a là %T, kích thước của a là %d byte\n", a, unsafe.Sizeof(a)) 
    fmt.Printf("kiểu dữ liệu của a1 là %T, kích thước của a1 là %d byte\n", a1, unsafe.Sizeof(a1)) 
    fmt.Printf("kiểu dữ liệu của b là %T, kích thước của b là %d bytes\n", b, unsafe.Sizeof(b)) 
    fmt.Printf("kiểu dữ liệu của b1 là %T, kích thước của b1 là %d bytes\n", b1, unsafe.Sizeof(b1)) 
    fmt.Printf("kiểu dữ liệu của c là %T, kích thước của c là %d bytes\n", c, unsafe.Sizeof(c))
    fmt.Printf("kiểu dữ liệu của c1 là %T, kích thước của c1 là %d bytes\n", c1, unsafe.Sizeof(c1)) 
    fmt.Printf("kiểu dữ liệu của d là %T, kích thước của d là %d bytes\n", d, unsafe.Sizeof(d))
    fmt.Printf("kiểu dữ liệu của d1 là %T, kích thước của d1 là %d bytes\n", d1, unsafe.Sizeof(d1)) 
	fmt.Printf("kiểu dữ liệu của e là %T, kích thước của e là %d bytes", e, unsafe.Sizeof(e)) 
}