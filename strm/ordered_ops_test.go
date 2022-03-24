package strm

import (
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
	if len(got) != 3 {
		t.Errorf("len(mappedStream) = %d; want 3", len(got))
	}
	if got[0] != 1 {
		t.Errorf("mappedStream[0] = %d; want 1", got[0])
	}
	if got[1] != 3 {
		t.Errorf("mappedStream[0] = %d; want 3", got[0])
	}
	if got[2] != 6 {
		t.Errorf("mappedStream[0] = %d; want 6", got[0])
	}
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
	if len(got) != 3 {
		t.Errorf("len(mappedStream) = %d; want 3", len(got))
	}
	if got[0] != 1 {
		t.Errorf("mappedStream[0] = %d; want 1", got[0])
	}
	if got[1] != 1 {
		t.Errorf("mappedStream[0] = %d; want 3", got[0])
	}
	if got[2] != 1 {
		t.Errorf("mappedStream[0] = %d; want 6", got[0])
	}
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
	if len(got) != 3 {
		t.Errorf("len(got) = %d; want 3", len(got))
	}
	if got[0] != 1 {
		t.Errorf("got[0] = %d; want 1", got[0])
	}
	if got[1] != 2 {
		t.Errorf("got[0] = %d; want 3", got[0])
	}
	if got[2] != 3 {
		t.Errorf("got[0] = %d; want 6", got[0])
	}
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
	if len(resCollector) != 3 {
		t.Errorf("len(got) = %d; want 3", len(got))
	}
	for _, elem := range resCollector {
		if elem != 2 {
			t.Errorf("%d; want 2", elem)
		}
	}
	for _, elem := range got {
		if elem != 1 {
			t.Errorf("%d; want 1", elem)
		}
	}
}

func TestApplyOnEach(t *testing.T) {
	// prepare
	initSlice := []int{1, 1, 1}

	// call
	got := From(initSlice).
		ApplyOnEach(func(it int) int { return it + 1 }).
		ToSlice()

	// assert
	if len(got) != 3 {
		t.Errorf("len(got) = %d; want 3", len(got))
	}
	for _, elem := range got {
		if elem != 2 {
			t.Errorf("%d; want 2", elem)
		}
	}
}

func TestReversed(t *testing.T) {
	// call
	slice1 := Of(1, 2, 3).Reversed().ToSlice()
	filteredSlice := Of(1, 2, 3).
		Filter(func(it int) bool { return it > 1 }).
		Reversed().
		ToSlice()

	// assert
	if len(slice1) != 3 {
		t.Errorf("len(slice1) = %d; want 3", len(slice1))
	}
	if len(filteredSlice) != 2 {
		t.Errorf("len(slice2) = %d; want 2", len(filteredSlice))
	}
	if (slice1[0] != 3) || (slice1[1] != 2) || (slice1[2] != 1) {
		t.Errorf("slice1 = %v; want [3,2,1]", slice1)
	}
	if (filteredSlice[0] != 3) || (filteredSlice[1] != 2) {
		t.Errorf("filteredSlice = %v; want [3,2]", filteredSlice)
	}
}

func TestDistinct(t *testing.T) {
	// call
	slice1 := Of(1, 2, 3, 3).Distinct().ToSlice()
	filteredSlice := Of(1, 2, 3, 3).
		Filter(func(it int) bool { return it > 1 }).
		Distinct().
		ToSlice()

	// assert
	if len(slice1) != 3 {
		t.Errorf("len(slice1) = %d; want 3", len(slice1))
	}
	if len(filteredSlice) != 2 {
		t.Errorf("len(slice2) = %d; want 2", len(filteredSlice))
	}
	if (slice1[2] != 3) || (slice1[1] != 2) || (slice1[0] != 1) {
		t.Errorf("slice1 = %v; want [1,2,3]", slice1)
	}
	if (filteredSlice[1] != 3) || (filteredSlice[0] != 2) {
		t.Errorf("filteredSlice = %v; want [2,3]", filteredSlice)
	}
}

func TestChunked(t *testing.T) {
	// call
	slice1 := Of(1, 2, 3, 4, 5, 6).Chunked(2)

	// assert
	if len(slice1) != 3 {
		t.Errorf("Chunked(2) = %d; want 3", len(slice1))
	}
	for _, slice := range slice1 {
		if len(slice) != 2 {
			t.Errorf("Chunked(2) = %d; want 2", len(slice))
		}
	}
}

func TestChunkedOdd(t *testing.T) {
	// call
	slice1 := Of(1, 2, 3, 4, 5, 6, 7).Chunked(2)

	// assert
	if len(slice1) != 4 {
		t.Errorf("Chunked(2) = %d; want 3", len(slice1))
	}
	if len(slice1[3]) != 1 {
		t.Errorf("Chunked(2) = %d; want 4", len(slice1[3]))
	}
	for i := 0; i < len(slice1)-1; i++ {
		if len(slice1[i]) != 2 {
			t.Errorf("Chunked(2) = %d; want 2", len(slice1[i]))
		}
	}
}

func TestWindowed(t *testing.T) {
	// call
	slice1 := Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15).Windowed(5, 3)

	// assert
	if len(slice1) != 4 {
		t.Errorf("Windowed(5, 3) = %d; want 3", len(slice1))
	}
	for i := 0; i < len(slice1); i++ {
		if len(slice1[i]) != 5 {
			t.Errorf("Windowed(5, 3) = %d; want 5", len(slice1[i]))
		}
	}
}

func TestWindowedPartial(t *testing.T) {
	// call
	slice1 := Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15).Windowed(5, 3, true)

	// assert
	if len(slice1) != 5 {
		t.Errorf("Windowed(5, 3) = %d; want 3", len(slice1))
	}
	if len(slice1[4]) != 3 {
		t.Errorf("Windowed(5, 3) = %d; want 3", len(slice1[4]))
	}
	for i := 0; i < len(slice1)-1; i++ {
		if len(slice1[i]) != 5 {
			t.Errorf("Windowed(5, 3) = %d; want 5", len(slice1[i]))
		}
	}
}

func TestWindowedBigWindow(t *testing.T) {
	// call
	slice1 := Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15).Windowed(16, 3)

	// assert
	if len(slice1) != 1 {
		t.Errorf("Windowed(16, 3) = %d; want 1", len(slice1))
	}
	if len(slice1[0]) != 15 {
		t.Errorf("Windowed(16, 3) = %d; want 15", len(slice1[0]))
	}
}

func TestTake(t *testing.T) {
	// call
	got := Of(0, 1, 2, 3).
		Take(2).
		ToSlice()

	// assert
	if len(got) != 2 {
		t.Errorf("len(take(2)) = %d; want 2", len(got))
	}
	if got[0] != 0 || got[1] != 1 {
		t.Errorf("take = %v", got)
	}
}

func TestTakeNothing(t *testing.T) {
	// call
	got := Of(0, 1, 2, 3).
		Take(0).
		ToSlice()

	// assert
	if len(got) != 0 {
		t.Errorf("len(take(0)) = %d; want 0", len(got))
	}
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
	if len(got) != 2 {
		t.Errorf("len(drop(2)) = %d; want 2", len(got))
	}
	if got[0] != 2 || got[1] != 3 {
		t.Errorf("drop = %v", got)
	}
}

func TestEmptyDrop(t *testing.T) {
	// call
	got := Of(0, 1, 2, 3).
		Drop(5).
		ToSlice()

	// assert
	if len(got) != 0 {
		t.Errorf("len(drop(2)) = %d; want 0", len(got))
	}
}

func TestDropNothing(t *testing.T) {
	// call
	got := Of(0, 1, 2, 3).
		Drop(0).
		ToSlice()

	// assert
	if len(got) != 4 {
		t.Errorf("len(drop(0)) = %d; want 4", len(got))
	}
}
