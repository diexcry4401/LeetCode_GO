package main

import "fmt"

/*
You are given a sorted unique integer array nums.

A range [a,b] is the set of all integers from a to b (inclusive).

Return the smallest sorted list of ranges that cover all the numbers in the array exactly. That is, each element of nums is covered by exactly one of the ranges, and there is no integer x such that x is in one of the ranges but not in nums.

Each range [a,b] in the list should be output as:

"a->b" if a != b
"a" if a == b

Example:

Input: nums = [0,2,3,4,6,8,9]
Output: ["0","2->4","6","8->9"]
Explanation: The ranges are:
[0,0] --> "0"
[2,4] --> "2->4"
[6,6] --> "6"
[8,9] --> "8->9"
*/

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}

	start := nums[0]
	prev := nums[0]
	var res []string
	for _, number := range nums[1:] {
		// fmt.Printf("start - %d\nprev - %d\n", start, prev)
		if number-prev > 1 {
			if start == prev {
				res = append(res, fmt.Sprintf("%d", start))
			} else {
				res = append(res, fmt.Sprintf("%d->%d", start, prev))
			}
			start = number
		}
		prev = number
	}
	if start == prev {
		res = append(res, fmt.Sprintf("%d", start))
	} else {
		res = append(res, fmt.Sprintf("%d->%d", start, prev))
	}
	return res
}

func main() {
	nums := []int{0, 1, 2, 4, 5, 7}
	fmt.Println(summaryRanges(nums))
	nums2 := []int{0, 2, 3, 4, 6, 8, 9}
	fmt.Println(summaryRanges(nums2))
}
