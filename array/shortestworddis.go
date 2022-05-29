package array

//// brute force
//// Time: O(n^2) Space: O(1)
//func shortestDistance(wordsDict []string, word1 string, word2 string) int {
//	min := len(wordsDict)
//	for i := 0; i < len(wordsDict); i++ {
//		if wordsDict[i] == word1 {
//			for j := 0; j < len(wordsDict); j++ {
//				if wordsDict[j] == word2 {
//					min = Min(Abs(i-j), min)
//				}
//			}
//		}
//	}
//	return min
//}

// one pass
func shortestDistance(wordsDict []string, word1 string, word2 string) int {
	i1, i2 := -1, -1
	min := len(wordsDict)
	for i := 0; i < len(wordsDict); i++ {
		if wordsDict[i] == word1 {
			i1 = i
		} else if wordsDict[i] == word2 {
			i2 = i
		}

		if i1 != -1 && i2 != -1 {
			// return 1
			if Abs(i1-i2) == 1 {
				return 1
			}
			min = Min(min, Abs(i1-i2))
		}
	}
	return min
}
