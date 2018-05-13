package main

import (  
    "fmt"
)

func main() {  
    temp := make([]int, 2, 10)
	fmt.Println(temp)
	temp = temp[:10]
	fmt.Println(temp)
}
