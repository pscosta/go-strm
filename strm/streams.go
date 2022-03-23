package strm

import (
	"reflect"
)

type predicate[T any] func(v T) bool
type mapper[IN any, OUT any] func(v IN) OUT
type reducer[OUT any, IN any] func(OUT, IN) OUT

type Stream[T any] struct {
	slice      []T
	filters    []predicate[T]
	streamType reflect.Kind
}

// Constructors

func From[T any](backingSlice []T) *Stream[T] {
	return &Stream[T]{
		slice:      backingSlice,
		streamType: typeOf[T](),
	}
}

func Of[T any](elems ...T) *Stream[T] {
	slice := make([]T, 0, len(elems))

	for _, elem := range elems {
		slice = append(slice, elem)
	}
	return From(slice)
}

func CopyFrom[T any](slice []T) *Stream[T] {
	copySlice := make([]T, len(slice))
	copy(copySlice, slice)

	return &Stream[T]{
		slice:      copySlice,
		streamType: reflect.TypeOf((*T)(nil)).Elem().Kind(),
	}
}

// Main functions

func (s *Stream[T]) Filter(p predicate[T]) *Stream[T] {
	s.filters = append(s.filters, p)
	return s
}

func Map[IN any, OUT any](s *Stream[IN], f mapper[IN, OUT]) *Stream[OUT] {
	var newSlice []OUT

	for _, elem := range s.filteredSlice() {
		newSlice = append(newSlice, f(elem))
	}
	return From(newSlice)
}

func FlatMap[IN any, OUT any](s *Stream[IN], f mapper[IN, *Stream[OUT]]) *Stream[OUT] {
	var newSlice []OUT

	for _, elem := range s.filteredSlice() {
		for _, slice := range f(elem).filteredSlice() {
			newSlice = append(newSlice, slice)
		}
	}
	return From(newSlice)
}

func Reduce[IN any, OUT any](s *Stream[IN], f reducer[OUT, IN], start ...OUT) (out OUT) {
	if len(start) > 0 {
		out = start[0]
	}
	for _, elem := range s.filteredSlice() {
		out = f(out, elem)
	}
	return out
}

func GroupBy[K comparable, V any](s *Stream[V], keySelector func(V) K) map[K][]V {
	grouping := make(map[K][]V, len(s.filteredSlice()))

	for _, elem := range s.slice {
		key := keySelector(elem)
		grouping[key] = append(grouping[key], elem)
	}
	return grouping
}

// Internal Ops

func typeOf[T any]() reflect.Kind {
	return reflect.TypeOf((*T)(nil)).Elem().Kind()
}

func (s *Stream[T]) filteredSlice() []T {
	applyFilters(&(s.slice), s.filters)
	s.filters = nil
	return s.slice
}

func applyFilters[T any](slice *[]T, filters []predicate[T]) {
	if len(filters) == 0 {
		return
	}
	i := 0
filtering:
	for _, elem := range *slice {
		for _, filter := range filters {
			if !filter(elem) {
				continue filtering
			}
		}
		(*slice)[i] = elem
		i++
	}
	// garbage-collects the unfiltered elements
	for j := i; j < len(*slice); j++ {
		var t T
		(*slice)[j] = t
	}
	*slice = (*slice)[:i]
}
