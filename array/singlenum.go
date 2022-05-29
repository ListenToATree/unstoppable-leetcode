package array

// t O(n^2) s O(n)
func singleNumber(nums []int) int {
	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = !set[num]
	}

	for k, v := range set {
		if v {
			return k
		}
	}
	return -1
}
