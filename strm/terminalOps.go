package strm

import (
	"fmt"
)

func (s *Stream[T]) ToSlice() []T {
	return s.filteredSlice()
}

func (s *Stream[T]) ForEach(f func(T)) {
	for _, elem := range s.filteredSlice() {
		f(elem)
	}
}

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
