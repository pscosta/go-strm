package strm

import (
	"fmt"
	"testing"
)

func TestOf(t *testing.T) {
	// call
	stream := Of(1, 2, 3)
	// assert
	if got := stream.ToSlice(); len(got) != 3 {
		t.Errorf("len(stream) = %d; want 3", len(got))
	}
}

func TestFrom(t *testing.T) {
	// call
	stream := From([][]int{{1}, {1, 2}, {1, 2, 3}})
	// assert
	if got := stream.ToSlice(); len(got) != 3 {
		t.Errorf("len(stream) = %d; want 3", len(got))
	}
}

func TestFilterCopyFrom(t *testing.T) {
	// prepare
	initSlice := [][]int{{1}, {1, 2}, {1, 2, 3}}
	// call
	got := CopyFrom(initSlice).
		Filter(func(it []int) bool { return len(it) > 2 }).
		ToSlice()
	// assert
	if len(got) != 1 {
		t.Errorf("len(initSlice) = %d; want 1", len(got))
	}
	if len(initSlice) != 3 {
		t.Errorf("len(initSlice) = %d; want 3", len(initSlice))
	}
	if len(got[0]) != 3 {
		t.Errorf("len(initSlice(0)) = %v; want 3", len(got[0]))
	}
}

func TestMapReduce(t *testing.T) {
	// prepare
	initSlice := [][]int{{1}, {1, 2}, {1, 2, 3}}

	// call
	got := Map(
		From(initSlice),
		func(it []int) int { return Reduce(From(it), func(a int, b int) int { return a + b }) },
	).ToSlice()

	// assert
	if len(got) != 3 {
		t.Errorf("len(mappedSlice) = %d; want 3", len(got))
	}
	if got[0] != 1 {
		t.Errorf("mappedSlice[0] = %d; want 1", got[0])
	}
	if got[1] != 3 {
		t.Errorf("mappedSlice[0] = %d; want 3", got[0])
	}
	if got[2] != 6 {
		t.Errorf("mappedSlice[0] = %d; want 6", got[0])
	}
}

func TestParallelLinearMapReduce(t *testing.T) {
	// prepare
	initSlice := [][]int{{1}, {1, 2}, {1, 2, 3}}

	// call
	got := PMap(
		From(initSlice),
		func(it []int) int { return Reduce(From(it), func(a int, b int) int { return a + b }) },
		true,
	).ToSlice()

	// assert
	if len(got) != 3 {
		t.Errorf("len(mappedSlice) = %d; want 3", len(got))
	}
	if got[0] != 1 {
		t.Errorf("mappedSlice[0] = %d; want 1", got[0])
	}
	if got[1] != 3 {
		t.Errorf("mappedSlice[0] = %d; want 3", got[0])
	}
	if got[2] != 6 {
		t.Errorf("mappedSlice[0] = %d; want 6", got[0])
	}
}

func TestParallelBatchedMapReduce(t *testing.T) {
	// prepare
	initSlice := [][]int{{1}, {1, 2}, {1, 2, 3}}

	// call
	got := PMap(
		From(initSlice),
		func(it []int) int { return Reduce(From(it), func(a int, b int) int { return a + b }) },
	).ToSlice()

	// assert
	if len(got) != 3 {
		t.Errorf("len(mappedSlice) = %d; want 3", len(got))
	}
	if got[0] != 1 {
		t.Errorf("mappedSlice[0] = %d; want 1", got[0])
	}
	if got[1] != 3 {
		t.Errorf("mappedSlice[0] = %d; want 3", got[0])
	}
	if got[2] != 6 {
		t.Errorf("mappedSlice[0] = %d; want 6", got[0])
	}
}

func TestMapReduceWithStart(t *testing.T) {
	// prepare
	initSlice := [][]int{{1}, {1, 2}, {1, 2, 3}}

	// call
	got := Map(
		From(initSlice),
		func(it []int) int { return Reduce(From(it), func(a int, b int) int { return a + b }, 1) },
	).ToSlice()

	// assert
	if len(got) != 3 {
		t.Errorf("len(mappedSlice) = %d; want 3", len(got))
	}
	if got[0] != 2 {
		t.Errorf("mappedSlice[0] = %d; want 1", got[0])
	}
	if got[1] != 4 {
		t.Errorf("mappedSlice[0] = %d; want 3", got[1])
	}
	if got[2] != 7 {
		t.Errorf("mappedSlice[0] = %d; want 6", got[2])
	}
}

func TestFlatMap(t *testing.T) {
	// call
	got := FlatMap(
		Of(1, 2, 4),
		func(e int) *Stream[string] { return Of(fmt.Sprint(e)) },
	).ToSlice()

	// assert
	if len(got) != 3 {
		t.Errorf("len(FlatMap) = %d; want 7", len(got))
	}
	if got[0] != "1" {
		t.Errorf("mappedSlice[0] = %v; want 1", got[0])
	}
	if got[1] != "2" {
		t.Errorf("mappedSlice[0] = %v; want 3", got[1])
	}
	if got[2] != "4" {
		t.Errorf("mappedSlice[0] = %v; want 6", got[2])
	}
}

func TestGroupBy(t *testing.T) {
	// prepare
	type Person struct {
		name string
		age  int
	}
	stream := Of(Person{"Tim", 30}, Person{"Bil", 40}, Person{"John", 30}, Person{"Tim", 35})

	// call
	byAge := GroupBy(stream, func(it Person) int { return it.age })
	byName := GroupBy(stream, func(it Person) string { return it.name })

	// assert
	if len(byAge) != 3 {
		t.Errorf("len(byAge) = %v; want 3", len(byAge))
	}
	if len(byName) != 3 {
		t.Errorf("len(byAge) = %v; want 3", len(byName))
	}
	if byAge[30][0].age != 30 {
		t.Errorf("byAge[30][0].age = %v; want 3", byAge[30][0].age)
	}
	if byAge[30][0].name != "Tim" {
		t.Errorf("byAge[30][0].name = %v; want 3", byAge[30][0].name)
	}
	if byAge[30][1].age != 30 {
		t.Errorf("byAge[30][1].age = %v; want 3", byAge[30][1].age)
	}
	if byAge[30][1].name != "John" {
		t.Errorf("len(byAge) = %v; want 3", byAge[30][1].name)
	}
	if byName["Bil"][0].age != 40 {
		t.Errorf("byName[\"Bil\"][0].age = %v; want 3", byName["Bil"][0].age)
	}
	if byName["Bil"][0].name != "Bil" {
		t.Errorf("byName[\"Bill\"][0].name = %v; want 3", byName["Bill"][0].name)
	}
}
