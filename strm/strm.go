package strm

import (
	"reflect"
	"runtime"
)

// internal types
type predicate[T any] func(v T) bool
type mapper[IN any, OUT any] func(v IN) OUT
type reducer[OUT any, IN any] func(OUT, IN) OUT

// Stream The Main struct
type Stream[T any] struct {
	slice      []T
	filters    []predicate[T]
	streamKind reflect.Kind
}

// Constructors

// From Creates a new Stream backed by the given [backingSlice]
// the state of the given backingSlice will be updated by operation applied to the returned Stream
func From[T any](backingSlice []T) *Stream[T] {
	return &Stream[T]{
		slice:      backingSlice,
		streamKind: typeOf[T](),
	}
}

// CopyFrom Creates a new Stream backed by a copy of the elements in the given [slice]
// the state of the given slice will be preserved
func CopyFrom[T any](slice []T) *Stream[T] {
	sliceCopy := make([]T, len(slice))
	copy(sliceCopy, slice)

	return &Stream[T]{
		slice:      sliceCopy,
		streamKind: reflect.TypeOf((*T)(nil)).Elem().Kind(),
	}
}

// Of Creates a new Stream backed by the given [elems]
func Of[T any](elems ...T) *Stream[T] {
	slice := make([]T, 0, len(elems))

	for _, elem := range elems {
		slice = append(slice, elem)
	}
	return From(slice)
}

// Main functions

// Filter Returns a Stream containing only elements matching the given [predicate].
// This operation is lazy and will be applied only upon calling a terminal operation on the Stream
func (s *Stream[T]) Filter(p predicate[T]) *Stream[T] {
	s.filters = append(s.filters, p)
	return s
}

// Map Returns a new Stream containing the results of applying the given function to each element in the given Stream
func Map[IN any, OUT any](s *Stream[IN], f mapper[IN, OUT]) *Stream[OUT] {
	var newSlice []OUT

	for _, elem := range s.filteredSlice() {
		newSlice = append(newSlice, f(elem))
	}
	return From(newSlice)
}

// PMap Returns a new Stream containing the results of applying the given function to each element in the given Stream
// in parallel. By default, the parallel work is batched by number of available CPU cores.
// If the [noBatching] flag is present, PMap will launch a new goroutine per each element present in the provided Stream
// - not recommended for very large Streams due to potentially large memory footprint.
func PMap[IN any, OUT any](s *Stream[IN], f mapper[IN, OUT], noBatching ...bool) *Stream[OUT] {
	if len(noBatching) == 0 {
		return parallelBatchingMap(s, f)
	} else {
		return parallelLinearMap(s, f)
	}
}

// FlatMap Returns a single Stream of all elements yielded from results of [mapper] function
// being invoked on each element of original Stream
func FlatMap[IN any, OUT any](s *Stream[IN], f mapper[IN, *Stream[OUT]]) *Stream[OUT] {
	var newSlice []OUT

	for _, elem := range s.filteredSlice() {
		for _, slice := range f(elem).filteredSlice() {
			newSlice = append(newSlice, slice)
		}
	}
	return From(newSlice)
}

// Reduce Accumulates value starting with the given [start] value if provided, or with first element and applying
// the given [reducer] operation from left to right to current accumulator value and each element.
//	operation: function that takes current accumulator value and an element, and calculates the next accumulator value.
func Reduce[IN any, OUT any](s *Stream[IN], f reducer[OUT, IN], start ...OUT) (out OUT) {
	if len(start) > 0 {
		out = start[0]
	}
	for _, elem := range s.filteredSlice() {
		out = f(out, elem)
	}
	return out
}

// GroupBy Groups elements of the given Stream by the key produced by the given [keySelector] applied to each element
// and returns a map where each group key is associated with a slice of corresponding elements.
// The returned map preserves the entry iteration order of the keys produced from the original Stream.
func GroupBy[K comparable, V any](s *Stream[V], keySelector func(V) K) map[K][]V {
	grouping := make(map[K][]V, len(s.filteredSlice()))

	for _, elem := range s.slice {
		key := keySelector(elem)
		grouping[key] = append(grouping[key], elem)
	}
	return grouping
}

// Internal Ops

// returns the King of the given generic type
func typeOf[T any]() reflect.Kind {
	return reflect.TypeOf((*T)(nil)).Elem().Kind()
}

// returns the filtered backing slice after applying all registered filters
func (s *Stream[T]) filteredSlice() []T {
	applyFilters(&(s.slice), s.filters)
	s.filters = nil
	return s.slice
}

// Applies all lazy filters to the Stream
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
		(*slice)[i], i = elem, i+1
	}
	// garbage-collects the unfiltered elements
	for j := i; j < len(*slice); j++ {
		var t T
		(*slice)[j] = t
	}
	*slice = (*slice)[:i]
}

// parallelLinearMap Returns a new Stream containing the results of applying the given function to each element
// in the given Stream in parallel. A new goroutine is launched per each element present in the provided Stream.
func parallelLinearMap[IN any, OUT any](s *Stream[IN], f mapper[IN, OUT]) *Stream[OUT] {
	resultSlice := make([]OUT, len(s.filteredSlice()))
	ch := make(chan struct{}, len(s.slice))

	// launching the goroutines
	for i := range s.slice {
		// launch goroutine that executes the mapping asynchronously
		go func(f mapper[IN, OUT], idx int, ch chan struct{}) {
			resultSlice[idx] = f(s.slice[idx])
			ch <- struct{}{} // signals completion
		}(f, i, ch)
	}
	// blocking: waits for all goroutines to complete
	for range s.slice {
		<-ch
	}
	return From(resultSlice)
}

// parallelBatchingMap Returns a new Stream containing the results of applying the given function to each element
// in the given Stream in parallel. The parallel work is batched by number of available CPU cores.
func parallelBatchingMap[IN any, OUT any](s *Stream[IN], f mapper[IN, OUT]) *Stream[OUT] {
	resultSlice := make([]OUT, len(s.filteredSlice()))
	cores := Min(Of(runtime.NumCPU(), len(s.slice)))
	ch := make(chan struct{}, cores)
	batchSize := Max(Of(len(s.slice)/cores, 1))

	// launching the goroutines
	for i := range s.slice {
		// launch goroutine that executes the batch mapping asynchronously
		go func(f mapper[IN, OUT], idx int, ch chan struct{}) {
			for j := idx; j < idx+batchSize && j < len(s.slice); j++ {
				resultSlice[j] = f(s.slice[j])
			}
			ch <- struct{}{} // signals completion
		}(f, i, ch)
	}
	// blocking: waits for all goroutines to complete
	for range s.slice {
		<-ch
	}
	return From(resultSlice)
}
