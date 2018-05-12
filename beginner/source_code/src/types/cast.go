package main

import (  
    "fmt"
)

func main() {  
    var result int32
    result = 4 / 2 
	result = int32(4) / int32(2.0) 
	fmt.Println(result)
}