package main

import (
	"fmt"
)

func BasicAtoi(s string) int {
	chaine := []rune(s)
	nbr := 0
	compt := 0
	res := false

	for _, val := range chaine {
		if val <= '9' || val >= '0' {
			res = true
		}
	}

	if res {
		for _, val := range chaine {
			for i := '0' ; i < val ; i++ {
				compt++
			}
			nbr = nbr *10 + compt
			compt = 0
		}
		return nbr 
	}
	return 0
}

func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
}