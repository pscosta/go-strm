package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapSum(t *testing.T) {
	// prepare
	initSlice := [][]int{{1}, {1, 2}, {1, 2, 3}}
	// call
	got := Map(
		From(initSlice),
		func(it []int) int { return Sum(From(it)) },
	).ToSlice()

	// assert
	assert.Equal(t, len(initSlice), len(got), "wrong length")
	assert.Equal(t, 1, got[0], "wrong sum")
	assert.Equal(t, 3, got[1], "wrong sum")
	assert.Equal(t, 6, got[2], "wrong sum")
}

func TestMapMin(t *testing.T) {
	// prepare
	initSlice := [][]int{{1}, {1, 2}, {3, 2, 1}}
	// call
	got := Map(
		From(initSlice),
		func(it []int) int { return Min(From(it)) },
	).ToSlice()

	// assert
	assert.Equal(t, len(initSlice), len(got), "wrong length")
	assert.Equal(t, 1, got[0], "wrong min")
	assert.Equal(t, 1, got[1], "wrong min")
	assert.Equal(t, 1, got[2], "wrong min")
}

func TestMapMax(t *testing.T) {
	// prepare
	initSlice := [][]int{{1}, {1, 2}, {1, 2, 3}}
	// call
	got := Map(
		From(initSlice),
		func(it []int) int { return Max(From(it)) },
	).ToSlice()

	// assert
	assert.Equal(t, len(initSlice), len(got), "wrong length")
	assert.Equal(t, 1, got[0], "wrong max")
	assert.Equal(t, 2, got[1], "wrong max")
	assert.Equal(t, 3, got[2], "wrong max")
}

func TestOnEach(t *testing.T) {
	// prepare
	initSlice := []int{1, 1, 1}
	var resCollector []int

	// call
	got := From(initSlice).
		OnEach(func(it int) { resCollector = append(resCollector, it+1) }).
		ToSlice()

	// assert
	assert.Equal(t, len(initSlice), len(resCollector), "wrong length")
	From(resCollector).ForEach(func(elem int) { assert.Equal(t, 2, elem, "wrong value") })
	From(got).ForEach(func(elem int) { assert.Equal(t, 1, elem, "wrong value") })
}

func TestApplyOnEach(t *testing.T) {
	// prepare
	initSlice := []int{1, 1, 1}

	// call
	got := From(initSlice).
		ApplyOnEach(func(it int) int { return it + 1 }).
		ToSlice()

	// assert
	assert.Equal(t, len(initSlice), len(got), "wrong length")
	From(got).ForEach(func(elem int) { assert.Equal(t, 2, elem, "wrong value") })
}

func TestReversed(t *testing.T) {
	// call
	reversedSlice := Of(1, 2, 3).Reversed().ToSlice()
	filteredSlice := Of(1, 2, 3).
		Filter(func(it int) bool { return it > 1 }).
		Reversed().
		ToSlice()

	// assert
	assert.Equal(t, 3, len(reversedSlice), "wrong length")
	assert.Equal(t, 2, len(filteredSlice), "wrong length")
	assert.Equal(t, 3, reversedSlice[0], "wrong value")
	assert.Equal(t, 2, reversedSlice[1], "wrong value")
	assert.Equal(t, 1, reversedSlice[2], "wrong value")
	assert.Equal(t, 3, filteredSlice[0], "wrong value")
	assert.Equal(t, 2, filteredSlice[1], "wrong value")
}

func TestDistinct(t *testing.T) {
	// call
	dedupedSlice := Of(1, 2, 3, 3).Distinct().ToSlice()
	filteredSlice := Of(1, 2, 3, 3).
		Filter(func(it int) bool { return it > 1 }).
		Distinct().
		ToSlice()

	// assert
	assert.Equal(t, 3, len(dedupedSlice), "wrong length")
	assert.Equal(t, 2, len(filteredSlice), "wrong length")
	assert.Equal(t, 1, dedupedSlice[0], "wrong value")
	assert.Equal(t, 2, dedupedSlice[1], "wrong value")
	assert.Equal(t, 3, dedupedSlice[2], "wrong value")
	assert.Equal(t, 2, filteredSlice[0], "wrong value")
	assert.Equal(t, 3, filteredSlice[1], "wrong value")
}

func TestDistinctStruct(t *testing.T) {
	// prepare
	type Person struct {
		name string
	}
	// call
	dedupedSlice := Of(Person{"Tim"}, Person{"Tim"}, Person{"Tom"}).Distinct().ToSlice()

	// assert
	assert.Equal(t, 2, len(dedupedSlice), "wrong length")
	assert.Equal(t, Person{"Tim"}, dedupedSlice[0], "wrong value")
	assert.Equal(t, Person{"Tom"}, dedupedSlice[1], "wrong value")
}

func TestDistinctSlice(t *testing.T) {
	// call
	dedupedSlice := Of([]int{1}, []int{1}, []int{1, 2}).Distinct().ToSlice()

	// assert
	assert.Equal(t, 2, len(dedupedSlice), "wrong length")
	assert.Equal(t, []int{1}, dedupedSlice[0], "wrong value")
}

func TestDistinctSeveralTypes(t *testing.T) {
	// prepare
	type Person struct {
		name string
		age  int
	}
	type DifficultPerson struct {
		age    int
		issues []string
	}

	// call
	emptySlice := Of(1).Filter(func(i int) bool { return false }).Distinct().ToSlice()
	Ints := Of(1, 2, 2, 2).Distinct().ToSlice()
	Ints2 := Of(1, 2, 3).Distinct().ToSlice()
	Structs := Of(Person{"Tim", 30}, Person{"Tim", 30}, Person{"Tom", 40}).Distinct().ToSlice()
	Strings := Of("Tim", "Tim", "Tom").Distinct().ToSlice()
	Slice := From([][]int{{1}, {1}, {1, 2}, {1, 2, 3}, {1, 2, 3}}).Distinct().ToSlice()
	SliceNil := From([][]int{{1}, {1}, {1, 2}, {1, 2, 3}, nil, nil}).Distinct().ToSlice()
	Slices := Of(DifficultPerson{39, []string{"covid"}}, DifficultPerson{39, []string{"covid"}}).Distinct().ToSlice()
	Slices2 := Of(DifficultPerson{39, []string{"covid"}}, DifficultPerson{39, []string{"covid", "19"}}).Distinct().ToSlice()
	Funs := Of(func(i int) {}, nil).Distinct().ToSlice()
	Maps := Of(map[int]int{1: 1}, map[int]int{1: 1}).Distinct().ToSlice()
	MapsNil := Of(map[int]int{1: 1}, map[int]int{1: 1}, nil).Distinct().ToSlice()
	arrays := Of([2]int{1: 1}, [2]int{0, 1}).Distinct().ToSlice()

	// assert
	assert.Equal(t, 0, len(emptySlice), "wrong Distinct size")
	assert.Equal(t, 2, len(Ints), "wrong Distinct size")
	assert.Equal(t, 3, len(Ints2), "wrong Distinct size")
	assert.Equal(t, 2, len(Structs), "wrong Distinct size")
	assert.Equal(t, 2, len(Strings), "wrong Distinct size")
	assert.Equal(t, 3, len(Slice), "wrong Distinct size")
	assert.Equal(t, 4, len(SliceNil), "wrong Distinct size")
	assert.Equal(t, 1, len(Slices), "wrong Distinct size")
	assert.Equal(t, 1, len(Slices2), "wrong Distinct size")
	assert.Equal(t, 2, len(Funs), "wrong Distinct size")
	assert.Equal(t, 1, len(Maps), "wrong Distinct size")
	assert.Equal(t, 2, len(MapsNil), "wrong Distinct size")
	assert.Equal(t, 1, len(arrays), "wrong Distinct size")
}

func TestChunkedOdd(t *testing.T) {
	// call
	batches := Of(1, 2, 3, 4, 5, 6, 7).Chunked(2)

	// assert
	assert.Equal(t, 4, len(batches), "wrong length")
	assert.Equal(t, 1, len(batches[3]), "wrong last batch length")
	From(batches).Take(3).ForEach(func(b []int) { assert.Equal(t, 2, len(b), "wrong batch length") })
}

func TestChunked(t *testing.T) {
	// call
	batches := Of(1, 2, 3, 4, 5, 6).Chunked(2)

	// assert
	assert.Equal(t, 3, len(batches), "wrong length")
	From(batches).ForEach(func(b []int) { assert.Equal(t, 2, len(b), "wrong batch length") })
}

func TestWindowed(t *testing.T) {
	// call
	windows := Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15).Windowed(5, 3)

	// assert
	assert.Equal(t, 4, len(windows), "wrong length")
	From(windows).ForEach(func(w []int) { assert.Equal(t, 5, len(w), "wrong window length") })
}

func TestWindowedPartial(t *testing.T) {
	// call
	windows := Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15).Windowed(5, 3, true)

	// assert
	assert.Equal(t, 5, len(windows), "wrong length")
	assert.Equal(t, 3, len(windows[4]), "wrong last window length")
	From(windows).Take(4).ForEach(func(w []int) { assert.Equal(t, 5, len(w), "wrong window length") })
}

func TestWindowedBigWindow(t *testing.T) {
	// call
	windows := Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15).Windowed(16, 3)

	// assert
	assert.Equal(t, 1, len(windows), "wrong length")
	assert.Equal(t, 15, len(windows[0]), "wrong length")
}

func TestTake(t *testing.T) {
	// call
	got := Of(0, 1, 2, 3).
		Take(2).
		ToSlice()

	// assert
	assert.Equal(t, 2, len(got), "wrong length")
	assert.Equal(t, 0, got[0], "wrong value")
	assert.Equal(t, 1, got[1], "wrong value")
}

func TestTakeNothing(t *testing.T) {
	// call
	got := Of(0, 1, 2, 3).
		Take(0).
		ToSlice()

	// assert
	assert.Equal(t, 0, len(got), "wrong length")
}

func TestTakeAll(t *testing.T) {
	// call
	got := Of(0, 1, 2, 3).
		Take(5).
		ToSlice()

	// assert
	if len(got) != 4 {
		t.Errorf("len(take(5)) = %d; want 4", len(got))
	}
}

func TestDrop(t *testing.T) {
	// call
	got := Of(0, 1, 2, 3).
		Drop(2).
		ToSlice()

	// assert
	assert.Equal(t, 2, len(got), "wrong length")
	assert.Equal(t, 2, got[0], "wrong value")
	assert.Equal(t, 3, got[1], "wrong value")
}

func TestEmptyDrop(t *testing.T) {
	// call
	got := Of(0, 1, 2, 3).
		Drop(5).
		ToSlice()

	// assert
	assert.Equal(t, 0, len(got), "wrong length")
}

func TestDropNothing(t *testing.T) {
	// call
	got := Of(0, 1, 2, 3).
		Drop(0).
		ToSlice()

	// assert
	assert.Equal(t, 4, len(got), "wrong length")
}
