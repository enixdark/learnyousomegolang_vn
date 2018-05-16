package main

import (
	"fmt"
)

type Computer struct {
	displayScreen string
    cpu string
    hdd int
    mainboard string
    ram int
} 

func main() {
	
	computer := Computer{
		displayScreen: "Samsung 24 inch",
		cpu: "Intel Core i3",
		hdd: 1024,
		mainboard: "Asus",
		ram: 1024,
    }

    fmt.Println("Bạn đang sở hữu một màn hình", computer.displayScreen)
    computer.displayScreen = "Samsung 27 inch"
    fmt.Println("Oh bạn vừa chuyển sang sử  dụng màn", computer.displayScreen)
}