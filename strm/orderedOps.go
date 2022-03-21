package strm

import "reflect"

func (s *Stream[T]) OnEach(f func(T)) *Stream[T] {
	for _, elem := range s.filteredSlice() {
		f(elem)
	}
	return s
}

func Max[O Ordered](s *Stream[O]) (max O) {
	max = s.First()
	for _, elem := range s.Slice {
		if elem > max {
			max = elem
		}
	}
	return
}

func Min[O Ordered](s *Stream[O]) (min O) {
	min = s.First()
	for _, elem := range s.Slice {
		if elem < min {
			min = elem
		}
	}
	return
}

func Sum[O Ordered](s *Stream[O]) (sum O) {
	for _, elem := range s.Slice {
		sum += elem
	}
	return
}

// Reversed reverses the backing Slice
func (s *Stream[T]) Reversed() *Stream[T] {
	for i := (len(s.filteredSlice()) / 2) - 1; i >= 0; i-- {
		opp := len(s.Slice) - 1 - i
		s.Slice[i], (s.Slice)[opp] = (s.Slice)[opp], (s.Slice)[i]
	}
	return s
}

// Distinct In-place deduplication of the backing Slice
func (s *Stream[T]) Distinct() *Stream[T] {
	var keySelector func(t T) any

	// decides whether to compare pointers or values
	switch s.streamType {
	case reflect.Array, reflect.Slice, reflect.Func, reflect.Map:
		keySelector = func(t T) any { return &t }
	default:
		keySelector = func(t T) any { return t }
	}

	keys := make(map[any]struct{}, len(s.filteredSlice()))
	j := 0

	for i := 0; i < len(s.Slice); i++ {
		if _, prs := keys[keySelector(s.Slice[i])]; prs {
			continue
		}
		keys[keySelector(s.Slice[i])] = struct{}{}
		s.Slice[j] = s.Slice[i]
		j++
	}
	s.Slice = s.Slice[:j]
	return s
}

// Contains Distinct In-place deduplication of the backing Slice
func (s *Stream[T]) Contains(t T) bool {
	var valSelector func(t T) any

	// always return false for container types
	switch s.streamType {
	case reflect.Array, reflect.Slice, reflect.Func, reflect.Map:
		return false
	default:
		valSelector = func(t T) any { return t }
	}

	// O(n) search
	for _, a := range s.filteredSlice() {
		if valSelector(a) == valSelector(t) {
			return true
		}
	}
	return false
}
