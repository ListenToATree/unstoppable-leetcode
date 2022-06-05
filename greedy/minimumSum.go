package greedy

import "sort"

func minimumSum(num int) int {
	var arr []int

	for num > 0 {
		arr = append(arr, num%10)
		num /= 10
	}

	// [0,1,2,3]
	sort.Ints(arr)
	return arr[0]*10 + arr[2] + arr[1]*10 + arr[3]
}
