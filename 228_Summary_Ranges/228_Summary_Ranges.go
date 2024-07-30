package main

import "fmt"

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
