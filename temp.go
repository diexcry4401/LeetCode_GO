package main

import "fmt"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}

	start := nums[0]
	prev := nums[0]
	var res []string
	for _, elem := range nums[1:] {
		if elem-prev > 1 {
			if start == prev {
				res = append(res, fmt.Sprintf("%d", start))
			} else {
				res = append(res, fmt.Sprintf("%d->%d", start, prev))
			}
			start = elem
		}
		prev = elem
	}
	if start == prev {
		res = append(res, fmt.Sprintf("%d", start))
	} else {
		res = append(res, fmt.Sprintf("%d->%d", start, prev))
	}
	return res
}

func main() {
	ex1 := []int{0, 2, 3, 4, 6, 8, 9}
	fmt.Println(summaryRanges(ex1))
}
