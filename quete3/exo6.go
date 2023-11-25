package main

import (
	"fmt"
)

func StrLen(s string) int {
	c := 0
	for i := range s {
		c++
		i++
		fmt.Println(c)
	}
	return c

	// ou

	// c := len(s)
	// return c
}


func main() {
	l := StrLen("Hello World!")
	fmt.Println(l)
}