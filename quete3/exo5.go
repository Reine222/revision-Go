package main

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, val := range s {
		z01.PrintRune(val)
	}
	z01.PrintRune('\n')
}

func main() {
	PrintStr("Hello World!")
}