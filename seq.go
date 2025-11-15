// Package seq provides generic helpers for working with slices and other
// sequences. It includes several functions that are not included in the
// standard library for working with slices which I have found helpful.
// It includes functions for mapping, filtering, grouping,
// chunking, de-duplicating, and computing aggregate values such as minima
// and maxima.
package seq

import (
	"cmp"
	"slices"
)

// Filter returns a new slice containing only the elements of slice for which
// keep returns true.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5, 6}
//	evens := Filter(numbers, func(n int) bool { return n%2 == 0 })
//	// evens == []int{2, 4, 6}
func Filter[T any](slice []T, keep func(T) bool) []T {
	out := make([]T, 0, len(slice))
	for _, v := range slice {
		if keep(v) {
			out = append(out, v)
		}
	}
	return out
}

// Map returns a new slice containing the results of applying f to each
// element of slice.
//
// Example:
//
//	numbers := []int{1, 2, 3}
//	strs := Map(numbers, func(n int) string {
//		return fmt.Sprintf("n=%d", n)
//	})
//	// strs == []string{"n=1", "n=2", "n=3"}
func Map[T any, R any](slice []T, f func(T) R) []R {
	out := make([]R, 0, len(slice))
	for _, v := range slice {
		out = append(out, f(v))
	}
	return out
}

// Reduce applies f to each element of slice, accumulating the result, and
// returns the final accumulated value. The accumulator is initialized with init.
//
// Example (sum):
//
//	numbers := []int{1, 2, 3, 4}
//	sum := Reduce(numbers, 0, func(acc, n int) int {
//		return acc + n
//	})
//	// sum == 10
//
// Example (concatenate):
//
//	words := []string{"go", " ", "lang"}
//	joined := Reduce(words, "", func(acc, s string) string {
//		return acc + s
//	})
//	// joined == "go lang"
func Reduce[T any, R any](slice []T, init R, f func(R, T) R) R {
	acc := init
	for _, v := range slice {
		acc = f(acc, v)
	}
	return acc
}

// LastIndex returns the index of the last occurrence of v in slice.
// If v is not found, the returned index is -1 and ok is false.
//
// Example:
//
//	values := []string{"a", "b", "c", "b"}
//	idx, ok := LastIndex(values, "b")
//	// idx == 3, ok == true
//
//	_, ok = LastIndex(values, "z")
//	// ok == false
func LastIndex[T comparable](slice []T, v T) (idx int, ok bool) {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == v {
			return i, true
		}
	}
	return -1, false
}

// Partition splits slice into two slices: matches, containing elements for
// which pred returns true, and nonMatches, containing the rest.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5, 6}
//	evens, odds := Partition(numbers, func(n int) bool { return n%2 == 0 })
//	// evens == []int{2, 4, 6}
//	// odds  == []int{1, 3, 5}
func Partition[T any](slice []T, pred func(T) bool) (matches, nonMatches []T) {
	matches = make([]T, 0, len(slice))
	nonMatches = make([]T, 0, len(slice))

	for _, v := range slice {
		if pred(v) {
			matches = append(matches, v)
		} else {
			nonMatches = append(nonMatches, v)
		}
	}
	return
}

// GroupBy groups the elements of slice into a map keyed by the value returned
// from keyFunc for each element.
//
// Example:
//
//	words := []string{"a", "bb", "ccc", "dd"}
//	byLen := GroupBy(words, func(s string) int { return len(s) })
//	// byLen[1] == []string{"a"}
//	// byLen[2] == []string{"bb", "dd"}
//	// byLen[3] == []string{"ccc"}
func GroupBy[T any, K comparable](slice []T, keyFunc func(T) K) map[K][]T {
	out := make(map[K][]T)

	for _, v := range slice {
		k := keyFunc(v)
		out[k] = append(out[k], v)
	}
	return out
}

// Flatten returns a new slice containing all the elements of slices
// flattened into a single slice.
//
// Example:
//
//	nested := [][]int{{1, 2}, {}, {3}, {4, 5}}
//	flat := Flatten(nested)
//	// flat == []int{1, 2, 3, 4, 5}
func Flatten[T any](slices [][]T) []T {
	var out []T
	for _, inner := range slices {
		out = append(out, inner...)
	}
	return out
}

// Unique returns a new slice containing only the unique elements of slice.
// The order of first occurrence is preserved.
//
// Example:
//
//	values := []int{1, 2, 1, 3, 2, 4, 4}
//	uniq := Unique(values)
//	// uniq == []int{1, 2, 3, 4}
func Unique[T comparable](slice []T) []T {
	unique := make([]T, 0, len(slice))
	seen := make(map[T]struct{})

	for _, item := range slice {
		if _, ok := seen[item]; !ok {
			seen[item] = struct{}{}
			unique = append(unique, item)
		}
	}

	return unique
}

// UniqueBy returns a new slice containing only the unique elements of slice,
// where uniqueness is determined by the key returned from keyFunc.
// The order of first occurrence (by key) is preserved.
//
// Example:
//
//	type User struct {
//		Name  string
//		Email string
//	}
//
//	users := []User{
//		{Name: "Alice", Email: "a@example.com"},
//		{Name: "Bob", Email: "b@example.com"},
//		{Name: "Alice Clone", Email: "a@example.com"},
//	}
//
//	uniqByEmail := UniqueBy(users, func(u User) string { return u.Email })
//	// uniqByEmail == []User{
//	//   {Name: "Alice", Email: "a@example.com"},
//	//   {Name: "Bob",   Email: "b@example.com"},
//	// }
func UniqueBy[T any, K comparable](slice []T, keyFunc func(T) K) []T {
	unique := make([]T, 0, len(slice))
	seen := make(map[K]struct{})

	for _, item := range slice {
		k := keyFunc(item)
		if _, ok := seen[k]; !ok {
			seen[k] = struct{}{}
			unique = append(unique, item)
		}
	}

	return unique
}

// Chunk splits slice into consecutive sub-slices of at most size elements.
// The final chunk may be smaller than size.
// The caller must ensure size > 0.
//
// Example:
//
//	values := []int{1, 2, 3, 4, 5}
//	chunks := Chunk(values, 2)
//	// chunks == [][]int{{1, 2}, {3, 4}, {5}}
func Chunk[T any](slice []T, size int) [][]T {
	var chunks [][]T
	for i := 0; i < len(slice); i += size {
		end := min(i+size, len(slice))
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

// MinMax returns the minimum and maximum values of slice.
// The caller must ensure that slice is non-empty otherwise MinMax will panic.
//
// Example:
//
//	values := []int{5, 2, 9, 1}
//	min, max := MinMax(values)
//	// min == 1
//	// max == 9
func MinMax[T cmp.Ordered](slice []T) (min, max T) {
	return slices.Min(slice), slices.Max(slice)
}

// MinMaxFunc returns the minimum and maximum values of slice using the
// comparison function less. The comparison function should return a negative
// value if a < b, zero if a == b, and a positive value if a > b.
// The caller must ensure that slice is non-empty otherwise MinMaxFunc will panic.
//
// Example:
//
//	type Person struct {
//		Name string
//		Age  int
//	}
//
//	people := []Person{
//		{Name: "Alice", Age: 30},
//		{Name: "Bob", Age: 25},
//		{Name: "Charlie", Age: 40},
//	}
//
//	lessByAge := func(a, b Person) int {
//		return cmp.Compare(a.Age, b.Age)
//	}
//
//	min, max := MinMaxFunc(people, lessByAge)
//	// min == Person{Name: "Bob", Age: 25}
//	// max == Person{Name: "Charlie", Age: 40}
func MinMaxFunc[T any](slice []T, less func(T, T) int) (min, max T) {
	return slices.MinFunc(slice, less), slices.MaxFunc(slice, less)
}
