package main

import (  
	"fmt"
	"unsafe"
)

func main() {  
	
	const a  = 1
	var b int8 = a
	var c int32 = a
	var d float32 = a
	var e complex64 = a

	type another_string_type string
	const text = "untype constant"
	var f another_string_type = text

	fmt.Printf("kiểu dữ liệu của a là %T, kích thước của a là %d byte\n", a, unsafe.Sizeof(a)) 
	fmt.Printf("kiểu dữ liệu của b là %T, kích thước của b là %d byte\n", b, unsafe.Sizeof(b)) 
	fmt.Printf("kiểu dữ liệu của c là %T, kích thước của c là %d byte\n", c, unsafe.Sizeof(c)) 
	fmt.Printf("kiểu dữ liệu của d là %T, kích thước của d là %d byte\n", d, unsafe.Sizeof(d)) 
	fmt.Printf("kiểu dữ liệu của e là %T, kích thước của e là %d byte\n", e, unsafe.Sizeof(e)) 
	fmt.Printf("kiểu dữ liệu của f là %T, kích thước của f là %d byte\n", e, unsafe.Sizeof(e))
}