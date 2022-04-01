# go-strm

This is package which make use of the Generics feature introduced in Go 1.18 for 
providing a functional oriented reach Map/Reduce API in Go.
As this is still an experimental package, some of its apis may change slightly in future releases.

## Background

This api leverages Go Generics for providing a set of higher-order Map/Reduce functions.
These functions, when chained together, allow for functional programming techniques that ultimately reduce code
duplication and make it easier to transform and iterate over collections of elements. 
The current Generics' implementation in Go 1.18 doesn't allow methods themselves to have additional type parameters. 
This limitation forces mapping functions, whose input and return type differs, to be defined as top-level, 
limiting their chaining capacity and compromising readability. 
Despite this limitation, some of the most common functional programming techniques are still possible to be implemented 
in Go, as samples bellow illustrate.

All available Api operations are [enumerated here](#api-operations-listing).

### Generic Usage

This api is a wrapper of its inputs, either a slice or a set of elements of a generic type.
This wrapper is internally backed from a slice - either passed explicitly by a constructor or created on the fly if 
built by a set of elements - providing a set of operations applied over the contents of its backing slice. 
All operations have been implemented aiming minimal memory allocation, hence never allocating intermediate slices.

### Building from elements
##### Creates a strm backed by slice containing the given elements
```go
stringStrm := strm.Of("Hey!", "Hello!", "Hi!")
intStrm := strm.Of(1, 2, 3, 4)
sliceStrm := strm.Of([]int{1}, []int{1, 2}, []int{1, 2, 3})
```

### Building from a slice
##### The `backingSlice` state will be modified by the operation applied to the strm
```go
backingSlice := []int {1, 2, 3}
intStrm := strm.From(backingSlice)
```

### Building from a copy of a slice
##### The strm will create its own `backing slice`, which will be a copy of the `initSlice`. The `initSlice` state remain unmodified, independent of the operation applied to the strm.

```go
initSlice := []int {1, 2, 3}
intStrm := strm.CopyFrom(initSlice)
```

### Converting back to a slice
##### The `backingSlice` will be returned after all the operations have been applied to the strm

```go
slice := strm.Of(1, 2, 3, 4).ToSlice()
```

### Generic Operations
#### Filtering

```go
isEven := func(n int) bool { return n%2 == 0 }

// Filtering even elements
// evenSlice -> [2 4]
evenSlice := strm.Of(1, 2, 3, 4, 5).
    Filter(isEven).
    ToSlice()

// filter chaining
// evenSlice -> [4]
evenSlice := strm.Of(1, 2, 3, 4, 5).
    Filter(isEven).
    Filter(func(n int) bool { return n > 2 }).
    ToSlice()
```

#### Iterating

```go
// iterating over all elements
// prints -> n: 2;   n: 4;   
strm.Of(1, 2, 3, 4, 5).
    Filter(isEven).
    ForEach(func(n int) { fmt.Printf("n: %v;\t", n) })

// printing each element and filtering
// prints -> n: 1
//           n: 2
//           even: 2
// slice -> [2]
slice := strm.Of(1, 2).
    OnEach(func(n int) { fmt.Printf("n: %v\n", n) }).
    Filter(isEven).
    OnEach(func(n int) { fmt.Printf("even: %v\n", n) }).
    ToSlice()
```

#### Mapping

```go
// applies a transformation (n->n) to each element
// sqrSlice -> [1 4 9 16 25 36]
sqrSlice := strm.Of(1, 2, 3, 4, 5, 6).
    ApplyOnEach(func(n int) int { return n * n }).
    ToSlice()

type Person struct { name string; age  int }
people := []Person{{"Peter", 30}, {"John", 18}, {"Sarah", 16}, {"Kate", 16}}

// maps a go-strm of (Person) to a slice of Person.name (string)  
// names -> [Peter John Sarah Kate]
names := Map(From(people), func(p Person) string { return p.name }).
        ToSlice()
```

#### Mapping & Reducing

```go
slices := [][]int{{1}, {1, 2}, {1, 2, 3}}

// sums -> [1 3 6]
sums := Map(
    From(slices),
    func(it []int) int { return Reduce(From(it), func(a int, b int) int { return a + b }) },
).ToSlice()

// sums -> [1 3 6]
sums := Map(
    From(slices),
    func(it []int) int { return Sum(From(it)) },
).ToSlice()

// mins -> [1 1 1]
mins := Map(
    From(slices),
    func(it []int) int { return Min(From(it)) },
).ToSlice()

// maxs -> [1 2 3]
maxs := Map(
    From(slices),
    func(it []int) int { return Max(From(it)) },
).ToSlice()

// flatSlice -> [1 1 2 1 2 3]
flatSlice := FlatMap(
    From(slices),
    func(it []int) *Stream[int] { return From(it) },
).ToSlice()
```

#### Parallel Mapping
A `PMap` function is available for applying the given mapping function over all stream elements in parallel, leveraging
goroutines. The `PMap` usage is similar to `Map`. By default, `PMap` will launch a new goroutine per each element
present in the given Stream. If the `batching` flag is provided, the parallel work is batched by number of
available logical CPUs.

```go
people := []Person{{"Peter", 30}, {"John", 18}, {"Sarah", 16}, {"Kate", 16}}

// maps a go-strm of (Person) to a slice of Person.name (string) in parallel 
// names -> [Peter John Sarah Kate]
names := strm.
    PMap(strm.From(people), func(p Person) string { return p.name }).
    ToSlice()

// maps a go-strm of (Person) to a slice of Person.name (string) in parallel without batching
// names -> [Peter John Sarah Kate]
names := strm.
    PMap(strm.From(people), func(p Person) string { return p.name }, true).
    ToSlice()
```

#### Grouping 

````go
people := strm.Of(Person{"Tim", 30}, Person{"Bil", 40}, Person{"John", 30}, Person{"Tim", 35})

// byAge -> map[30:[{Tim 30} {John 30}] 35:[{Tim 35}] 40:[{Bil 40}]]
byAge := strm.GroupBy(people, func(it Person) int { return it.age })

// byName -> map[Bil:[{Bil 40}] John:[{John 30}] Tim:[{Tim 30} {Tim 35}]]
byName := strm.GroupBy(people, func(it Person) string { return it.name })
````

#### De-duping and Reversing
`Distinct` de-dupes strms of both `Comparable` and `Non-Comparable` types

```go
// deduped -> [2 3 4 5 6]
deduped := strm.Of(2, 2, 3, 4, 4, 5, 6, 6).
    Distinct().
    ToSlice()

// dedupedStruct -> [{Peter 18} {Bruce 48}]
dedupedStruct := strm.Of(Person{"Peter", 18}, Person{"Peter", 18}, Person{"Bruce", 48}).
    Distinct().
    ToSlice()

// dedupedSlices -> [[1 2] [3 4]]
dedupedSlices := strm.Of([]int{1, 2}, []int{1, 2}, []int{3, 4}).
    Distinct().
    ToSlice()

// reversed -> [6 5 4 3 2 1]
reversed := strm.Of(1, 2, 3, 4, 5, 6).
    Reversed().
    ToSlice()
```

#### The usual terminal operations

```go
// all -> false
all := strm.Of(1, 2, 3, 4, 5).
	All(func(n int) bool { return n < 5 })

// none ->  false
none := strm.Of(1, 2, 3, 4, 5).
	None(func(n int) bool { return n < 5 })

// any -> true
any := strm.Of("Hey!", "Hello!", "Hi!").
	Any(func(n string) bool { return n == "Hi!" })

// count ->  3
count := strm.Of(2, 2, 3, 4, 4, 5, 6, 6).
	Filter(isEven).
	Distinct().
	Count()

// Sum only even numbers
// sumBy ->  6
sumBy := strm.Of(1, 2, 3, 4, 5).
	SumBy(func(n int) int { if n%2 == 0 { return n } else { return 0 } })

// Count only even numbers
// countBy ->  2
countBy := strm.Of(1, 2, 3, 4, 5).
	CountBy(isEven)

// contains -> true 
contains := strm.Of(Person{"Peter", 18}, Person{"John", 30}, Person{"Bruce", 48}).
	Contains(Person{"Bruce", 48})

// contains -> true
contains := strm.Of([]int{1, 2}, []int{3, 4}).
	Contains([]int{1, 2})

// names -> Peter,John,Sarah,Kate
people := From([]Person{{"Peter", 30}, {"John", 18}, {"Sarah", 16}, {"Kate", 16}})
names := Map(people, func(p Person) string { return p.name }).
	JoinToString(",")
```

#### Chunked and Windowed

````go
// converting to batches of 2 elements each
// batches -> [[1 2] [4 6] [14 1] [2]]
batches := strm.Of(1, 2, 4, 6, 14, 1, 2).
	Chunked(2)

// converting to windows of 5 elements with a step of 3, without partial windows at the end
// windows -> [[1 2 3 4 5] [4 5 6 7 8] [7 8 9 10 11] [10 11 12 13 14]]
windows := strm.Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15).
	Windowed(5, 3)

// converting to windows of 5 elements with a step of 3, preserving all partial windows at the end
// partWindows -> [[1 2 3 4 5] [4 5 6 7 8] [7 8 9 10 11] [10 11 12 13 14] [13 14 15]]
partWindows := strm.Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15).
	Windowed(5, 3, true)
````

#### Picking elements

````go
// first -> {Peter 18}
first := strm.Of(Person{"Peter", 18}, Person{"John", 30}, Person{"Bruce", 18}).
	First()

// firstBy -> {John 30}
firstBy := strm.Of(Person{"Peter", 18}, Person{"John", 30}, Person{"Bruce", 18}).
	FirstBy(func(p Person) bool { return p.age > 18 })

// last -> {"Bruce", 18}
last := strm.Of(Person{"Peter", 18}, Person{"John", 30}, Person{"Bruce", 18}).
	Last()

// take -> [0 1]
take := strm.Of(0, 1, 2, 3).
	Take(2).
	ToSlice()

// drop -> [1 2 3]
drop := strm.Of(0, 1, 2, 3).
	Drop(1).
	ToSlice()
````

#### Int Ranges 

An int range, represented by the `IntStream` type, is also available for convenience of use `Sum`, `Min`, `Max`, `Avg` and `Sorted` ops.
This type encloses a `*Stream[int]` exclusively, leveraging all methods already available in the Stream type. 

```go
// sum -> 15
sum := strm.Range(1,5).Sum()

// min -> 1
min := strm.RangeOf(1, 2, 3, 4, 5).Min()

// max -> 5
max := strm.RangeFrom([]int{1, 2, 3, 4, 5}).Max()

// avg -> 2
avg := strm.Range(1,4).Avg()

// sorted -> [1 2 3 4 5]
sorted := strm.RangeOf(5, 4, 2, 1, 3).
	Sorted().
	ToSlice()

// mappedSlice -> [2 3 4 5 6]
mappedSlice := Map(
	RangeOf(5, 4, 2, 1, 3).Sorted().ToStrm(),
    func(it int) int { return it+1 },
).ToSlice()

// batches -> [[2 3] [4 5] [6]]
batches := strm.RangeOf(5, 4, 2, 1, 3, 6).
	Sorted().
	Filter(func(it int) bool { return it > 1 }).
	Chunked(2)
```

### API Benchmarking 

Performance-wise, single mapping and filtering ops perform very well. Chained operations like applying several mappings 
and filters over the strm, can be slower than just performing native for loops.
The following benchmarks can be found at [strm_test.go](strm/strm_test.go). 
```
goos: darwin
goarch: amd64
pkg: strm/strm
cpu: Intel(R) Core(TM) i9-8950HK CPU @ 2.90GHz
name               old time/op   new time/op   delta
Filter-12           114ns ± 2%     98ns ± 5%   -13.81%  (p=0.008 n=5+5)
Distinct-12         526ns ± 3%    461ns ± 5%   -12.48%  (p=0.008 n=5+5)
Map-12             1.36µs ± 3%   1.55µs ± 3%   +13.35%  (p=0.008 n=5+5)
ChainedFilters-12   117ns ± 4%    138ns ± 5%   +17.97%  (p=0.008 n=5+5)
MapFilter-12       1.29µs ± 5%   1.65µs ± 3%   +27.77%  (p=0.008 n=5+5)
```

### API Operations Listing 

```go
// Constructors	
func Of[T any](elems ...T) *Stream[T]
func From[T any](backingSlice []T) *Stream[T]
func CopyFrom[T any](slice []T) *Stream[T]

// Top-Level functions
func Map[IN any, OUT any](s *Stream[IN], f func(IN) OUT) *Stream[OUT]
func PMap[IN any, OUT any](s *Stream[IN], f func(IN) OUT) *Stream[OUT]
func FlatMap[IN any, OUT any](s *Stream[IN], f func(v IN) *Stream[OUT]) *Stream[OUT]
func Reduce[IN any, OUT any](s *Stream[IN], f reducer[OUT, IN], start ...OUT) OUT
func GroupBy[K comparable, V any](s *Stream[V], keySelector func(V) K) map[K][]V
func Max[O Ordered](s *Stream[O]) O
func Min[O Ordered](s *Stream[O]) O
func Sum[O Ordered](s *Stream[O]) O
func Merge[T any](streams ...*Stream[T]) *Stream[T]

// go-strm operations
func Filter(predicate func(T) bool) *Stream[T]
func ApplyOnEach(action func(T) T) *Stream[T]
func OnEach(f func(T)) *Stream[T]
func Plus(other *Stream[T]) *Stream[T]
func Append(elems []T) *Stream[T]
func Take(n int) *Stream[T]
func Drop(n int) *Stream[T]
func Reversed() *Stream[T]
func Distinct() *Stream[T]

// Terminal go-strm operations
func ToSlice() []T
func ForEach(action func(T))
func Any(predicate func(T) bool) bool
func All(predicate func(T) bool) bool
func None(predicate func(T) bool) bool
func Count() int
func CountBy(predicate func(T) bool) int
func SumBy(selector func(T) int) int
func FirstBy(predicate func(T) bool) T
func First() T
func Last() T
func Contains(element T) bool
func JoinToString(delimiter string) string
func Chunked(batchSize int) [][]T
func Windowed(size int, step int, partialWindows ...bool) [][]T

// Int Ranges operations
func Range(from int, to int) *IntStream
func RangeOf(elems ...int) *IntStream
func RangeFrom(backingIntSlice []int) *IntStream
func RangeCopyFrom(backingIntSlice []int) *IntStream
func Sorted() *IntStream
func Sum() int
func Min() int
func Max() int
func Avg() int
func ToStrm() *Stream[int]
```
