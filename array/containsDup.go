package array

var cache = map[int]bool{}

func containsDuplicate(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		if _, ok := cache[nums[i]]; ok {
			return true
		}
		cache[nums[i]] = true
	}
	return false
}
