package dp

//// Time: O(n) Space: O(n)
//// tabulation
//func minCostClimbingStairs(cost []int) int {
//	stairs := make([]int, len(cost)+1)
//
//	for i := 2; i <= len(cost); i++ {
//		lastOne := stairs[i-1] + cost[i-1]
//		lastTwo := stairs[i-2] + cost[i-2]
//		stairs[i] = min(lastOne, lastTwo)
//	}
//
//	return stairs[len(cost)]
//}
//
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// recursion + memo
var memo = make(map[int]int)

func minCostClimbingStairs(cost []int) int {
	return minCost(len(cost), cost)
}

func minCost(i int, cost []int) int {
	if i <= 1 {
		return 0
	}

	if v, ok := memo[i]; ok {
		return v
	}

	lastOne := cost[i-1] + minCost(i-1, cost)
	lastTwo := cost[i-2] + minCost(i-2, cost)

	memo[i] = min(lastTwo, lastOne)

	return memo[i]
}
