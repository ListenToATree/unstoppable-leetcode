package dp

func fib(N int) int {
	fibs := make([]int, 31)
	if N < 2 {
		return N
	}
	fibs[0], fibs[1] = 0, 1
	for i := 2; i <= N; i++ {
		fibs[i] = fibs[i-1] + fibs[i-2]
	}
	return fibs[N]
}
