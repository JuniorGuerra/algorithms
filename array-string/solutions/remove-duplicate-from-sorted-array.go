package solutions

import "fmt"

func RemoveDuplicates(nums []int) int {

	prev := nums[0]

	l := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != prev {
			nums[l] = nums[i]
			l++
		}
		prev = nums[i]
	}
	fmt.Println(nums)
	return l
}
