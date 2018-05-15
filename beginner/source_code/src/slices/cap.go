package main

import (
    "fmt"
)

func main() {
	days := [7]string{"Thứ Hai", "Thứ Ba", "Thứ Tư", "Thứ Năm", "Thứ Sáu", "Thứ Bảy", "Chủ Nhật"}

	schools := days[0:5]

	fmt.Println("len:", len(schools), "cap:", cap(schools))
    	
	fmt.Println("Học sinh phải cắp sách tới trường vào ngày:", schools)
	fmt.Println("Hm nhưng nhà trường chợt nhận ra các em học sinh rất ham hoc vì vậy")

	schools = schools[:6]
	fmt.Println("Học sinh phải cắp sách tới trường vào ngày:", schools)
}