package main
import (
	"fmt"
)

func main() {
	var data map[string]int
	origin := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	} 

	data = origin 
	fmt.Println("giá trị của origin là:", origin)

	data["D"] = 4
	fmt.Println("giá trị của origin sau khi chỉnh sửa tử  biến data là:", origin)
	append(data, "t", 1)
}
