package main

import "fmt"

func main() {  
    var money int 
    fmt.Printf("trong ví tôi có: %d VND.\n", money)
    money = 10000 // gán giá trị cho biến vừa khởi tạo
    fmt.Printf("trong ví tôi có: %d VND.\n", money)
    money = 20000 // tiếp tục gán giá trị cho biến vừa khởi tạo
    fmt.Printf("trong ví tôi có: %d VND.", money)
}

