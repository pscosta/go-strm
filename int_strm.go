package main

import "sort"

// IntStream The Main struct
type IntStream struct {
	*Stream[int]
}

/*
 * Constructors
 */

// Range Returns a sequential ordered IntStream from
// [from] (inclusive) to [to] (inclusive) by an incremental step of 1.
func Range(from int, to int) *IntStream {
	if to < from {
		return &IntStream{From([]int{})}
	}
	intSlice := make([]int, 0, to-from)
	for i := from; i <= to; i++ {
		intSlice = append(intSlice, i)
	}
	return &IntStream{From(intSlice)}
}

// RangeOf Creates a new IntStream backed by the given [elems]
func RangeOf(elems ...int) *IntStream {
	intSlice := make([]int, 0, len(elems))
	for _, elem := range elems {
		intSlice = append(intSlice, elem)
	}
	return &IntStream{From(intSlice)}
}

// RangeFrom Creates a new IntStream backed by the given [backingSlice]
// the state of the given backingSlice will be updated by operation applied to the returned IntStream
func RangeFrom(backingIntSlice []int) *IntStream {
	return &IntStream{From(backingIntSlice)}
}

// RangeCopyFrom Creates a new IntStream backed by a copy of the elements in the given [slice]
// the state of the given slice will be preserved
func RangeCopyFrom(backingIntSlice []int) *IntStream {
	return &IntStream{CopyFrom(backingIntSlice)}
}

/*
 * Main Ops
 */

// Sum Returns the sum of elements in this IntStream
func (s *IntStream) Sum() (sum int) {
	if len(s.filteredSlice()) == 0 {
		return
	}
	for _, elem := range s.Stream.slice {
		sum += elem
	}
	return
}

// Min Returns the minimum element of this IntStream, or 0 if this IntStream is empty
func (s *IntStream) Min() int {
	if len(s.filteredSlice()) == 0 {
		return 0
	}
	min := s.Stream.slice[0]
	for _, elem := range s.Stream.slice {
		if elem < min {
			min = elem
		}
	}
	return min
}

// Max Returns the maximum element of this IntStream, or 0 if this IntStream is empty
func (s *IntStream) Max() int {
	if len(s.filteredSlice()) == 0 {
		return 0
	}
	max := s.Stream.slice[0]
	for _, elem := range s.Stream.slice {
		if elem > max {
			max = elem
		}
	}
	return max
}

// Avg Returns the arithmetic mean of elements of this IntStream, or 0 if this IntStream is empty
func (s *IntStream) Avg() int {
	if len(s.filteredSlice()) == 0 {
		return 0
	}
	return s.Sum() / len(s.slice)
}

// Sorted sorts the IntStream in increasing order.
func (s *IntStream) Sorted() *IntStream {
	if len(s.filteredSlice()) == 0 {
		return s
	}
	sort.Ints(s.Stream.slice)
	return s
}

/*
 * Adapter Ops
 */

// Filter see Stream.Filter
func (s *IntStream) Filter(p predicate[int]) *IntStream {
	s.Stream.Filter(p)
	return s
}

// ApplyOnEach see Stream.ApplyOnEach
func (s *IntStream) ApplyOnEach(action func(int) int) *IntStream {
	s.Stream.ApplyOnEach(action)
	return s
}

// OnEach see Stream.OnEach
func (s *IntStream) OnEach(action func(int)) *IntStream {
	s.Stream.OnEach(action)
	return s
}

// Distinct see Stream.Distinct
func (s *IntStream) Distinct() *IntStream {
	s.Stream.Distinct()
	return s
}

// Reversed see Stream.Reversed
func (s *IntStream) Reversed() *IntStream {
	s.Stream.Reversed()
	return s
}

// Take see Stream.Take
func (s *IntStream) Take(n int) *IntStream {
	s.Stream.Take(n)
	return s
}

// Drop see Stream.Drop
func (s *IntStream) Drop(n int) *IntStream {
	s.Stream.Drop(n)
	return s
}

// ToStrm Returns the enclosed *Stream[int] from this IntStream
func (s *IntStream) ToStrm() *Stream[int] {
	return s.Stream
}
