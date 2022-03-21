package strm

import "testing"

func TestOf(t *testing.T) {
	// call
	stream := Of(1, 2, 3)
	// assert
	if got := stream.Slice; len(got) != 3 {
		t.Errorf("len(stream) = %d; want 3", len(got))
	}
}

func TestFrom(t *testing.T) {
	// prepare
	slice := [][]int{{1}, {1, 2}, {1, 2, 3}}
	// call
	stream := From(slice)
	// assert
	if got := stream.Slice; len(got) != 3 {
		t.Errorf("len(stream) = %d; want 3", len(got))
	}
}

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

func TestFilter(t *testing.T) {
	// prepare
	initSlice := [][]int{{1}, {1, 2}, {1, 2, 3}}
	// call
	got := From(initSlice).
		Filter(func(it []int) bool { return len(it) > 2 }).
		ToSlice()
	// assert
	if len(got) != 1 {
		t.Errorf("len(initSliceinitslice) = %d; want 1", len(got))
	}
	if len(got[0]) != 3 {
		t.Errorf("len(initSliceinitslice(0)) = %v; want 3", len(got[0]))
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

func TestChainedFilters(t *testing.T) {
	// prepare
	initSlice := [][]int{{1}, {1, 2}, {1, 2, 3}}
	// call
	got := From(initSlice).
		Filter(func(it []int) bool { return len(it) > 1 }).
		Filter(func(it []int) bool { return len(it) < 3 }).
		ToSlice()

	// assert
	if len(got) != 1 {
		t.Errorf("len(stream) = %d; want 1", len(got))
	}
	if len(got[0]) != 2 {
		t.Errorf("len(stream(0)) = %v; want 2", len(got[0]))
	}
}

func TestLazyFilters(t *testing.T) {
	// prepare
	initSlice := [][]int{{1}, {1, 2}, {1, 2, 3}}
	// call
	got := From(initSlice).
		Filter(func(it []int) bool { return len(it) > 1 }).
		Filter(func(it []int) bool { return len(it) < 3 })

	// assert
	if len(got.Slice) != 3 {
		t.Errorf("len(stream.Slice) = %d; want 3", len(got.Slice))
	}
}

func TestMapReduce(t *testing.T) {
	// prepare
	initSlice := [][]int{{1}, {1, 2}, {1, 2, 3}}

	// call
	got := Map(
		From(initSlice),
		func(it []int) int { return Reduce(From(it), 0, func(a int, b int) int { return a + b }) },
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
	initSlice := [][]int{{1}, {1, 2}, {1, 2, 3}}
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

func TestGroupBy(t *testing.T) {
	// prepare
	type Person struct {
		name string
		age  int
	}
	slice := []Person{Person{"Tim", 30}, Person{"Bil", 40}, Person{"John", 30}, Person{"Tim", 35}}

	// call
	byAge := GroupBy(From(slice), func(it Person) int { return it.age })
	byName := GroupBy(From(slice), func(it Person) string { return it.name })

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
