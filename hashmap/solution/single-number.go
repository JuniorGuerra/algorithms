package solution

import "fmt"

func SingleNumber(nums []int) int {
	result := 0

	for _, val := range nums {
		result ^= val
		fmt.Println(result)

	}

	return result
}
