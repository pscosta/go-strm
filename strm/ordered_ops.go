package strm

import (
	"golang.org/x/exp/constraints"
	"reflect"
)

// OnEach executes the given [action] on each element and returns the unchanged Stream afterwards.
// used mainly for debugging purposes.
func (s *Stream[T]) OnEach(f func(T)) *Stream[T] {
	for _, elem := range s.filteredSlice() {
		f(elem)
	}
	return s
}

// ApplyOnEach Applies the given [action] on each element of the backing slice returns the Stream afterwards.
// Similar to Map, but the returning type of the given [action] but math this Stream's type
func (s *Stream[T]) ApplyOnEach(action func(T) T) *Stream[T] {
	for i, elem := range s.filteredSlice() {
		s.slice[i] = action(elem)
	}
	return s
}

// Take Returns this Stream containing first [n] elements.
// [n] must be positive.
func (s *Stream[T]) Take(n int) *Stream[T] {
	if n == 0 || len(s.filteredSlice()) <= n {
		return s
	}
	s.slice = s.slice[:n]
	return s
}

// Drop Returns this Stream containing all elements except first [n] elements.
// [n] must be positive.
func (s *Stream[T]) Drop(n int) *Stream[T] {
	if n == 0 {
		return s
	}
	if len(s.filteredSlice()) <= n {
		// the nil slice is the preferred way
		s.slice = nil
	}
	s.slice = s.slice[n:]
	return s
}

// Max Returns the largest element of the Ordered constrained Stream.
func Max[O constraints.Ordered](s *Stream[O]) (max O) {
	max = s.First()
	for _, elem := range s.slice {
		if elem > max {
			max = elem
		}
	}
	return
}

// Min Returns the smallest element of the Ordered constrained Stream.
func Min[O constraints.Ordered](s *Stream[O]) (min O) {
	min = s.First()
	for _, elem := range s.slice {
		if elem < min {
			min = elem
		}
	}
	return
}

// Sum Returns the sum of all elements in this Ordered constrained Stream.
func Sum[O constraints.Ordered](s *Stream[O]) (sum O) {
	for _, elem := range s.filteredSlice() {
		sum += elem
	}
	return
}

// Reversed reverses the elements order of this Stream
func (s *Stream[T]) Reversed() *Stream[T] {
	for i := len(s.filteredSlice())/2 - 1; i >= 0; i-- {
		opp := len(s.slice) - 1 - i
		s.slice[i], (s.slice)[opp] = (s.slice)[opp], (s.slice)[i]
	}
	return s
}

// Distinct In-place deduplication of the backing slice, guided with a map.
// Streams of structs, pointers and primitive types are correctly de-duped
func (s *Stream[T]) Distinct() *Stream[T] {
	var keySelector func(t T) any

	// decides whether to compare pointers or values
	switch s.streamKind {
	case reflect.Array, reflect.Slice, reflect.Func, reflect.Map:
		// this is a hack to allow for key hashing over mutable/container types
		keySelector = func(t T) any { return &t }
	default:
		keySelector = func(t T) any { return t }
	}

	keys := make(map[any]struct{}, len(s.filteredSlice()))
	j := 0

	for i := 0; i < len(s.slice); i++ {
		if _, ok := keys[keySelector(s.slice[i])]; ok {
			continue
		}
		keys[keySelector(s.slice[i])] = struct{}{}
		s.slice[j], j = s.slice[i], j+1
	}
	s.slice = s.slice[:j]
	return s
}

// Contains Returns true if [element] is found in the Stream.
// Only works for Streams of structs, pointers and primitive types
func (s *Stream[T]) Contains(element T) bool {
	var valSelector func(t T) any

	switch s.streamKind {
	case reflect.Array, reflect.Slice, reflect.Func, reflect.Map:
		// always return false for mutable/container types
		return false
	default:
		valSelector = func(t T) any { return t }
	}

	// O(n) search
	for _, a := range s.filteredSlice() {
		if valSelector(a) == valSelector(element) {
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
