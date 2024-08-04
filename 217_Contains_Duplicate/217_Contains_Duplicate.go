package main

import (
	"fmt"
	"sort"
)

func containsDuplicate(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

func main() {
	ex1 := []int{1, 2, 3, 1}
	ex2 := []int{1, 2, 3, 4}
	ex3 := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println(containsDuplicate(ex1))
	fmt.Println(containsDuplicate(ex2))
	fmt.Println(containsDuplicate(ex3))
}
