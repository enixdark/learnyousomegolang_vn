package main

import (
    "fmt"
)

func main() {
	days := [7]string{"Thứ Hai", "Thứ Ba", "Thứ Tư", "Thứ Năm", "Thứ Sáu", "Thứ Bảy", "Chủ Nhật"}

	school := days[0:5]
	off := days[5:]

	fmt.Println("Học sinh phải cắp sách tới trường vào ngày:", school)
	
	fmt.Println("Học sinh được nghỉ học vào ngày", off)
}