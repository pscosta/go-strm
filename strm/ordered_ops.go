package strm

import "reflect"

func (s *Stream[T]) OnEach(f func(T)) *Stream[T] {
	for _, elem := range s.filteredSlice() {
		f(elem)
	}
	return s
}

func (s *Stream[T]) ApplyOnEach(f func(T) T) *Stream[T] {
	for i, elem := range s.filteredSlice() {
		s.slice[i] = f(elem)
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

// Contains In-place deduplication of the backing slice
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

// Chunked Splits this Stream into several slices each not exceeding the given [size]
// The last list may have fewer elements than the given [size].
//	 size: the nr. of elems to take in each slice, must be >0 and can be greater than the nr of elems in this stream
func (s *Stream[T]) Chunked(batchSize int) [][]T {
	batches := make([][]T, 0, (len(s.filteredSlice())+batchSize-1)/batchSize)

	for batchSize < len(s.slice) {
		s.slice, batches = s.slice[batchSize:], append(batches, s.slice[0:batchSize:batchSize])
	}
	return append(batches, s.slice)
}

// Windowed Returns a slice of slices of the window of the given size, sliding along this Stream with the given [step].
// Several last slices may have fewer elements than the given [size].
// Both [size] and [step] must be positive and can be greater than the number of elements in this Stream.
//	 size: the number of elements to take in each window
// 	 step: the number of elements to move the window forward by on each step
//	 partialWindows: controls whether to keep partial windows in the end if any, false by default
func (s *Stream[T]) Windowed(size int, step int, partialWindows ...bool) [][]T {
	// returns the input slice as the first element
	if len(s.filteredSlice()) <= size {
		return [][]T{s.slice}
	}
	// no partial windows by default
	partialWin := false
	if len(partialWindows) > 0 {
		partialWin = partialWindows[0]
	}
	// allocate slice at the requested
	res := make([][]T, 0, len(s.slice)/size+1)

	winSize := step
	for i, j := 0, size; j <= len(s.slice) && i <= len(s.slice) && i < j; i, j = i+step, j+winSize {
		res = append(res, s.slice[i:j])
		if (len(s.slice)-j) < winSize && partialWin {
			winSize, j = 0, len(s.slice)
		}
	}
	return res
}
