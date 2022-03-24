# go-strm

This is an experimental package which make use of the Generics feature introduced in Go 1.18 for 
providing a functional oriented Map/Reduce api in Go.

## Background

This api leverages Go Generics for providing a set of higher-order Map/Reduce functions.
These functions, when chained together, allow for functional programming techniques that ultimately reduce code
duplication and make it easier to transform and iterate over collections of elements. 
The current Generics' implementation in Go 1.18 don't allow methods themselves to have additional type parameters. 
This limitation forces mapping functions, whose input and return type differs, to be defined as top-level, 
limiting their chaining capacity and compromising readability. 
Despite this limitation, some of the most common functional programming techniques are still possible to be implemented 
in Go, as samples bellow illustrate.


### Generic Usage

#### Building from elements

```go
stringStrm := strm.Of("Hey!", "Hello!", "Hi!")
intStrm := strm.Of(1, 2, 3, 4)
sliceStrm := strm.Of([]int{1}, []int{1, 2}, []int{1, 2, 3})
```

#### Building from slices

```go
slice := []int {1, 2, 3}
intStrm := strm.From(slice)
```

#### Building from copy slices

```go
slice := []int {1, 2, 3}
intStrm := strm.CopyFrom(slice)
```

#### Converting back to slices

```go
intSlice := strm.Of(1, 2, 3, 4).ToSlice()
```

#### Filtering

```go
// Filtering a strm
isEven := func(n int) bool { return n%2 == 0 }

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

// iterating over all elements
// prints -> n: 2;   n: 4;   
strm.Of(1, 2, 3, 4, 5).
    Filter(isEven).
    ForEach(func(n int) { fmt.Printf("n: %v;\t", n) })
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

// maps a strm of (Person) to a slice of Person.name (string)  
// names -> [Peter John Sarah Kate]
names := strm.
    Map(strm.From(people), func(p Person) string { return p.name }).
    ToSlice()
```

#### De-duping and Reversing

```go
// deduped -> [2 3 4 5 6]
deduped := strm.Of(2, 2, 3, 4, 4, 5, 6, 6).
    Distinct().
    ToSlice()

// dedupedStruct -> [{Peter 18} {Bruce 48}]
dedupedStruct := strm.Of(Person{"Peter", 18}, Person{"Peter", 18}, Person{"Bruce", 48}).
    Distinct().
    ToSlice()

// reversed -> [6 5 4 3 2 1]
reversed := strm.Of(1, 2, 3, 4, 5, 6).
    Reversed().
    ToSlice()
```

#### Usual terminal operations

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


// names -> Peter,John,Sarah,Kate
people := strm.From([]Person{{"Peter", 30}, {"John", 18}, {"Sarah", 16}, {"Kate", 16}})
names := strm.Map(people, func(p Person) string { return p.name }).
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