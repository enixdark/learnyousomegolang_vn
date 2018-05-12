package main

import (
    "fmt"
	"math/rand"
	"time"
)

func main() {  
	rand.Seed(time.Now().UnixNano())
	
    lucky_number := rand.Intn(10)
    
    switch (lucky_number) {
        case 1:
            fmt.Println("1 là số  may mắn")
        case 2:
            fmt.Println("2 là số  may mắn")
        case 3:
            fmt.Println("3 là số  may mắn")
        case 4:
            fmt.Println("4 là số  may mắn")
        case 5:
            fmt.Println("5 là số  may mắn")
        case 6: 
            fmt.Println("6 là số  may mắn")
        case 7:
            fmt.Println("7 là số  may mắn")
        case 8:
            fmt.Println("8 là số  may mắn")
        case 9:
            fmt.Println("9 là số  may mắn")
        case 10:
            fmt.Println("10 là số  may mắn")
    }
}