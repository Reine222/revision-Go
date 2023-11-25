package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 1 {
		chaine := args[0]
		verif := false
		//fmt.Println(chaine)
		for _, v := range chaine {
			if v >= '0' || v <= '9' {
				verif = true
			}
		}

		count := 0
		rep := 0

		if verif {
			for _, v := range chaine {
				for i := '0'; i < v; i++ {
					count++
				}
				rep = rep*10 + count
				count = 0
			}
		}

		fmt.Println(rep)

	} else {
		z01.PrintRune('\n')
	}

}
