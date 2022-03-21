package strm

import "reflect"

func (s *Stream[T]) OnEach(f func(T)) *Stream[T] {
	for _, elem := range s.filteredSlice() {
		f(elem)
	}
	return s
}

func Max[O ordered](s *Stream[O]) (max O) {
	max = s.First()
	for _, elem := range s.slice {
		if elem > max {
			max = elem
		}
	}
	return
}

func Min[O ordered](s *Stream[O]) (min O) {
	min = s.First()
	for _, elem := range s.slice {
		if elem < min {
			min = elem
		}
	}
	return
}

func Sum[O ordered](s *Stream[O]) (sum O) {
	for _, elem := range s.slice {
		sum += elem
	}
	return
}

// Reversed reverses the backing slice
func (s *Stream[T]) Reversed() *Stream[T] {
	s.filteredSlice()
	for i := len(s.slice)/2 - 1; i >= 0; i-- {
		opp := len(s.slice) - 1 - i
		s.slice[i], (s.slice)[opp] = (s.slice)[opp], (s.slice)[i]
	}
	return s
}

// Distinct In-place deduplication of the backing slice
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

	for i := 0; i < len(s.slice); i++ {
		if _, prs := keys[keySelector(s.slice[i])]; prs {
			continue
		}
		keys[keySelector(s.slice[i])] = struct{}{}
		s.slice[j] = s.slice[i]
		j++
	}
	s.slice = s.slice[:j]
	return s
}

// Contains Distinct In-place deduplication of the backing slice
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
