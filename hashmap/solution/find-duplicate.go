package solution

func FindDuplicates(arr []int) bool {
	hashMap := map[int]int{}
	for _, val := range arr {
		if _, ok := hashMap[val]; ok {
			return true
		}
		hashMap[val] = 0
	}
	return false
}
