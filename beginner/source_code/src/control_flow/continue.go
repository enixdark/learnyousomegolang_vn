package main

import (
    "fmt"
)

func main() {  
    
    i := 0

    for {
        i++
        if i <= 100 {
           continue
        }
        fmt.Println("level loop:", i)
        fmt.Println("Infinity gauntlet")
        break
    }
}