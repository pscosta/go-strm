package strm

import (
	"reflect"
)

type predicate[T any] func(v T) bool
type mapper[IN any, OUT any] func(v IN) OUT
type reducer[OUT any, IN any] func(OUT, IN) OUT
type ordered interface{ int | int64 | float32 | string }

type Stream[T any] struct {
	Slice      []T
	filters    []predicate[T]
	streamType reflect.Kind
}

// Constructors

func From[T any](backingSlice []T) *Stream[T] {
	return &Stream[T]{
		Slice:      backingSlice,
		streamType: reflect.TypeOf((*T)(nil)).Elem().Kind(),
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
		Slice:      copySlice,
		streamType: reflect.TypeOf((*T)(nil)).Elem().Kind(),
	}
}

// Main functions

func (s *Stream[T]) Filter(f predicate[T]) *Stream[T] {
	s.filters = append(s.filters, f)
	return s
}

func Map[IN any, OUT any](s *Stream[IN], f mapper[IN, OUT]) *Stream[OUT] {
	var newSlice []OUT

	for _, elem := range s.filteredSlice() {
		newSlice = append(newSlice, f(elem))
	}
	return From(newSlice)
}

func Reduce[IN any, OUT any](s *Stream[IN], start OUT, f reducer[OUT, IN]) OUT {
	out := start
	for _, elem := range s.Slice {
		out = f(out, elem)
	}
	return out
}

func GroupBy[K comparable, V any](s *Stream[V], keySelector func(V) K) map[K][]V {
	grouping := make(map[K][]V, len(s.filteredSlice()))

	for _, elem := range s.Slice {
		key := keySelector(elem)
		grouping[key] = append(grouping[key], elem)
	}
	return grouping
}

// Internal Ops

func (s *Stream[T]) filteredSlice() []T {
	applyFilters(&(s.Slice), s.filters)
	s.filters = nil
	return s.Slice
}

func applyFilters[T any](slice *[]T, filters []predicate[T]) {
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
