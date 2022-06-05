package array

//// I look up the solutions to find if there is any simple way to determine whether the array is inc.
//// Then I realized that no simple way. It is a paradox! The goal is find it out.
//// So the solution mostly takes the two ways to consider. No optimization at the beginning.
//func isMonotonic(nums []int) bool {
//	if len(nums) <= 1 {
//		return true
//	}
//	isInc := (nums[1] - nums[0]) > 0
//	for i := 1; i < len(nums); i++ {
//		if isInc && (nums[i-1] > nums[i]) {
//			return false
//		}
//		if !isInc && (nums[i-1] < nums[i]) {
//			return false
//		}
//	}
//	return true
//}

//// Time: O(2n) Space: O(1)
//// two pass
//func isMonotonic(nums []int) bool {
//	return isInc(nums) || isDec(nums)
//}
//
//func isInc(nums []int) bool {
//	for i := 0; i < len(nums)-1; i++ {
//		if nums[i] > nums[i+1] {
//			return false
//		}
//	}
//	return true
//}
//
//func isDec(nums []int) bool {
//	for i := 0; i < len(nums)-1; i++ {
//		if nums[i] < nums[i+1] {
//			return false
//		}
//	}
//	return true
//}

// one pass
// Time: O(n) Space: O(1)
func isMonotonic(nums []int) bool {
	inc, dec := true, true
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			inc = false
		}
		if nums[i] < nums[i+1] {
			dec = false
		}
	}
	return inc || dec
}
