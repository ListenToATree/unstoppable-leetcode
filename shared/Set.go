package shared

// Set is a multi set data structure.
type Set[T comparable] struct {
	items map[T]int
}

// NewSet returns a new Set.
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		items: make(map[T]int),
	}
}

// Has returns whether the Set has
// a value (true) or not (false).
func (s *Set[T]) Has(value T) bool {
	_, exists := s.items[value]
	return exists
}

// Get returns the occurrence of the value.
func (s *Set[T]) Get(value T) int {
	occurrence, exists := s.items[value]
	if !exists {
		return -1
	}
	return occurrence
}

// Add adds a value in the Set.
func (s *Set[T]) Add(value T) {
	if s.Has(value) {
		s.items[value]++
	} else {
		s.items[value] = 1
	}
}

// Remove removes a value from the Set.
func (s *Set[T]) Remove(value T) {
	if occurrence, exists := s.items[value]; exists {
		if occurrence == 1 {
			delete(s.items, value)
		} else {
			s.items[value]--
		}
	}
}
