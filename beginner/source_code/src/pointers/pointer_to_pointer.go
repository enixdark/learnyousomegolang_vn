package main
import "fmt"

func main() {
  var a int = 10
  var p *int = &a
  var pp **int = &p
  
  fmt.Println("Giá trị của con trỏ p sau khi gán là:", p)
  fmt.Println("Giá trị của biến được trỏ là:", *p)

  fmt.Println("Giá trị của con trỏ pp sau khi gán là:", pp)
  fmt.Println("Giá trị của cấp 1 là:", *pp)
  fmt.Println("Giá trị của cáp 2 là:", **pp)

}
