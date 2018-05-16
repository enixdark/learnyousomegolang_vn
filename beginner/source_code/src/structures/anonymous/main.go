package main

import (
	"fmt"
)

func main() {
	
	var computerA struct {
        displayScreen string
        cpu string
        hdd int
        mainboard string
        ram int
    } = struct {
        displayScreen string
        cpu string
        hdd int
        mainboard string
        ram int
    } {
        displayScreen: "Samsung 24 inch",
		cpu: "Intel Core i3",
		hdd: 1024,
		mainboard: "Asus",
		ram: 1024,
	}
	
	computerB := struct { 
		string
		int
	} { "Samsung 20 inch", 256 }

    fmt.Println("Chiếc máy tính A có thông số  là:", computerA)
	fmt.Println("Chiếc máy tính B có thông số  là:", computerB)
	fmt.Println("HDD của chiếc máy tính B là:", computerB.int)

   
}