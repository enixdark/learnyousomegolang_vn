package main

import (
    "fmt"
)

func main() {  
    
    i := 0

    for {
        i++
        if i > 10 {
            break
        }
        fmt.Println("Infinity gauntlet")
    }
}