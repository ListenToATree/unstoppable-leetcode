package string

func slowestKey(releaseTimes []int, keysPressed string) byte {
	max := releaseTimes[0]
	index := 0
	for i := 1; i < len(releaseTimes); i++ {
		dur := releaseTimes[i] - releaseTimes[i-1]
		if dur > max || (max == dur && keysPressed[i] > keysPressed[index]) {
			max = dur
			index = i
		}
	}
	return keysPressed[index]
}
