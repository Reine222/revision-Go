package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		z01.PrintRune('\n')
	} else {
		chain := args[0]
		rune_chain := []rune(chain)
		var tab_space []int
		var tab_mot []string

		for i, v := range chain {
			if v == ' ' || v == '	' {
				tab_space = append(tab_space, i)
			}

		}

		j := 0
		for _, sp := range tab_space {
			tab_mot = append(tab_mot, string(rune_chain[j:sp]))
			if sp == tab_space[len(tab_space)-1] {
				tab_mot = append(tab_mot, string(rune_chain[sp+1:]))
			}
			j = sp + 1

		}
		//fmt.Println(tab_mot)

		// **************************************************************** firstword ******************************************************************************

		for _, v := range tab_mot[0] {
			z01.PrintRune(v)
		}

		fmt.Println()
		// ***************************************************************************** lastword **********************************************************

		for _, v := range tab_mot[len(tab_mot)-1] {
			z01.PrintRune(v)
		}

	}

}
