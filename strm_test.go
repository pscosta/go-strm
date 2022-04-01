package strm

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOf(t *testing.T) {
	// call
	slice := Of(1, 2, 3).ToSlice()
	// assert
	assert.Equal(t, 3, len(slice), "wrong length")
}

func TestFrom(t *testing.T) {
	// call
	slice := From([][]int{{1}, {1, 2}, {1, 2, 3}}).ToSlice()
	// assert
	assert.Equal(t, 3, len(slice), "wrong length")
}

func TestFilterCopyFrom(t *testing.T) {
	// prepare
	initSlice := [][]int{{1}, {1, 2}, {1, 2, 3}}
	// call
	got := CopyFrom(initSlice).
		Filter(func(it []int) bool { return len(it) > 2 }).
		ToSlice()
	initSlice[0] = []int{1, 2}

	// assert
	assert.Equal(t, 1, len(got), "wrong length")
	assert.Equal(t, 3, len(initSlice), "wrong length")
	assert.Equal(t, 3, len(got[0]), "wrong length")
}

func TestMapReduce(t *testing.T) {
	// prepare
	initSlice := [][]int{{1}, {1, 2}, {1, 2, 3}}

	// call
	got := Map(
		From(initSlice),
		func(it []int) int { return Reduce(From(it), func(a int, b int) int { return a + b }) },
	).ToSlice()

	// assert
	assert.Equal(t, 3, len(got), "wrong length")
	assert.Equal(t, 1, got[0], "wrong value")
	assert.Equal(t, 3, got[1], "wrong value")
	assert.Equal(t, 6, got[2], "wrong value")
}

func TestParallelLinearMapReduce(t *testing.T) {
	// prepare
	initSlice := [][]int{{1}, {1, 2}, {1, 2, 3}}

	// call
	got := PMap(
		From(initSlice),
		func(it []int) int { return Reduce(From(it), func(a int, b int) int { return a + b }) },
	).ToSlice()

	// assert
	assert.Equal(t, 3, len(got), "wrong length")
	assert.Equal(t, 1, got[0], "wrong value")
	assert.Equal(t, 3, got[1], "wrong value")
	assert.Equal(t, 6, got[2], "wrong value")
}

func TestParallelBatchedMapReduce(t *testing.T) {
	// prepare
	initSlice := [][]int{{1}, {1, 2}, {1, 2, 3}}

	// call
	got := PMap(
		From(initSlice),
		func(it []int) int { return Reduce(From(it), func(a int, b int) int { return a + b }) },
		true,
	).ToSlice()

	// assert
	assert.Equal(t, 3, len(got), "wrong length")
	assert.Equal(t, 1, got[0], "wrong value")
	assert.Equal(t, 3, got[1], "wrong value")
	assert.Equal(t, 6, got[2], "wrong value")
}

func TestMapReduceWithStart(t *testing.T) {
	// prepare
	initSlice := [][]int{{1}, {1, 2}, {1, 2, 3}}

	// call
	got := Map(
		From(initSlice),
		func(it []int) int { return Reduce(From(it), func(a int, b int) int { return a + b }, 1) },
	).ToSlice()

	// assert
	assert.Equal(t, 3, len(got), "wrong length")
	assert.Equal(t, 2, got[0], "wrong value")
	assert.Equal(t, 4, got[1], "wrong value")
	assert.Equal(t, 7, got[2], "wrong value")
}

func TestFlatMap(t *testing.T) {
	// call
	got := FlatMap(
		Of(1, 2, 4),
		func(e int) *Stream[string] { return Of(fmt.Sprint(e)) },
	).ToSlice()

	// assert
	assert.Equal(t, 3, len(got), "wrong length")
	assert.Equal(t, "1", got[0], "wrong value")
	assert.Equal(t, "2", got[1], "wrong value")
	assert.Equal(t, "4", got[2], "wrong value")
}

func TestGroupBy(t *testing.T) {
	// prepare
	type Person struct {
		name string
		age  int
	}
	stream := Of(Person{"Tim", 30}, Person{"Bil", 40}, Person{"John", 30}, Person{"Tim", 35})

	// call
	byAge := GroupBy(stream, func(it Person) int { return it.age })
	byName := GroupBy(stream, func(it Person) string { return it.name })

	// assert
	assert.Equal(t, 3, len(byAge), "wrong length")
	assert.Equal(t, 3, len(byName), "wrong length")
	assert.Equal(t, 30, byAge[30][0].age, "wrong value")
	assert.Equal(t, "Tim", byAge[30][0].name, "wrong value")
	assert.Equal(t, 30, byAge[30][1].age, "wrong value")
	assert.Equal(t, "John", byAge[30][1].name, "wrong value")
	assert.Equal(t, 40, byName["Bil"][0].age, "wrong value")
	assert.Equal(t, "Bil", byName["Bil"][0].name, "wrong value")
}

func BenchmarkApiFilter(b *testing.B) {
	var result []int

	for i := 0; i < b.N; i++ {
		result = From([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).
			Filter(func(n int) bool { return n%2 == 0 }).
			ToSlice()
	}

	require.Equal(b, []int{2, 4, 6, 8, 10}, result)
}

func BenchmarkStandardFilter(b *testing.B) {
	var res []int

	for i := 0; i < b.N; i++ {
		res = nil
		for _, n := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
			if n%2 == 0 {
				res = append(res, n)
			}
		}
	}

	require.Equal(b, []int{2, 4, 6, 8, 10}, res)
}

func BenchmarkApiChainedFilters(b *testing.B) {
	var result []int

	for i := 0; i < b.N; i++ {
		result = From([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).
			Filter(func(n int) bool { return n%2 == 0 }).
			Filter(func(n int) bool { return n <= 10 }).
			ToSlice()
	}

	require.Equal(b, []int{2, 4, 6, 8, 10}, result)
}

func BenchmarkStandardChainedFilters(b *testing.B) {
	var res []int

	for i := 0; i < b.N; i++ {
		res = nil
		for _, n := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
			if n%2 == 0 && n <= 10 {
				res = append(res, n)
			}
		}
	}

	require.Equal(b, []int{2, 4, 6, 8, 10}, res)
}

type Person struct {
	name string
	age  int
}

func BenchmarkApiDistinct(b *testing.B) {
	var result []Person
	people := []Person{{"Peter", 30}, {"Peter", 30}, {"John", 30}, {"Tim", 30}, {"Tom", 30}}

	for i := 0; i < b.N; i++ {
		result = From(people).
			Distinct().
			ToSlice()
	}

	require.Equal(b, []Person{{"Peter", 30}, {"John", 30}, {"Tim", 30}, {"Tom", 30}}, result)
}

func BenchmarkStandardDistinct(b *testing.B) {
	var people []Person

	for idx := 0; idx < b.N; idx++ {
		people = []Person{{"Peter", 30}, {"Peter", 30}, {"John", 30}, {"Tim", 30}, {"Tom", 30}}
		keys := make(map[any]struct{}, len(people))
		j := 0
		for i := 0; i < len(people); i++ {
			if _, ok := keys[people[i]]; ok {
				continue
			}
			keys[people[i]] = struct{}{}
			people[j], j = people[i], j+1
		}
		people = people[:j]
	}

	require.Equal(b, []Person{{"Peter", 30}, {"John", 30}, {"Tim", 30}, {"Tom", 30}}, people)
}

func BenchmarkApiMap(b *testing.B) {
	var result []string
	var people []Person
	for i := 0; i <= 100; i++ {
		people = append(people, Person{"Person", i})
	}

	for i := 0; i < b.N; i++ {
		result = Map(From(people),
			func(p Person) string { return p.name },
		).ToSlice()
	}

	require.Equal(b, len(people), len(result))
}

func BenchmarkApiPMap(b *testing.B) {
	var result []Person
	var people []string
	for i := 0; i <= 10000000; i++ {
		people = append(people, `{"name": "peter", "age": 39}`)
	}

	for i := 0; i < b.N; i++ {
		result = PMap(From(people),
			func(s string) Person {
				var p Person
				_ = json.Unmarshal([]byte(s), &p)
				_, _ = json.Marshal(`{"name": "peter", "age": 39}`) // just another heavy op
				return p
			},
		).ToSlice()
	}

	require.Equal(b, len(people), len(result))
}

func BenchmarkStandardMap(b *testing.B) {
	var result []string
	var people []Person
	for i := 0; i <= 100; i++ {
		people = append(people, Person{"Person", i})
	}

	for i := 0; i < b.N; i++ {
		result = nil
		for _, p := range people {
			result = append(result, p.name)
		}
	}

	require.Equal(b, len(people), len(result))
}

func BenchmarkApiMapFilter(b *testing.B) {
	var result []string
	var people []Person
	for i := 0; i <= 100; i++ {
		people = append(people, Person{"Person", i})
	}

	for i := 0; i < b.N; i++ {
		result = Map(CopyFrom(people).Filter(func(p Person) bool { return p.age < 30 }),
			func(p Person) string { return p.name },
		).ToSlice()
	}

	require.Equal(b, 30, len(result))
}

func BenchmarkStandardMapFilter(b *testing.B) {
	var result []string
	var people []Person
	for i := 0; i <= 100; i++ {
		people = append(people, Person{"Person", i})
	}

	for i := 0; i < b.N; i++ {
		result = nil
		var ppl []Person
		// performs 2 loops
		for _, p := range people {
			if p.age < 30 {
				ppl = append(ppl, p)
			}
		}
		for _, p := range ppl {
			result = append(result, p.name)
		}
	}

	require.Equal(b, 30, len(result))
}

func BenchmarkApiForEach(b *testing.B) {
	var people []Person
	for i := 0; i <= 100; i++ {
		people = append(people, Person{"Person", i})
	}

	for i := 0; i < b.N; i++ {
		From(people).ForEach(func(p Person) { p.age++ })
	}
}

func BenchmarkStandardForEach(b *testing.B) {
	var people []Person
	for i := 0; i <= 100; i++ {
		people = append(people, Person{"Person", i})
	}

	for i := 0; i < b.N; i++ {
		for _, p := range people {
			p.age++
		}
	}
}

func BenchmarkApiAny(b *testing.B) {
	var result bool
	var people []Person
	for i := 0; i <= 100; i++ {
		people = append(people, Person{"Person", i})
	}

	for i := 0; i < b.N; i++ {
		result = From(people).Any(func(p Person) bool { return p.age > 100 })
	}

	require.Equal(b, false, result)
}

func BenchmarkStandardAny(b *testing.B) {
	var result bool
	var people []Person
	for i := 0; i <= 100; i++ {
		people = append(people, Person{"Person", i})
	}

	for i := 0; i < b.N; i++ {
		for _, p := range people {
			if p.age > 100 {
				result = true
			}
		}
	}

	require.Equal(b, false, result)
}

func BenchmarkApiSumBy(b *testing.B) {
	var result int
	var people []Person
	for i := 0; i <= 100; i++ {
		people = append(people, Person{"Person", i})
	}

	for i := 0; i < b.N; i++ {
		result = From(people).SumBy(func(p Person) int { return p.age })
	}

	require.Equal(b, 5050, result)
}

func BenchmarkStandardSumBy(b *testing.B) {
	var people []Person
	for i := 0; i <= 100; i++ {
		people = append(people, Person{"Person", i})
	}
	sum := 0
	for i := 0; i < b.N; i++ {
		sum = 0
		for _, p := range people {
			sum += p.age
		}
	}

	require.Equal(b, 5050, sum)
}
