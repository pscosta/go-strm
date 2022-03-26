package strm

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOf(t *testing.T) {
	// call
	slice := Of(1, 2, 3).ToSlice()
	// assert
	assert.Equal(t, 3, len(slice), "wrong length")
}

func TestFrom(t *testing.T) {
	// call
	slice := From([][]int{{1}, {1, 2}, {1, 2, 3}}).ToSlice()
	// assert
	assert.Equal(t, 3, len(slice), "wrong length")
}

func TestFilterCopyFrom(t *testing.T) {
	// prepare
	initSlice := [][]int{{1}, {1, 2}, {1, 2, 3}}
	// call
	got := CopyFrom(initSlice).
		Filter(func(it []int) bool { return len(it) > 2 }).
		ToSlice()
	initSlice[0] = []int{1, 2}

	// assert
	assert.Equal(t, 1, len(got), "wrong length")
	assert.Equal(t, 3, len(initSlice), "wrong length")
	assert.Equal(t, 3, len(got[0]), "wrong length")
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
	assert.Equal(t, 3, len(got), "wrong length")
	assert.Equal(t, 1, got[0], "wrong value")
	assert.Equal(t, 3, got[1], "wrong value")
	assert.Equal(t, 6, got[2], "wrong value")
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
	assert.Equal(t, 3, len(got), "wrong length")
	assert.Equal(t, 1, got[0], "wrong value")
	assert.Equal(t, 3, got[1], "wrong value")
	assert.Equal(t, 6, got[2], "wrong value")
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
	assert.Equal(t, 3, len(got), "wrong length")
	assert.Equal(t, 1, got[0], "wrong value")
	assert.Equal(t, 3, got[1], "wrong value")
	assert.Equal(t, 6, got[2], "wrong value")
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
	assert.Equal(t, 3, len(got), "wrong length")
	assert.Equal(t, 2, got[0], "wrong value")
	assert.Equal(t, 4, got[1], "wrong value")
	assert.Equal(t, 7, got[2], "wrong value")
}

func TestFlatMap(t *testing.T) {
	// call
	got := FlatMap(
		Of(1, 2, 4),
		func(e int) *Stream[string] { return Of(fmt.Sprint(e)) },
	).ToSlice()

	// assert
	assert.Equal(t, 3, len(got), "wrong length")
	assert.Equal(t, "1", got[0], "wrong value")
	assert.Equal(t, "2", got[1], "wrong value")
	assert.Equal(t, "4", got[2], "wrong value")
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
	assert.Equal(t, 3, len(byAge), "wrong length")
	assert.Equal(t, 3, len(byName), "wrong length")
	assert.Equal(t, 30, byAge[30][0].age, "wrong value")
	assert.Equal(t, "Tim", byAge[30][0].name, "wrong value")
	assert.Equal(t, 30, byAge[30][1].age, "wrong value")
	assert.Equal(t, "John", byAge[30][1].name, "wrong value")
	assert.Equal(t, 40, byName["Bil"][0].age, "wrong value")
	assert.Equal(t, "Bil", byName["Bil"][0].name, "wrong value")
}
