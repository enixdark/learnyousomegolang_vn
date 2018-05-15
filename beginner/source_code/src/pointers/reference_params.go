package main
import "fmt"

func ChangeValue(p *int, value int) {
  *p = value
}

func appendString(s* string, word string) {
  *s += word
}

func main() {
  var a int = 10
  var s string = "Hello "

  fmt.Println("Giá trị của a là:", a)
  fmt.Println("Giá trị của chuỗi s là:", s)

  ChangeValue(&a, 20)
  appendString(&s, "Go!")

  fmt.Println("Giá trị của a là:", a)
  fmt.Println("Giá trị của chuỗi s là:", s)
}