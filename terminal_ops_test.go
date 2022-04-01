package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToSlice(t *testing.T) {
	// prepare
	initSlice := []int{1, 2, 3}
	// call
	resSlice := From(initSlice).ToSlice()

	// assert
	assert.Equal(t, len(initSlice), len(resSlice), "should have same length")
	for i := range resSlice {
		assert.Equal(t, initSlice[i], resSlice[i], "should be equal")
	}
}

func TestForEach(t *testing.T) {
	// prepare
	initSlice := []int{1, 1, 1}
	var resCollector []int

	// call
	From(initSlice).ForEach(func(it int) { resCollector = append(resCollector, it+1) })

	// assert
	assert.Equal(t, 3, len(resCollector), "should have initial length")
	assert.Equal(t, len(initSlice), len(resCollector), "should have same length")
	for _, elem := range resCollector {
		assert.Equal(t, 2, elem, "should have been incremented in one unit")
	}
}

func TestAny(t *testing.T) {
	// prepare
	initSlice := []string{"Hi!", "Hello!", "Hey"}
	// call
	gotLen := From(initSlice).Any(func(it string) bool { return len(it) > 3 })
	gotMatch := From(initSlice).Any(func(it string) bool { return it == "Hey" })
	gotBye := From(initSlice).Any(func(it string) bool { return it == "Bye" })

	// assert
	assert.True(t, gotLen, "Any(len): expecting true")
	assert.True(t, gotMatch, "Any(Hey): expecting true")
	assert.False(t, gotBye, "Any(Bye): expecting false")
}

func TestAll(t *testing.T) {
	// prepare
	initSlice := []string{"Hi!", "Hello!", "Hey"}
	// call
	gotLen := From(initSlice).All(func(it string) bool { return len(it) >= 3 })
	gotMatch := From(initSlice).All(func(it string) bool { return it == "Hey" })
	gotLenFiltered := Of("Hi!", "Hello!", "Hey").
		Filter(func(it string) bool { return len(it) > 4 }).
		All(func(it string) bool { return it == "Hello!" })

	// assert
	assert.True(t, gotLen, "All: expecting true")
	assert.True(t, gotLenFiltered, "Any: expecting true")
	assert.False(t, gotMatch, "All: expecting false")
}

func TestNone(t *testing.T) {
	// prepare
	initSlice := []string{"Hi!", "Hello!", "Hey", ""}
	// call
	noneMatch := From(initSlice).None(func(it string) bool { return it == "Hey" })
	noneLen := From(initSlice).None(func(it string) bool { return len(it) < 3 })
	noneLenFiltered := Of("Hi!", "Hello!", "Hey", "").
		Filter(func(it string) bool { return len(it) > 0 }).
		None(func(it string) bool { return len(it) < 3 })

	// assert
	assert.True(t, noneLenFiltered, "None: expecting true")
	assert.False(t, noneLen, "None: expecting true")
	assert.False(t, noneMatch, "None: expecting false")
}

func TestCount(t *testing.T) {
	// prepare
	var slice2 []int
	// call
	count1 := Of(1, 2, 3).Filter(func(it int) bool { return it > 1 }).Count()
	count2 := From(slice2).Count()
	// assert
	assert.Equal(t, 2, count1, "wrong count")
	assert.Equal(t, 0, count2, "wrong count")
}

func TestCountBy(t *testing.T) {
	// prepare
	var slice2 []int
	// call
	count2 := From(slice2).CountBy(func(it int) bool { return it > 2 })
	count1 := Of(1, 2, 3).
		Filter(func(it int) bool { return it > 1 }).
		CountBy(func(it int) bool { return it > 2 })

	// assert
	assert.Equal(t, 1, count1, "wrong count")
	assert.Equal(t, 0, count2, "wrong count")
}

func TestSumBy(t *testing.T) {
	// prepare
	var slice2 []int
	// call
	sum1 := From(slice2).SumBy(func(it int) int { return it })
	sumBy := Of(1, 2, 3, 4, 5).
		SumBy(func(n int) int {
			if n%2 == 0 {
				return n
			} else {
				return 0
			}
		})

	// assert
	assert.Equal(t, 6, sumBy, "wrong sum")
	assert.Equal(t, 0, sum1, "wrong sum")
}

func TestFirst(t *testing.T) {
	// prepare
	var slice2 []int
	// call
	first1 := From(slice2).First()
	first2 := Of(1, 2, 3).
		Filter(func(it int) bool { return it > 1 }).
		First()

	// assert
	assert.Equal(t, 0, first1, "wrong first elem")
	assert.Equal(t, 2, first2, "wrong first elem")
}

func TestFirstBy(t *testing.T) {
	// prepare
	var slice []int
	// call
	first1 := From(slice).FirstBy(func(it int) bool { return it > 1 })
	first2 := Of(1, 2, 3).
		Filter(func(it int) bool { return it > 1 }).
		FirstBy(func(it int) bool { return it > 2 })

	// assert
	assert.Equal(t, 0, first1, "wrong first elem")
	assert.Equal(t, 3, first2, "wrong first elem")
}

func TestLast(t *testing.T) {
	// prepare
	var slice []int
	// call
	last1 := From(slice).Last()
	last2 := Of(1, 2, 3).
		Filter(func(it int) bool { return it > 1 }).
		Last()

	// assert
	assert.Equal(t, 0, last1, "wrong last elem")
	assert.Equal(t, 3, last2, "wrong last elem")
}

func TestJoinToString(t *testing.T) {
	// prepare
	var slice []int
	// call
	str1 := From(slice).JoinToString(",")
	str2 := Of(1, 2, 3).
		Filter(func(it int) bool { return it > 1 }).
		JoinToString("-")

	// assert
	assert.Equal(t, "", str1, "wrong string")
	assert.Equal(t, "2-3", str2, "wrong string")
}

func TestContains(t *testing.T) {
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
	hasInt := Of(1, 2, 3).Contains(2)
	hasInt2 := Of(1, 2, 3).Contains(4)
	hasStruct := Of(Person{"Tim", 30}, Person{"Tom", 40}).Contains(Person{"Tom", 40})
	hasString := Of("Tim", "Tom").Contains("tom")
	hasSlice := From([][]int{{1}, {1, 2}, {1, 2, 3}}).Contains([]int{1})
	hasSlices := Of(DifficultPerson{39, []string{"covid"}}).Contains(DifficultPerson{39, []string{"covid"}})
	hasFuns := Of(func(i int) {}).Contains(func(i int) {})
	hasMaps := Of(map[int]int{1: 1}).Contains(map[int]int{1: 1})
	hasArray := Of([2]int{1: 1}, [2]int{1: 2}).Contains([2]int{1: 1})
	hasArray2 := Of([2]int{1: 1}).Contains([2]int{0, 1})

	// assert
	assert.True(t, hasInt, "wrong Contains value")
	assert.False(t, hasInt2, "wrong Contains value")
	assert.True(t, hasStruct, "wrong Contains value")
	assert.False(t, hasString, "wrong Contains value")
	assert.True(t, hasSlice, "wrong Contains value")
	assert.True(t, hasSlices, "wrong Contains value")
	assert.False(t, hasFuns, "wrong Contains value")
	assert.True(t, hasMaps, "wrong Contains value")
	assert.True(t, hasArray, "wrong Contains value")
	assert.True(t, hasArray2, "wrong Contains value")
}
