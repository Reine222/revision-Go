package main

import (
	"fmt"
)

func UltimatePointOne(n ***int) {
	***n++
}

func main() {
	a := 0
	b := &a
	n := &b
	UltimatePointOne(&n)
	fmt.Println(a)
}