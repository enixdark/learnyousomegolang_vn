package main

import (  
    "fmt"
)

func main() {  
    temp := make([]int, 2, 10)
	fmt.Println("Gía trị bạn đâu là:", temp)
	temp = append(temp, 1, 2)
	fmt.Println("Gía trị sau khi gọi hàm append là:", temp)
}