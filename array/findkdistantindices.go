package array

// func findKDistantIndices(nums []int, key int, k int) []int {
//	idx, tmp := []int{}, make(map[int]bool)
//	for i, v := range nums {
//		if v == key {
//			idx = append(idx, i)
//		}
//	}
//	for i := range nums {
//		for _, j := range idx {
//			if Abs(i-j) <= k {
//				tmp[i] = true
//			}
//		}
//	}
//	res := []int{}
//	for k, _ := range tmp {
//		res = append(res, k)
//	}
//	sort.Ints(res)
//	return res
//}

// super-fast!
func findKDistantIndices(nums []int, key int, k int) []int {
	ans := []int{}
	for i, v := range nums {
		if v == key {
			left := 0
			if len(ans) != 0 {
				left = ans[len(ans)-1] + 1
			}
			left = Max(left, i-k)
			for j := left; j <= i+k && j < len(nums); j++ {
				ans = append(ans, j)
			}
		}
	}
	return ans
}
