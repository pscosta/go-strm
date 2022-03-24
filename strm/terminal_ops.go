package strm

import (
	"fmt"
	"reflect"
	"strings"
)

// ToSlice returns a Slice containing all the elements of this Stream
func (s *Stream[T]) ToSlice() []T {
	return s.filteredSlice()
}

// ForEach Performs the given action on each element of the Stream
func (s *Stream[T]) ForEach(action func(T)) {
	for _, elem := range s.filteredSlice() {
		action(elem)
	}
}

// Any Returns true if at least one element matches the given predicate.
func (s *Stream[T]) Any(p predicate[T]) bool {
	for _, elem := range s.filteredSlice() {
		if p(elem) {
			return true
		}
	}
	return false
}

// All Returns true if all elements match the given predicate.
func (s *Stream[T]) All(p predicate[T]) bool {
	for _, elem := range s.filteredSlice() {
		if !p(elem) {
			return false
		}
	}
	return true
}

// None Returns true if no elements match the given predicate.
func (s *Stream[T]) None(p predicate[T]) bool {
	return !s.Any(p)
}

// Count Returns the number of elements in this Stream
func (s *Stream[T]) Count() int {
	return len(s.filteredSlice())
}

// CountBy Returns the number of elements matching the given predicate [p].
func (s *Stream[T]) CountBy(p predicate[T]) (count int) {
	for _, elem := range s.filteredSlice() {
		if p(elem) {
			count++
		}
	}
	return
}

// SumBy Returns the sum of all values produced by [selector] function applied to each element in the Stream.
func (s *Stream[T]) SumBy(selector func(t T) int) (sum int) {
	for _, elem := range s.filteredSlice() {
		sum += selector(elem)
	}
	return
}

// FirstBy Returns the first element of this Stream matching the given predicate [p].
func (s *Stream[T]) FirstBy(p predicate[T]) (t T) {
	for _, elem := range s.filteredSlice() {
		if p(elem) {
			return elem
		}
	}
	return
}

// First Returns the first element of this Stream
func (s *Stream[T]) First() (t T) {
	if len(s.filteredSlice()) == 0 {
		return
	}
	return s.slice[0]
}

// Last Returns the last element of this Stream
func (s *Stream[T]) Last() (t T) {
	if len(s.filteredSlice()) == 0 {
		return
	}
	return s.slice[len(s.slice)-1]
}

// Contains Returns true if [element] is found in the Stream.
// Works for Streams of hashtable types like structs, pointers and primitive types
func (s *Stream[T]) Contains(element T) bool {
	var valSelector func(t T) any

	switch s.streamKind {
	case reflect.Array, reflect.Slice, reflect.Func, reflect.Map:
		// always return false for unhashable types
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

// JoinToString Creates a string from all the elements separated using [separator].
// makes use of fmt.Sprint() to dynamically convert the incoming generic type T to a string representation
func (s *Stream[T]) JoinToString(delimiter string) string {
	var sb strings.Builder
	for i := 0; i < len(s.filteredSlice()); i++ {
		sb.WriteString(fmt.Sprint(s.slice[i]))
		if i+1 < len(s.slice) {
			sb.WriteString(delimiter)
		}
	}
	return sb.String()
}
