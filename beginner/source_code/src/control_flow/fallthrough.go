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

    switch (true) {
        case money - ( egg_price_A * 3 + tomato_price_A * 6 ) >= 0:
            fmt.Println("Mua hàng ở quầy A")
            fallthrough
        case money - ( egg_price_B * 3 + tomato_price_B * 6 ) >= 0:
            fmt.Println("Mua hàng ở quầy B")
		case money - ( egg_price_C * 3 + tomato_price_C * 6 ) >= 0:
            fmt.Println("Mua hàng ở quầy C")
    }
}