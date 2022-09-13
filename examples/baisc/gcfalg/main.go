package main

import (
	"fmt"
)

type obj struct{}

func main() {
	a := &obj{}
	fmt.Printf("%p\n", a)
	b := &obj{}
	println(b)
}
