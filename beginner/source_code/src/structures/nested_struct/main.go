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
    computer Computer
    price int
}

func main() {
	
	computer := Computer{
		displayScreen: "Samsung 24 inch",
		cpu: "Intel Core i3",
		hdd: 1024,
		mainboard: "Asus",
		ram: 1024,
	}
	
	product := Product{
        computer: computer,
        price: 12000000,
	}
	
	fmt.Println("Thông tin sản phẩm:", product)
}