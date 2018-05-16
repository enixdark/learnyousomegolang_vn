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
	
	var clone Computer = computer
	fmt.Println("Trạng thái ban đâu là", computer)

	clone.hdd = 256
	fmt.Println("Trạng thái máy tính lúc sau là", computer)

	var pointer* Computer = &computer
	(*pointer).hdd =  512
	fmt.Println("Trạng thái máy tính lúc sau là", computer)

	var phdd *int = &computer.hdd
	*phdd = 2048
	fmt.Println("Trạng thái máy tính tiếp sau là", computer)
}