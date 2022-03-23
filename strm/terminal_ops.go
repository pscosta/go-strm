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

func (s *Stream[T]) Any(p predicate[T]) bool {
	for _, elem := range s.filteredSlice() {
		if p(elem) {
			return true
		}
	}
	return false
}

func (s *Stream[T]) All(p predicate[T]) bool {
	for _, elem := range s.filteredSlice() {
		if !p(elem) {
			return false
		}
	}
	return true
}

func (s *Stream[T]) None(p predicate[T]) bool {
	return !s.Any(p)
}

func (s *Stream[T]) Count() int {
	return len(s.filteredSlice())
}

func (s *Stream[T]) CountBy(p predicate[T]) (count int) {
	for _, elem := range s.filteredSlice() {
		if p(elem) {
			count++
		}
	}
	return
}

func (s *Stream[T]) FirstBy(p predicate[T]) (t T) {
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
	return s.slice[0]
}

func (s *Stream[T]) Last() (t T) {
	if len(s.filteredSlice()) == 0 {
		return
	}
	return s.slice[len(s.slice)-1]
}

func (s *Stream[T]) JoinToString(delimiter string) (joined string) {
	for i := 0; i < len(s.filteredSlice()); i++ {
		joined += fmt.Sprint(s.slice[i])
		if i+1 < len(s.slice) {
			joined += delimiter
		}
	}
	return
}
