package main

import "fmt"

func Max(arr []int) int {
	res := arr[0]
	for i := range arr {
		if res > arr[i] {
		} else {
			res = arr[i]
		}

	}
	return res
}

func main() {
	arrInt := []int{23, 123, 1, 11, 550, 93}
	max := Max(arrInt)
	fmt.Println(max)
}
