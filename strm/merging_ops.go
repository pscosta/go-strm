package strm

// Plus copies the backing slices contents of both Streams into a new *Stream[T]
func (s *Stream[T]) Plus(other *Stream[T]) *Stream[T] {
	merged := make([]T, len(s.filteredSlice()), len(s.slice)+len(other.filteredSlice()))
	copy(merged, s.slice)
	return From(append(merged, other.slice...))
}

// Append appends the element of the given slice to this Stream
func (s *Stream[T]) Append(elems []T) *Stream[T] {
	s.filteredSlice()
	s.slice = append(s.slice, elems...)
	return s
}

// Merge merges the given Streams into a single one
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
