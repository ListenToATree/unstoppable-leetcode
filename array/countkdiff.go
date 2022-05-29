package array

func countKDifference(nums []int, k int) int {
	count := 0
	m := [101]int{}
	for _, num := range nums {
		m[num]++
	}
	for i := 0; i < len(m)-k; i++ {
		count += m[i] * m[i+k]
	}
	return count
}
