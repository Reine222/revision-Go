package main
import (
	"fmt"
)

func PointOne(n *int) {
	*n++
}

func main() {
	n := 0
	PointOne(&n)
	fmt.Println(n)
}