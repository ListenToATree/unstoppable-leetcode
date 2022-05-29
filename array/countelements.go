package array

//// brute force
//// Time: O(n^2) Space: O(1)
//func countElements(arr []int) int {
//	sort.Ints(arr)
//	count := 0
//	for i := 0; i < len(arr); i++ {
//		for j := i + 1; j < len(arr); j++ {
//			if arr[j] == arr[i]+1 {
//				count++
//				break
//			}
//		}
//	}
//	return count
//}

// Time: O(n) Space: O(n)
func countElements(arr []int) int {
	hash := make(map[int]bool)
	count := 0
	for i := range arr {
		hash[arr[i]] = true
	}
	for j := range arr {
		if hash[arr[j]+1] {
			count += 1
		}
	}
	return count
}
