package hashset

type HashSet[T comparable] map[T]struct{}

func NewHashSet[T comparable]() HashSet[T] {
	return make(HashSet[T])
}

func HashSetFromSlice[T comparable](slice []T) HashSet[T] {
	set := NewHashSet[T]()

	for _, v := range slice {
		set.Add(v)
	}

	return set
}

func (s HashSet[T]) Contains(value T) bool {
	_, found := s[value]
	return found
}

func (s *HashSet[T]) Remove(value T) {
	delete((*s), value)
}

func (s *HashSet[T]) Add(value ...T) {
	for _, v := range value {
		(*s)[v] = struct{}{}
	}
}

func (s HashSet[T]) Union(other HashSet[T]) HashSet[T] {
	panic("Not Implemented")
}

func (s HashSet[T]) Intersect(other HashSet[T]) HashSet[T] {
	panic("Not Implemented")
}
