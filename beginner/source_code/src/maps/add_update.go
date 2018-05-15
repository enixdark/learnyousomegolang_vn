package main
import (
	"fmt"
)

func main() {
	init := make(map[string]int)
	
    fmt.Println("Map được khơi tạo cùng với giá trị mặc định là:", init)
    
    init["A"] = 1
    fmt.Println("Map sau khi được thêm dữ liệu:", init)
    
    init["A"] = 10
    fmt.Println("Map sau khi sửa đổi dữ liệu:", init)
}