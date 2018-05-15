package main
import (
	"fmt"
)

func main() {
	var init map[string]int
	data := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	} 
	init["d"] = 1
	fmt.Println("Map được khơi tạo cùng với giá trị mặc định là:", init)
	fmt.Println("Map được khơi tạo cùng với giá trị xác định trước là:", data)
}
