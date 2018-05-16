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

type Product struct {
	price int
	Computer
}

func main() {
	
	computer := Computer{
		displayScreen: "Samsung 24 inch",
		cpu: "Intel Core i3",
		hdd: 1024,
		mainboard: "Asus",
		ram: 1024,
	}
	
	product := Product{ price: 12000000 }
	product.Computer = computer
	
	fmt.Println("Thông tin sản phẩm:", product)
}