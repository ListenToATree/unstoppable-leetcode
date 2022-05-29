package array

//import "sort"

// Time: O(nlgn) Space: O(1)
//func missingNumber(nums []int) int {
//	sort.Ints(nums)
//	// edge case: the checking number range is 0 to n,
//	// but the index cannot be n (should be n-1)
//	var i int
//	for i = 0; i <= len(nums); i++ {
//		if i != nums[i] {
//			return i
//		}
//	}
//	return i
//}

// Gauß sum
// Time: O(n) Space: O(1)
func missingNumber(nums []int) int {
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	gauss := n * (n + 1) / 2
	return gauss - sum
}
