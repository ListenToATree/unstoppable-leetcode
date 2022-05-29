package array

// Set is a multi set data structure.
type Set struct {
	items map[int]int
}

// NewMultiSet returns a new Set.
func NewMultiSet() *Set {
	return &Set{
		items: make(map[int]int),
	}
}

// Has returns whether the Set has
// a number (true) or not (false).
func (s *Set) Has(number int) bool {
	_, exists := s.items[number]
	return exists
}

// Get returns the occurence of the number.
func (s *Set) Get(number int) int {
	occurrence, exists := s.items[number]
	if !exists {
		return -1
	}
	return occurrence
}

// Add adds a number in the Set.
func (s *Set) Add(number int) {
	if s.Has(number) {
		s.items[number]++
	} else {
		s.items[number] = 1
	}
}

// Remove removes a number from the Set.
func (s *Set) Remove(number int) {
	if occurrence, exists := s.items[number]; exists {
		if occurrence == 1 {
			delete(s.items, number)
		} else {
			s.items[number]--
		}
	}
}
