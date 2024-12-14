package helpers

type HashSet[T comparable] map[T]struct{}

func (s HashSet[T]) Add(item T) {
	s[item] = struct{}{}
}

func (s HashSet[T]) Remove(item T) {
	delete(s, item)
}

func (s HashSet[T]) Has(item T) bool {
	_, ok := s[item]
	return ok
}
