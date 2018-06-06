package main

import (
	"os"
	"fmt"
)

type Cat interface {
	mew()
}

type Dog interface {
	go()
}

type CatDog interface {
	Cat
	Dog
}

func main() {
	var w Writer

	// os.Stdout implements Writer
	w = os.Stdout

	fmt.Fprintf(w, "hello, writer\n")
}