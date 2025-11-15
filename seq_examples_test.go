package seq_test

import (
	"fmt"

	"github.com/banky/seq"
)

func ExampleFilter() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	evens := seq.Filter(numbers, func(n int) bool { return n%2 == 0 })
	fmt.Println(evens)
	// Output: [2 4 6]
}

func ExampleMap() {
	numbers := []int{1, 2, 3}
	strs := seq.Map(numbers, func(n int) string { return fmt.Sprintf("n=%d", n) })
	fmt.Println(strs)
	// Output: [n=1 n=2 n=3]
}

func ExampleReduce() {
	numbers := []int{1, 2, 3, 4}
	sum := seq.Reduce(numbers, 0, func(acc, n int) int { return acc + n })
	fmt.Println(sum)
	// Output: 10
}

func ExampleLastIndex() {
	values := []string{"a", "b", "c", "b"}
	idx, ok := seq.LastIndex(values, "b")
	fmt.Println(idx, ok)

	_, ok = seq.LastIndex(values, "z")
	fmt.Println(ok)
	// Output:
	// 3 true
	// false
}

func ExamplePartition() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	evens, odds := seq.Partition(numbers, func(n int) bool { return n%2 == 0 })
	fmt.Println(evens, odds)
	// Output: [2 4 6] [1 3 5]
}

func ExampleGroupBy() {
	words := []string{"a", "bb", "ccc", "dd", "e"}
	byLen := seq.GroupBy(words, func(s string) int { return len(s) })

	// Print groups in length order for stable output.
	fmt.Println(byLen[1])
	fmt.Println(byLen[2])
	fmt.Println(byLen[3])
	// Output:
	// [a e]
	// [bb dd]
	// [ccc]
}

func ExampleFlatten() {
	nested := [][]int{{1, 2}, {}, {3}, {4, 5}}
	flat := seq.Flatten(nested)
	fmt.Println(flat)
	// Output: [1 2 3 4 5]
}

func ExampleUnique() {
	values := []int{1, 2, 1, 3, 2, 4, 4}
	uniq := seq.Unique(values)
	fmt.Println(uniq)
	// Output: [1 2 3 4]
}

func ExampleUniqueBy() {
	type User struct {
		Name  string
		Email string
	}

	users := []User{
		{Name: "Alice", Email: "a@example.com"},
		{Name: "Bob", Email: "b@example.com"},
		{Name: "Alice Clone", Email: "a@example.com"},
	}

	uniq := seq.UniqueBy(users, func(u User) string { return u.Email })
	fmt.Println(uniq)
	// Output: [{Alice a@example.com} {Bob b@example.com}]
}

func ExampleChunk() {
	values := []int{1, 2, 3, 4, 5}
	chunks := seq.Chunk(values, 2)
	fmt.Println(chunks)
	// Output: [[1 2] [3 4] [5]]
}

func ExampleMinMax() {
	values := []int{5, 2, 9, 1}
	min, max := seq.MinMax(values)
	fmt.Println(min, max)
	// Output: 1 9
}

func ExampleMinMaxFunc() {
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 40},
	}

	less := func(a, b Person) int {
		return compare(a.Age, b.Age)
	}

	min, max := seq.MinMaxFunc(people, less)
	fmt.Println(min, max)
	// Output: {Bob 25} {Charlie 40}
}

// helper — replaces cmp.Compare but avoids extra imports
func compare(a, b int) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}
