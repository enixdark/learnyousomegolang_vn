package main

import (  
    "fmt"
    "unsafe"
)

func main() {  
    var a complex64 = 1 + 2i
    var b float32 = -909.2
    var c float64 = 11111

    var d = complex(b, b)
    var e = complex(c, c)

    fmt.Println("giá trị của a là:", a, 
				"\ngiá trị của b là:", b, 
				"\ngía trị của c là:", c, 
				"\ngiá trị của d là:", d, 
				"\ngía trị của e là:", e)
	
	fmt.Printf("kiểu dữ liệu của a là %T, kích thước của a là %d byte\n", a, unsafe.Sizeof(a)) 
	fmt.Printf("kiểu dữ liệu của b là %T, kích thước của b là %d bytes\n", b, unsafe.Sizeof(b)) 
	fmt.Printf("kiểu dữ liệu của c là %T, kích thước của c là %d bytes\n", c, unsafe.Sizeof(c))
	fmt.Printf("kiểu dữ liệu của d là %T, kích thước của d là %d bytes\n", d, unsafe.Sizeof(d))
	fmt.Printf("kiểu dữ liệu của e là %T, kích thước của e là %d bytes", e, unsafe.Sizeof(e)) 
}