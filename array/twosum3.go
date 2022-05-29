package array

// TwoSum ** use custom set
//type TwoSum struct {
//	set *Set
//}
//
//func Constructor() TwoSum {
//	return TwoSum{set: NewSet()}
//}
//
//func (this *TwoSum) Add(number int) {
//	this.set.Add(number)
//}
//
//func (this *TwoSum) Find(value int) bool {
//	for key, _ := range this.set.items {
//		target := value - key
//		if (key == target && this.set.Get(target) > 1) ||
//			(key != target && this.set.Get(target) > 0) {
//			return true
//		}
//	}
//	return false
//}

// TwoSum ** use map
type TwoSum struct {
	cache map[int]int
}

func Constructor() TwoSum {
	return TwoSum{make(map[int]int)}
}

func (this *TwoSum) Add(number int) {
	if _, ok := this.cache[number]; !ok {
		this.cache[number] = 1
	} else {
		this.cache[number]++
	}
}

func (this *TwoSum) Find(value int) bool {
	for k, _ := range this.cache {
		target := value - k
		if count, _ := this.cache[target]; (k == target && count > 1) || (k != target && count > 0) {
			return true
		}
	}
	return false
}
