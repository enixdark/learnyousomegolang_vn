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

	fmt.Println("Độ dài của map là:", len(data))
	fmt.Println("Các giá trị có trong map là")
	for key, value := range data {
        fmt.Println(key, " -> ", value)

    }
}