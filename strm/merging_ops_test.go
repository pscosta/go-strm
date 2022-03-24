package strm

import (
	"testing"
)

func TestAppend(t *testing.T) {
	// call
	slice1 := Of(1, 2).Append([]int{3}).ToSlice()

	// assert
	if len(slice1) != 3 {
		t.Errorf("len(Append) = %d; want 3", len(slice1))
	}
	if slice1[2] != 3 {
		t.Errorf("Append[2] = %d; want 3", slice1[0])
	}
}

func TestAppendEmpty(t *testing.T) {
	// call
	slice1 := Of(1, 2).Append([]int{}).ToSlice()

	// assert
	if len(slice1) != 2 {
		t.Errorf("len(Append) = %d; want 3", len(slice1))
	}
}

func TestPlus(t *testing.T) {
	// call
	slice1 := Of(1, 2).Plus(Of(3)).ToSlice()

	// assert
	if len(slice1) != 3 {
		t.Errorf("len(Plus) = %d; want 3", len(slice1))
	}
	if slice1[2] != 3 {
		t.Errorf("Plus[2] = %d; want 3", slice1[2])
	}
}

func TestMerge(t *testing.T) {
	// call
	slice1 := Merge(Of(1, 2), Of(3), Of(4, 5)).ToSlice()

	// assert
	if len(slice1) != 5 {
		t.Errorf("len(Merged) = %d; want 5", len(slice1))
	}
	if slice1[4] != 5 {
		t.Errorf("Merged[4] = %d; want 5", slice1[2])
	}
}
