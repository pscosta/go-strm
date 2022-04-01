package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRange(t *testing.T) {
	// call
	slice := Range(1, 3).ToSlice()
	// assert
	assert.Equal(t, 3, len(slice), "wrong length")
}

func TestRangeOf(t *testing.T) {
	// call
	slice := RangeOf(1, 2, 3).ToSlice()
	// assert
	assert.Equal(t, 3, len(slice), "wrong length")
}

func TestRangeFrom(t *testing.T) {
	// call
	slice := RangeFrom([]int{1, 2, 3}).ToSlice()
	// assert
	assert.Equal(t, 3, len(slice), "wrong length")
}

func TestFilterRangeCopyFrom(t *testing.T) {
	// prepare
	initSlice := []int{1, 2, 3}
	// call
	got := RangeCopyFrom(initSlice).
		Filter(func(it int) bool { return it >= 2 }).
		ToSlice()
	initSlice[0] = 3

	// assert
	assert.Equal(t, 2, len(got), "wrong length")
	assert.Equal(t, 3, len(initSlice), "wrong length")
	assert.Equal(t, 2, got[0], "wrong value")
}

func TestSum(t *testing.T) {
	// prepare
	initSlice := []int{1, 2, 3}
	// call
	sum := RangeFrom(initSlice).
		Sum()
	// assert
	assert.Equal(t, 6, sum, "wrong sum")
}

func TestMin(t *testing.T) {
	// call
	min := Range(1, 3).
		Min()
	// assert
	assert.Equal(t, 1, min, "wrong min")
}

func TestMax(t *testing.T) {
	// call
	max := Range(1, 3).
		Max()
	// assert
	assert.Equal(t, 3, max, "wrong max")
}

func TestAvg(t *testing.T) {
	// call
	max := RangeOf(1, 2, 3).
		Avg()
	// assert
	assert.Equal(t, 2, max, "wrong average")
}

func TestSorted(t *testing.T) {
	// call
	max := RangeOf(3, 2, 1).
		Sorted().
		ToSlice()
	// assert
	assert.Equal(t, []int{1, 2, 3}, max, "wrong sorting")
}
