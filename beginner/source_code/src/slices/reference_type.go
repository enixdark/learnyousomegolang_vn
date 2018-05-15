package main

import (
    "fmt"
)

func main() {
	days := [7]string{"Thứ Hai", "Thứ Ba", "Thứ Tư", "Thứ Năm", "Thứ Sáu", "Thứ Bảy", "Chủ Nhật"}

	slice_days := days[0:5]
    
    slice_days[0] = "2"

	fmt.Println("Danh Sách ngày trong tuần:", days)
	
}