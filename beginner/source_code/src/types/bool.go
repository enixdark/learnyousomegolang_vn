package main

import "fmt"

func main() {  
	x := true
	y := false
	z := x && y
	k := x || y
	m := !x
	n := x != y
    fmt.Println("x:", x, "y:", y, "z:", z, "k:",k, "m:", m, "n:", n)
}