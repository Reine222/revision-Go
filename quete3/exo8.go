package main

import (
	"fmt"
)

func StrRev(s string) string {
	chaine := []rune(s) 

	for a,b := 0, len(s)-1 ; a < b ; a,b = a+1, b-1 {
		chaine[a], chaine[b] = chaine[b], chaine[a]
	}
	return string(chaine)
}

func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}