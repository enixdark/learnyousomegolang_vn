package main
import "fmt"

func main() {
  var a int = 10
  var p *int 

  fmt.Println("Giá trị của biến a là:", a)
  fmt.Println("Địa chỉ của biến a là:", &a)
  fmt.Println("Giá trị của con trỏ p là:", p)
  fmt.Println("Địa chỉ của con trỏ p là:", &p)

  p = &a
  fmt.Println("")
  
  fmt.Println("Giá trị của con trỏ p sau khi gán là:", p)
  fmt.Println("Giá trị của biến được trỏ là:", *p)

  *p = 20

  fmt.Println("Giá trị của biến a sau khi thay dổi con trỏ là:", a)
}
