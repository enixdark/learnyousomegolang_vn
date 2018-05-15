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

    fmt.Println("A ->", data["A"])
	fmt.Println("D ->", data["D"])
	
	_, ok := data["D"]
	if ok != true {
        fmt.Println("Khóa D không tồn tại trong map")
	}
    
}