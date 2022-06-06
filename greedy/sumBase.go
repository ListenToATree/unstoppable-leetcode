package greedy

func sumBase(n int, k int) (sum int) {
	for ; n > 0; n = n / k {
		sum += n % k
	}
	return
}
