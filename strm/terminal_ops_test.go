package strm

import (
	"testing"
)

func TestToSlice(t *testing.T) {
	// prepare
	initSlice := []int{1, 2, 3}
	// call
	resSlice := From(initSlice).ToSlice()
	// assert
	if got := resSlice; len(got) != 3 {
		t.Errorf("len(stream) = %d; want 3", len(got))
	}
	for i := range resSlice {
		if got := resSlice[i]; initSlice[i] != resSlice[i] {
			t.Errorf("initSlice[i] != resSlice[i] = %d", got)
		}
	}
}

func TestForEach(t *testing.T) {
	// prepare
	initSlice := []int{1, 1, 1}
	var resCollector []int

	// call
	From(initSlice).ForEach(func(it int) { resCollector = append(resCollector, it+1) })

	// assert
	if len(resCollector) != 3 {
		t.Errorf("len(got) = %d; want 3", len(initSlice))
	}
	for _, elem := range resCollector {
		if elem != 2 {
			t.Errorf("%d; want 2", elem)
		}
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
	if !gotLen {
		t.Errorf("Any (len): %v want true", gotLen)
	}
	if !gotMatch {
		t.Errorf("Any (match): %v want true", gotLen)
	}
	if gotBye {
		t.Errorf("Any (match): %v want false", gotLen)
	}
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
	if !gotLen {
		t.Errorf("All (len): %v want true", gotLen)
	}
	if !gotLenFiltered {
		t.Errorf("All (len filtered): %v want true", gotLen)
	}
	if gotMatch {
		t.Errorf("All (match): %v want false", gotMatch)
	}
}

func TestNone(t *testing.T) {
	// prepare
	initSlice := []string{"Hi!", "Hello!", "Hey", ""}
	// call
	gotMatch := From(initSlice).None(func(it string) bool { return it == "Hey" })
	gotLen := From(initSlice).None(func(it string) bool { return len(it) < 3 })
	gotLenFiltered := Of("Hi!", "Hello!", "Hey", "").
		Filter(func(it string) bool { return len(it) > 0 }).
		None(func(it string) bool { return len(it) < 3 })

	// assert
	if !gotLenFiltered {
		t.Errorf("None (len Filtered): %v want false", gotLen)
	}
	if gotLen {
		t.Errorf("None (len): %v want false", gotLen)
	}
	if gotMatch {
		t.Errorf("None (match): %v want false", gotMatch)
	}
}

func TestCount(t *testing.T) {
	// prepare
	var slice2 []int
	// call
	count1 := Of(1, 2, 3).Filter(func(it int) bool { return it > 1 }).Count()
	count2 := From(slice2).Count()
	// assert
	if count1 != 2 {
		t.Errorf("count1 = %d; want 2", count1)
	}
	if count2 != 0 {
		t.Errorf("count2 = %d; want 0", count2)
	}
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
	if count1 != 1 {
		t.Errorf("count1 = %d; want 1", count1)
	}
	if count2 != 0 {
		t.Errorf("count2 = %d; want 0", count2)
	}
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
	if sumBy != 6 {
		t.Errorf("sumBy = %d; want 4", sumBy)
	}
	if sum1 != 0 {
		t.Errorf("sum1 = %d; want 0", sum1)
	}
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
	if first1 != 0 {
		t.Errorf("first1 = %d; want 0", first1)
	}
	if first2 != 2 {
		t.Errorf("first2 = %d; want 2", first2)
	}
}

func TestFirstBy(t *testing.T) {
	// prepare
	var slice2 []int
	// call
	first1 := From(slice2).FirstBy(func(it int) bool { return it > 1 })
	first2 := Of(1, 2, 3).
		Filter(func(it int) bool { return it > 1 }).
		FirstBy(func(it int) bool { return it > 2 })

	// assert
	if first1 != 0 {
		t.Errorf("first1 = %d; want 0", first1)
	}
	if first2 != 3 {
		t.Errorf("first2 = %d; want 3", first2)
	}
}

func TestLast(t *testing.T) {
	// prepare
	var slice2 []int
	// call
	last1 := From(slice2).Last()
	last2 := Of(1, 2, 3).
		Filter(func(it int) bool { return it > 1 }).
		Last()

	// assert
	if last1 != 0 {
		t.Errorf("last1 = %d; want 0", last1)
	}
	if last2 != 3 {
		t.Errorf("last2 = %d; want 2", last2)
	}
}

func TestJoinToString(t *testing.T) {
	// prepare
	var slice2 []int
	// call
	str1 := From(slice2).JoinToString(",")
	str2 := Of(1, 2, 3).
		Filter(func(it int) bool { return it > 1 }).
		JoinToString("-")

	// assert
	if str1 != "" {
		t.Errorf("str1 = %v; want empty", str1)
	}
	if str2 != "2-3" {
		t.Errorf("str2 = %v; want \"2-3\"", str2)
	}
}

func TestContains(t *testing.T) {
	// prepare
	type Person struct {
		name string
		age  int
	}
	// call
	containInt := Of(1, 2, 3).Contains(2)
	containsInt2 := Of(1, 2, 3).Contains(4)
	containsStruct := Of(Person{"Tim", 30}, Person{"Tom", 40}).Contains(Person{"Tom", 40})
	containsString := Of("Tim", "Tom").Contains("tom")
	containsSlice := From([][]int{{1}, {1, 2}, {1, 2, 3}}).Contains([]int{1})

	// assert
	if !containInt {
		t.Errorf("containInt(2) = %v; want true", containInt)
	}
	if containsInt2 {
		t.Errorf("containInt(4) = %v; want false", containsInt2)
	}
	if !containsStruct {
		t.Errorf("containInt(Person{\"Tom\", 40}) = %v; want true", containsStruct)
	}
	if containsString {
		t.Errorf("containInt(\"tom\") = %v; want false", containsString)
	}
	if containsSlice {
		t.Errorf("containInt(\"[]int{1}\") = %v; want false", containsSlice)
	}
}
