package main

import (  
    "fmt"
)

func count(number int, numbers ...int) int {  
	number_count := 0

    for _, v := range numbers {
        if v == number {
            number_count++
        }
    }
    return number_count
}

func main() {  
	numbers := []int{1,2,3,4,5,6,5,4,3,2,1}
	number := 5
	total := count(number, numbers...)
	
	fmt.Println("Tổng giá trị ̀5 xuất hiện trong mảng là ", total)

    total = count(6, numbers...)
	fmt.Println("Tổng giá trị ̀6 xuất hiện trong mảng là ", total)

	total = count(0, numbers...)
	fmt.Println("Tổng giá trị ̀0 xuất hiện trong mảng là ", total)
}