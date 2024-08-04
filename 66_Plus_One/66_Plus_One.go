package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] += 1
			return digits
		} else {
			digits[i] = 0
		}
	}
	var res = make([]int, n+1)
	res[0] = 1
	return res
}

func main() {
	nums1 := []int{9, 9}
	nums2 := []int{4, 3, 2, 1}
	fmt.Println(plusOne(nums1))
	fmt.Println(plusOne(nums2))
}
