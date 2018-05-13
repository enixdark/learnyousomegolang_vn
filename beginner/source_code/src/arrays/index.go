package main

import "fmt"

func main() {

	var stores [10]int

	sum := 0

	for i := 1; i <= 10; i++ {
		sum += i
		stores[i-1] = sum 
	}

	fmt.Println("kết quả lưu trữ tổng qua các lần gọi là:", stores);
}
