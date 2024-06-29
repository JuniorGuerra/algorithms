package solutions

import "fmt"

func RemoveElement(nums []int, val int) int {

	narr := nums
	nums = []int{}
	for _, valu := range narr {
		if valu != val {
			nums = append(nums, valu)

		}

	}
	narr = nums
	fmt.Println(narr)
	return len(nums)
}
