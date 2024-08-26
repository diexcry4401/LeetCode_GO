package main

import "fmt"

/*
Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.

Example 1:
Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]

Example 2:
Input: nums = [0]
Output: [0]
*/

// func moveZeroes(nums []int) {
// 	n := len(nums)
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < n-i-1; j++ {
// 			if nums[j] == 0 {
// 				nums[j], nums[j+1] = nums[j+1], nums[j]
// 			}
// 		}
// 	}
// 	fmt.Println(nums)
// }

func moveZeroes(nums []int) {
	lastIndxNonZero := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			fmt.Printf("nums[%d] = %d switch nums[%d] = %d \n", i, nums[i], lastIndxNonZero, nums[lastIndxNonZero])
			nums[lastIndxNonZero], nums[i] = nums[i], nums[lastIndxNonZero]
			lastIndxNonZero++
		}
	}
}

func main() {
	ex1 := []int{0, 1, 0, 3, 12}
	// ex2 := []int{1, 0, 0}
	// ex3 := []int{0, 0, 1}
	moveZeroes(ex1)
	// moveZeroes(ex2)
	// moveZeroes(ex3)

}
