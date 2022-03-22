package strm

// Plus copies the backing slices contents of both Streams into a new *Stream[T]
func (s *Stream[T]) Plus(other *Stream[T]) *Stream[T] {
	merged := make([]T, len(s.filteredSlice()), len(s.slice)+len(other.filteredSlice()))
	copy(merged, s.slice)
	return From(append(merged, other.slice...))
}

// Append
func (s *Stream[T]) Append(elems []T) *Stream[T] {
	s.filteredSlice()
	s.slice = append(s.slice, elems...)
	return s
}

// Merge
func Merge[T any](streams ...*Stream[T]) *Stream[T] {
	lt := 0
	for _, s := range streams {
		lt += len(s.filteredSlice())
	}
	merged := make([]T, 0, lt)
	for _, s := range streams {
		merged = append(merged, s.slice...)
	}
	return From(merged)
}
