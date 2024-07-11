package solution

func TwoSum(nums []int, target int) []int {
	hashMap := map[int]int{}
	for i, num := range nums {
		if v, ok := hashMap[target-num]; ok {
			return []int{v, i}
		}
		hashMap[num] = i
	}
	return []int{0, 0}
}
