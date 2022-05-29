package array

func twoSum(nums []int, target int) [2]int {
	set := make(map[int]int)
	for i, num := range nums {
		c := target - num
		if j, ok := set[c]; ok {
			return [2]int{i, j}
		}
		set[num] = i
	}
	return [2]int{-1, -1}
}
