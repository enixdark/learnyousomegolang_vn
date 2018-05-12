package main

import (  
	"fmt"
	"unsafe"
)

func main() {  
    first_name := "Cong"
    last_name := "Quan"
	fmt.Printf("Tên của tôi là %s %s\n", first_name, last_name)
	fmt.Printf("Độ dài là %d\n", len(first_name)) 
    fmt.Printf("kiểu dữ liệu là %T, kích thước là %d byte\n", 
                first_name, 
                unsafe.Sizeof(first_name)) 
    fmt.Printf("kiểu dữ liệu của một phần tử  là %T, kích thước của một phần tử  là %d byte\n", 
             first_name[0], 
             unsafe.Sizeof(first_name[0])) 
}