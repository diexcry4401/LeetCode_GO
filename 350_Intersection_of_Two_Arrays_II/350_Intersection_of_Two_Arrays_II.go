package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	freq := make(map[int]int)
	res := make([]int, 0)

	for _, num := range nums1 {
		freq[num]++
	}

	for _, num := range nums2 {
		if count, ok := freq[num]; ok && count > 0 {
			res = append(res, num)
			freq[num]--
		}
	}
	return res
}

func main() {
	nums1 := []int{4, 9, 5, 9}
	nums2 := []int{9, 4, 9, 8, 4}
	fmt.Println(intersect(nums1, nums2))
}
