package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppend(t *testing.T) {
	// call
	slice1 := Of(1, 2).Append([]int{3}).ToSlice()

	// assert
	assert.Equal(t, 3, len(slice1), "wrong length")
	assert.Equal(t, 3, slice1[2], "wrong value")
}

func TestAppendEmpty(t *testing.T) {
	// call
	slice1 := Of(1, 2).Append([]int{}).ToSlice()

	// assert
	assert.Equal(t, 2, len(slice1), "wrong length")
}

func TestPlus(t *testing.T) {
	// call
	slice1 := Of(1, 2).Plus(Of(3)).ToSlice()

	// assert
	assert.Equal(t, 3, len(slice1), "wrong length")
	assert.Equal(t, 3, slice1[2], "wrong value")
}

func TestMerge(t *testing.T) {
	// call
	slice1 := Merge(Of(1, 2), Of(3), Of(4, 5)).ToSlice()

	// assert
	assert.Equal(t, 5, len(slice1), "wrong length")
	assert.Equal(t, 5, slice1[4], "wrong value")
}
