package main

import "fmt"

func main() {

	b := [5]int{1,2,3,4,5}
	a := b

	a[1] = 0

	fmt.Println(a)
}
