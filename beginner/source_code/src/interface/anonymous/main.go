package main

import (
	"fmt"
)

type NoInf interface {}

func test(i NoInf) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func main() {  
    str_ := "Hi"
    test(str_)
    int_ := 55
    test(int_)
    struct_ := struct {
        name string 
    }{
        name: "Cong Quan",
    }
    test(struct_)
}