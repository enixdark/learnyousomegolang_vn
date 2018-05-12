package main

import (  
	"fmt"
	"unsafe"
)

func main() {  
	
	var a int = 1
	const b bool = true

	type another_string_type string
	const text = "untype constant"
	var c another_string_type = text



	fmt.Printf("kiểu dữ liệu của a là %T, kích thước của a là %d byte\n", a, unsafe.Sizeof(a)) 
	fmt.Printf("kiểu dữ liệu của b là %T, kích thước của b là %d byte\n", b, unsafe.Sizeof(b)) 
	fmt.Printf("kiểu dữ liệu của c là %T, kích thước của c là %d byte\n", c, unsafe.Sizeof(c))
}