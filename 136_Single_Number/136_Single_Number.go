package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	res := 0
	for _, val := range nums {
		res ^= val
	}
	return res
}

func main() {
	ex1 := []int{2, 2, 1}
	ex2 := []int{4, 1, 2, 1, 2}
	ex3 := []int{1}
	fmt.Println(singleNumber(ex1))
	fmt.Println(singleNumber(ex2))
	fmt.Println(singleNumber(ex3))
}
