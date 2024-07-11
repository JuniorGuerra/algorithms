package solution

func FindRestaurant(list1 []string, list2 []string) []string {
	if len(list1) < len(list2) {
		list1, list2 = list2, list1
	}
	hashMap := map[string]int{}
	for i, val := range list1 {
		hashMap[val] = i
	}
	var result []string
	var least int = len(list1) + len(list2)
	for i, val := range list2 {
		if idx, ok := hashMap[val]; ok {
			sum := i + idx
			if sum < least {
				least = sum
				result = []string{val}
				continue
			} else if sum == least {
				result = append(result, val)
			}
		}
	}

	return result
}
