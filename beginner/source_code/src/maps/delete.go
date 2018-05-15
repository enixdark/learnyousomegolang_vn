package main
import (
	"fmt"
)

func main() {
	data := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}

	fmt.Println("Dữ liệu trước khi xóa", data)
	
	delete(data, "A")
	delete(data, "B")
	delete(data, "A")

	fmt.Println("Dữ liệu sau khi xóa", data)

}