package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 2 {
		chain1 := args[0]
		rune_chain1 := []rune(chain1)

		chain2 := args[1]
		rune_chain2 := []rune(chain2)
		var res string
		//j := 0
		i := 0
		for j := range rune_chain2 {
			if i < len(rune_chain1) {
				if rune_chain1[i] == rune_chain2[j] {
					res += string(rune_chain2[j])
					j = j + 1
					i = i + 1

				} else {
					j = j + 1

				}
			}

		}

		fmt.Println(res)
		for _, v := range res {
			z01.PrintRune(v)
		}
	}

}
