package solution

func Intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	hashMap := map[int]bool{}
	for _, v := range nums1 {
		hashMap[v] = true
	}
	var nums []int
	for _, v := range nums2 {
		if add, ok := hashMap[v]; ok && add {
			hashMap[v] = false
			nums = append(nums, v)
		}
	}
	return nums
}

/*
// validate values
	[4, 9, 5] | [9, 4, 9, 8, 4]
	4, 9


*/
