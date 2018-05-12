package main

import (  
    "fmt"
)

func number() int {  
        num := 15 * 5 
        return num
}

func main() {

    switch num := number(); { //num is not a constant
    case num < 80:
		fmt.Printf("%d is lesser than 50\n", num)
        fallthrough
    case num > 90:
		fmt.Printf("%d is lesser than 1\n", num)
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
    case num < 200:
		fmt.Printf("%d is lesser than 200", num)
		
    }

}