package main

import (
	"fmt"
)

type Person struct {
	Name string
	age  int
}

func main() {
	// from elements
	//stringStrm := strm.Of("Hey!", "Hello!", "Hi!")
	//intStrm := strm.Of(1, 2, 3, 4)
	//sliceStrm := strm.Of([]int{1}, []int{1, 2}, []int{1, 2, 3})
	//
	//// from slices
	//slice := []int{1, 2, 3}
	//intSliceStrm := strm.From(slice)
	//
	//// copy slice
	//slice := []int{1, 2, 3}
	//intSliceStrm := strm.CopyFrom(slice)
	//
	//// filtering
	//
	//isEven := func(n int) bool { return n%2 == 0 }
	//
	//evenSlice := strm.Of(1, 2, 3, 4, 5).
	//	Filter(isEven).
	//	ToSlice()
	//
	//// filter chaining
	//evenSlice := strm.Of(1, 2, 3, 4, 5).
	//	Filter(isEven).
	//	Filter(func(n int) bool { return n > 2 }).
	//	ToSlice()
	//
	//slice := strm.Of(1, 2, 4, 6, 14).
	//	OnEach(func(c int) { fmt.Println(c) }).
	//	Filter(func(v int) bool { return v%2 == 0 }).
	//	Filter(func(v int) bool { return v > 1 }).
	//	ToSlice()
	//
	//// iterating
	//
	//strm.Of(1, 2, 3, 4, 5).
	//	Filter(isEven).
	//	ForEach(func(n int) { fmt.Printf("n: %v\n", n) })
	//
	////dedup
	//
	//dedupp := strm.Of(2, 2, 3, 4, 4, 5, 6, 6).
	//	Filter(isEven).
	//	Distinct().
	//	ToSlice()

	//count := strm.Of(2, 2, 3, 4, 4, 5, 6, 6).
	//	Filter(isEven).
	//	Distinct().
	//	Count()
	//
	//reversed := strm.Of(2, 2, 3, 4, 4, 5, 6, 6).
	//	Reversed().
	//	ToSlice()
	//
	//applyO := strm.Of(1, 2, 3, 4, 5, 6).
	//	ApplyOnEach(func(n int) int { return n * n }).
	//	ToSlice()
	//
	//applyO := strm.Of(1, 2, 3, 4, 5, 6).
	//	ApplyOnEach(func(n int) int { return n + 1 }).
	//	Filter(isEven).
	//	ToSlice()
	//
	//sum := strm.Of(1, 2, 3, 4, 5).
	//	SumBy(func(n int) int { return n%2 == 0 })
	////
	//all := strm.Of(1, 2, 3, 4, 5).
	//	All(func(n int) bool { return n < 5 })
	////
	//none := strm.Of(1, 2, 3, 4, 5).
	//	None(func(n int) bool { return n < 5 })
	////
	//any := strm.Of(1, 2, 3, 4, 5).
	//	Filter(isEven).
	//	Any(func(n int) bool { return n < 5 })
	//
	//type Person struct {
	//	name string
	//	age  int
	//}

	//people := []Person{{"Peter", 30}, {"John", 18}, {"Sarah", 16}, {"Cate", 16}}

	//grouping := strm.GroupBy(
	//	strm.CopyFrom(people).Filter(func(p Person) bool { return p.age > 18 }),
	//	func(p Person) int { return p.age },
	//)
	//
	//ageSum := strm.From(people).SumBy(func(p Person) int { return p.age })
	//
	//strm.Map(
	//	strm.CopyFrom(people),
	//	func(p Person) string { return fmt.Sprintf("Name:%v, Age:%v", p.name, p.age) },
	//).
	//ApplyOnEach(func(s string) string { return strings.ToUpper(s) }).
	//ForEach(func(s string) { fmt.Println(s) })
	//
	//names := strm.Map(strm.From(people), func(p Person) string { return p.name })
	//namesLen := strm.Reduce(names, func(acc int, age string) int { return acc + len(age) },0)
	//
	//
	//
	//// call
	//sumAllSlices := strm.Map(
	//	strm.From(initSlice),
	//	func(it []int) int { return strm.Reduce(strm.From(it), func(a int, b int) int { return a + b }) },
	//).ToSlice()
	//

	//initSlice := [][]int{{1}, {1, 2}, {1, 2, 3}}
	//flatSlice := strm.FlatMap(
	//	strm.From(initSlice),
	//	func(it []int) *strm.Stream[int] { return strm.From(it) },
	//).ToSlice()
	//fmt.Println(flatSlice)

	//sumAllSlices2 := strm.Sum(numsFlatt)
	//
	//fmt.Println(sumAllSlices2)

	//fmt.Println(strm.Reduce(strm.Of(1, 2, 4, 6, 14, 1, 2), func(i string, v int) string { return i + strconv.Itoa(v) }, ""))
	//fmt.Println(strm.Reduce(strm.Of(1, 2, 4, 6, 14, 1, 2), func(i int, v int) int { return i + v }, 0))

	//strings := strm.Of("ola", "123", "123", "hi", "!", "")
	//grouping := strm.GroupBy(strings, func(s string) int { return len(s) })
	//
	//fmt.Println(grouping)
	//
	//fmt.Println(strm.Max(strm.Of(1, 2, 4, 6, 14)))
	//fmt.Println(strm.Min(strm.Of(1, 2, 4, 6, 14)))
	//fmt.Println(strm.Sum(strm.Of(1, 2, 4, 6, 14, 1, 2)))
	//fmt.Println(strm.Of(1, 2, 4, 6, 14, 1, 2).Distinct().Filter(func(i int) bool { return i > 1 }).ToSlice())
	//
	//fmt.Println(strm.Of(1).Filter(func(i int) bool { return i > 1 }).Distinct().ToSlice())
	//
	fmt.Println(Of(1, 2, 4, 6, 14, 1, 2).Distinct().ToSlice())
	//strm.Of([]int{1, 2}, []int{3, 4}, []int{5, 6}).ForEach(func(v []int) { fmt.Printf("%p\n", v) })
	//fmt.Println(strm.Of(1, 2, 4, 6, 14, 1, 2).Distinct().ToSlice())
	//fmt.Println(strm.Of(Person{"Peter", 18}, Person{"John", 30}, Person{"Bruce", 18}).Distinct().ToSlice())
	//fmt.Println(strm.Of([]int{1, 2}, []int{3, 4}, []int{5, 6}, []int{1, 2}).Distinct().ToSlice())
	//
	//fmt.Println(strm.Of(Person{"Peter", 18}, Person{"John", 30}, Person{"Bruce", 18}).Contains(Person{"Bruce", 18}))
	//fmt.Println(strm.Of(1, 2, 4, 6, 14, 1, 2).Distinct().Contains(2))
	//fmt.Println(strm.Of([]int{1, 2}, []int{3, 4}, []int{5, 6}, []int{1, 2}).Contains([]int{1, 2}))
	//fmt.Println(strm.Of(1).Filter(func(i int) bool { return i > 1 }).Reversed().Distinct())
	//fmt.Println(strm.Of(1).Filter(func(i int) bool { return i > 1 }).Reversed().Distinct())
	//fmt.Println(strm.Sum(strm.Of("ola", "123", "123", "hi", "!", "")))
	//fmt.Println(strm.Of(1, 2, 4, 6, 14, 1, 2).JoinToString(""))
	//fmt.Println(strm.GroupBy(strm.Of(Person{"Peter", 18}, Person{"John", 30}, Person{"Bruce", 18}), func(p Person) int { return p.age }))
	//fmt.Println(strm.GroupBy(strm.Of(Person{"Peter", 18}, Person{"John", 30}, Person{"Bruce", 18}), func(p Person) string { return p.name }))
	//
	//fmt.Println(
	//	strm.PMap(
	//		strm.Of(1, 2, 4, 6, 14, 1, 2),
	//		func(e int) *strm.Stream[string] { return strm.Of(fmt.Sprint(e)) },
	//	).ToSlice(),
	//)
	//
	//fmt.Println(strm.Of(1, 2, 3, 4, 5).
	//	CountBy(isEven))

	//slice3 := []int{3}
	//fmt.Println(strm.From([][]int{{1}, {2}, slice3, slice3}).Distinct().ToSlice())
	//fmt.Println(strm.Of(1, 2).Plus(strm.Of(3, 4, 5)).ToSlice())
	//fmt.Println(strm.Merge(strm.Of(1, 2), strm.Of(3), strm.Of(4, 5)).ToSlice())
	//filters := []int{0}
	//filters = nil
	//fmt.Println(len(filters))
	//people := strm.Of(Person{"Tim", 30}, Person{"Bil", 40}, Person{"John", 30}, Person{"Tim", 35})
	//byAge := strm.GroupBy(people, func(it Person) int { return it.age })
	//byName := strm.GroupBy(people, func(it Person) string { return it.name })

	//strm1 := strm.Of(1)
	//sl1 := func() {}
	//fmt.Println(byAge)
	//fmt.Println(byName)
	//fmt.Println(strm.Of(&sl1, &sl1).Distinct().ToSlice())
	//fmt.Println(strm.Of(Person{"Peter", 18}, Person{"Peter", 18}, Person{"Bruce", 18}).Distinct().ToSlice())
	//fmt.Println(strm.Of(Person{"Peter", 18}, Person{"Peter", 18}, Person{"Bruce", 18}).
	//	SumBy(func(p Person) int { return p.age }))
	//fmt.Println(strm.Sum(strm.Of(1).Filter(func(i int) bool { return false })))
	//fmt.Println(strm.Of(map[string]int{"a": 1, "b": 2}, map[string]int{"hello": 1, "bye": 2}).JoinToString(","))

	//initSlice := []string{"ola", "adeus", "goodbye", "byebye", "salut", "adieu", "oh"}
	//window := make([]string, 2)
	//window = append(initSlice[1:3])
	////copy(window, initSlice[1:3])
	//fmt.Println(window)
	////fmt.Println(initSlice[1:3])
	//initSlice[1] = "puta"
	//fmt.Println(window)
	//batches := strm.CopyFrom(initSlice).Chunked(2)
	//fmt.Println(batches)
	////batches[0][1] = "puta"
	//batches[0] = append(batches[0], "puta")
	//fmt.Println(batches)
	//
	//fmt.Println(strm.Of([]int{1, 2}, []int{3, 4}, []int{5, 6}, []int{1, 2}).Contains([]int{1, 2}))
	//
	//type Cenas struct {
	//	s []int
	//	b int
	//}
	//fmt.Println(strm.Of(Cenas{s: []int{1, 2}, b: 2}).Contains(Cenas{s: []int{1, 2}, b: 2}))
	//
	//fmt.Println(strm.Of(Cenas{s: []int{1, 2}, b: 2}, Cenas{s: []int{1, 2}, b: 2}).Distinct().ToSlice())
	//
	//type Person struct {
	//	name string
	//	age  int
	//}
	//people := []Person{{"Peter", 30}, {"Peter", 30}, {"Cate", 16}}
	//fmt.Println(strm.From(people).Distinct().ToSlice())
	//fmt.Println(strm.Of("a", "b", "c").Sum())
	//fmt.Println(strm.Of(1, 2, 3).Distinct().ToSlice())
	//fmt.Println(strm.Of([]int{1}, []int{1}, []int{1}).Distinct().ToSlice())
	//fmt.Println(strm.Of(Person{"Peter", 30}, Person{"Peter", 30}).Distinct().ToSlice())
	//fmt.Println(strm.Of([]int{1, 2}, []int{1, 2}, []int{3, 4}).
	//	Distinct().
	//	ToSlice())
	//	Contains([]int{1, 2}))
	//
	s := Of([]int{1, 2}, []int{3, 4}).
		Contains([]int{1, 2})
	fmt.Println(s)
}

//type Storage[T any] interface {
//	store(thing T)
//}
//type Box[T any] struct {
//	things []T
//}
//
//func (b *Box[T]) store(thing T) {
//	b.things = append(b.things, thing)
//}
//func (b *Box[int]) store(thing int) {
//	b.things = append(b.things, thing)
//}
//
//func main() {
//	var box Box[string]
//	box.store("lightSaber")
//}

//
//func (m Box[T]) Add(bags ...[]T) Box[[]T] {
//	sliceOfBags := make([][]T, 0, len(bags))
//	for _, bag := range bags {
//		sliceOfBags = append(sliceOfBags, bag)
//	}
//	return Box[[]T]{sliceOfBags}
//}
//
//func main() {
//	applesBag := []string{"green apple", "ref apple"}
//	orangesBag := []string{"juicy orange", "dry orange"}
//
//	Box[string]{}.Add(applesBag, orangesBag)
//}
//
//type Stream[T any] struct {
//	slice []T
//}

//func (s *Stream[T]) Chunked(batchSize int) *Stream[[]T] {
//	batches := make([][]T, 0, (len(s.slice)+batchSize-1)/batchSize)
//
//	for batchSize < len(s.slice) {
//		s.slice, batches = s.slice[batchSize:], append(batches, s.slice[0:batchSize:batchSize])
//	}
//	batches = append(batches, s.slice)
//	return &Stream[[]T]{batches}
//}
