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
	
	computerA := Computer{
		displayScreen: "Samsung 24 inch",
		cpu: "Intel Core i3",
		hdd: 1024,
		mainboard: "Asus",
		ram: 1024,
    }

    computerB := Computer{"Samsung 20 inch", "AMD Ryzen 5", 256, "Asrock", 8000}
    
    computerC := Computer{
        displayScreen: "Dell 2̃7 inch",
    }

    fmt.Println("Anh A đang sở hữu chiếc máy tính có cấu hình", computerA)
    fmt.Println("Anh B đang sở hữu chiếc máy tính có cấu hình", computerB)
    fmt.Println("Anh C đang sở hữu chiếc máy tính có cấu hình", computerC)

}