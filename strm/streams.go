package strm

import (
	"fmt"
	"reflect"
)

type Predicate[T any] func(v T) bool
type Mapper[IN any, OUT any] func(v IN) OUT
type Reducer[OUT any, IN any] func(OUT, IN) OUT

type Stream[T any] struct {
	Slice      []T
	filters    []Predicate[T]
	streamType reflect.Kind
}

// Internal methods

func (s *Stream[T]) filteredSlice() []T {
	applyFilters(&(s.Slice), s.filters)
	s.filters = nil
	return s.Slice
}

func applyFilters[T any](slice *[]T, filters []Predicate[T]) {
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

// initializers

func From[T any](slice []T) *Stream[T] {
	return &Stream[T]{
		Slice:      slice,
		streamType: reflect.TypeOf((*T)(nil)).Elem().Kind(),
	}
}

func Of[T any](elems ...T) *Stream[T] {
	slice := make([]T, 0, len(elems))

	for _, elem := range elems {
		slice = append(slice, elem)
	}

	return &Stream[T]{
		Slice:      slice,
		streamType: reflect.TypeOf((*T)(nil)).Elem().Kind(),
	}
}

// Main functions

func (s *Stream[T]) Filter(f Predicate[T]) *Stream[T] {
	s.filters = append(s.filters, f)
	return s
}

func Map[IN any, OUT any](s *Stream[IN], f Mapper[IN, OUT]) *Stream[OUT] {
	s.filteredSlice()
	var newSlice []OUT

	for _, elem := range s.Slice {
		newSlice = append(newSlice, f(elem))
	}

	return From(newSlice)
}

func Reduce[IN any, OUT any](s *Stream[IN], start OUT, f Reducer[OUT, IN]) OUT {
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

func (s *Stream[T]) OnEach(f func(T)) *Stream[T] {
	for _, elem := range s.filteredSlice() {
		f(elem)
	}
	return s
}

func Max[C int](s *Stream[C]) (max C) {
	if len(s.filteredSlice()) == 0 {
		return
	}
	max = s.Slice[0]
	for _, elem := range s.Slice {
		if elem > max {
			max = elem
		}
	}
	return
}

func Min[C int](s *Stream[C]) (min C) {
	if len(s.filteredSlice()) == 0 {
		return
	}
	min = s.Slice[0]
	for _, elem := range s.Slice {
		if elem < min {
			min = elem
		}
	}
	return
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

// Reversed reverses the backing Slice
func (s *Stream[T]) Reversed() *Stream[T] {
	for i := (len(s.filteredSlice()) / 2) - 1; i >= 0; i-- {
		opp := len(s.Slice) - 1 - i
		s.Slice[i], (s.Slice)[opp] = (s.Slice)[opp], (s.Slice)[i]
	}
	return s
}

// Terminal operations

func (s *Stream[T]) Any(p Predicate[T]) bool {
	for _, elem := range s.filteredSlice() {
		if p(elem) {
			return true
		}
	}
	return false
}

func (s *Stream[T]) All(p Predicate[T]) bool {
	for _, elem := range s.filteredSlice() {
		if !p(elem) {
			return false
		}
	}
	return true
}

func (s *Stream[T]) None(p Predicate[T]) bool {
	return !s.Any(p)
}

func (s *Stream[T]) Count() int {
	return len(s.filteredSlice())
}

func (s *Stream[T]) CountBy(p Predicate[T]) (count int) {
	for _, elem := range s.filteredSlice() {
		if p(elem) {
			count++
		}
	}
	return
}

func (s *Stream[T]) FirstBy(p Predicate[T]) (t T) {
	for _, elem := range s.filteredSlice() {
		if p(elem) {
			return elem
		}
	}
	return
}

func (s *Stream[T]) ForEach(f func(T)) {
	for _, elem := range s.filteredSlice() {
		f(elem)
	}
}

func (s *Stream[T]) First() (t T) {
	if len(s.filteredSlice()) == 0 {
		return
	}
	return s.Slice[0]
}

func (s *Stream[T]) Last() (t T) {
	if len(s.filteredSlice()) == 0 {
		return
	}
	return s.filteredSlice()[len(s.Slice)-1]
}

func (s *Stream[T]) JoinToString(delimiter string) (joined string) {
	s.filteredSlice()
	for i := 0; i < len(s.Slice); i++ {
		joined += fmt.Sprint(s.Slice[i])
		if i+1 < len(s.Slice) {
			joined += delimiter
		}
	}
	return
}

func (s *Stream[T]) ToSlice() []T {
	return s.filteredSlice()
}
