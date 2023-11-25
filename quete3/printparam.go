package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	for i, val := range args {
		for _, v := range val {
			z01.PrintRune(v)
		}
		if i != len(args)-1 {
			z01.PrintRune(' ')
		}

	}
}
