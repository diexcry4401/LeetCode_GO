package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	prev := 0
	for i := 1; i < len(nums); i++ {
		if nums[prev] != nums[i] {
			prev++
			nums[prev] = nums[i]
		}
	}
	return prev + 1
}

// func removeDuplicates(nums []int) int {
// 	ln := len(nums)
// 	if ln <= 1 {
// 		return ln
// 	}

// 	j := 0 // points to  the index of last filled position
// 	for i := 1; i < ln; i++ {
// 		if nums[j] != nums[i] {
// 			j++
// 			nums[j] = nums[i]
// 		}
// 	}
// 	fmt.Println(nums)
// 	return j + 1
// }

func main() {
	ex1 := []int{1, 1, 2}
	ex2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(ex1))
	fmt.Println(removeDuplicates(ex2))
}
