package main

import (
	"github.com/01-edu/z01"
)

func main () {
	for i:='a' ; i<='z' ; i++ {
		a := i
		z01.PrintRune(a)
	}
	z01.PrintRune('\n')
	
}