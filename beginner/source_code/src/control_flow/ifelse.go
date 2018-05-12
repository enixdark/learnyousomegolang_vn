package main

import (
    "fmt"
)

func main() {  
    
    money := 15000

    egg_price_A := 3000 
    tomato_price_A := 1000
    
    egg_price_B := 2500 
    tomato_price_B := 2500

    egg_price_C := 3000 
    tomato_price_C := 1000

    if money - ( egg_price_A * 3 + tomato_price_A * 6 ) >= 0 {
        fmt.Println("Mua hàng ở quầy A")
    } else if money - ( egg_price_B * 3 + tomato_price_B * 6 ) >= 0 {
        fmt.Println("Mua hàng ở quầy B")
    } else if money - ( egg_price_C * 3 + tomato_price_C * 6 ) >= 0 {
        fmt.Println("Mua hàng ở quầy C")
    } else {
        fmt.Println("Không đủ tiền thì về  nhà thôi")
    }
}