package seq

import (
	"cmp"
	"slices"
	"testing"
)

func TestFilter(t *testing.T) {
	t.Run("keep evens", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5, 6}
		got := Filter(in, func(v int) bool { return v%2 == 0 })
		want := []int{2, 4, 6}

		if !slices.Equal(got, want) {
			t.Fatalf("Filter() = %v, want %v", got, want)
		}
	})

	t.Run("keep none", func(t *testing.T) {
		in := []int{1, 3, 5}
		got := Filter(in, func(v int) bool { return v%2 == 0 })
		if len(got) != 0 {
			t.Fatalf("Filter() = %v, want empty slice", got)
		}
	})

	t.Run("empty input", func(t *testing.T) {
		var in []int
		got := Filter(in, func(v int) bool { return v > 0 })
		if len(got) != 0 {
			t.Fatalf("Filter() = %v, want empty slice", got)
		}
	})
}

func TestMap(t *testing.T) {
	t.Run("int to string", func(t *testing.T) {
		in := []int{1, 2, 3}
		got := Map(in, func(v int) string { return "n=" + string(rune('0'+v)) })
		want := []string{"n=1", "n=2", "n=3"}

		if !slices.Equal(got, want) {
			t.Fatalf("Map() = %v, want %v", got, want)
		}
	})

	t.Run("empty input", func(t *testing.T) {
		var in []int
		got := Map(in, func(v int) int { return v * 2 })
		if len(got) != 0 {
			t.Fatalf("Map() = %v, want empty slice", got)
		}
	})
}

func TestReduce(t *testing.T) {
	t.Run("sum ints", func(t *testing.T) {
		in := []int{1, 2, 3, 4}
		got := Reduce(in, 0, func(acc, v int) int { return acc + v })
		const want = 10
		if got != want {
			t.Fatalf("Reduce() = %d, want %d", got, want)
		}
	})

	t.Run("concat strings", func(t *testing.T) {
		in := []string{"a", "b", "c"}
		got := Reduce(in, "", func(acc, v string) string { return acc + v })
		const want = "abc"
		if got != want {
			t.Fatalf("Reduce() = %q, want %q", got, want)
		}
	})

	t.Run("empty slice returns init", func(t *testing.T) {
		var in []int
		got := Reduce(in, 42, func(acc, v int) int { return acc + v })
		const want = 42
		if got != want {
			t.Fatalf("Reduce() = %d, want %d", got, want)
		}
	})
}

func TestLastIndex(t *testing.T) {
	t.Run("value found", func(t *testing.T) {
		in := []int{1, 2, 3, 2, 4}
		idx, ok := LastIndex(in, 2)
		if !ok {
			t.Fatalf("LastIndex() = (_, false), want true")
		}
		if idx != 3 {
			t.Fatalf("LastIndex() idx = %d, want 3", idx)
		}
	})

	t.Run("value not found", func(t *testing.T) {
		in := []int{1, 2, 3}
		idx, ok := LastIndex(in, 4)
		if ok {
			t.Fatalf("LastIndex() = (_, true), want false")
		}
		if idx != -1 {
			t.Fatalf("LastIndex() idx = %d, want -1", idx)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		var in []int
		idx, ok := LastIndex(in, 1)
		if ok {
			t.Fatalf("LastIndex() = (_, true), want false")
		}
		if idx != -1 {
			t.Fatalf("LastIndex() idx = %d, want -1", idx)
		}
	})
}

func TestPartition(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 6}
	matches, nonMatches := Partition(in, func(v int) bool { return v%2 == 0 })

	wantMatches := []int{2, 4, 6}
	wantNonMatches := []int{1, 3, 5}

	if !slices.Equal(matches, wantMatches) {
		t.Fatalf("Partition() matches = %v, want %v", matches, wantMatches)
	}
	if !slices.Equal(nonMatches, wantNonMatches) {
		t.Fatalf("Partition() nonMatches = %v, want %v", nonMatches, wantNonMatches)
	}
}

func TestGroupBy(t *testing.T) {
	in := []string{"a", "bb", "ccc", "dd", "e"}
	got := GroupBy(in, func(s string) int { return len(s) })

	// expected:
	// 1 -> ["a", "e"]
	// 2 -> ["bb", "dd"]
	// 3 -> ["ccc"]
	if !slices.Equal(got[1], []string{"a", "e"}) {
		t.Fatalf("GroupBy()[1] = %v, want %v", got[1], []string{"a", "e"})
	}
	if !slices.Equal(got[2], []string{"bb", "dd"}) {
		t.Fatalf("GroupBy()[2] = %v, want %v", got[2], []string{"bb", "dd"})
	}
	if !slices.Equal(got[3], []string{"ccc"}) {
		t.Fatalf("GroupBy()[3] = %v, want %v", got[3], []string{"ccc"})
	}
}

func TestFlatten(t *testing.T) {
	in := [][]int{
		{1, 2},
		{},
		{3},
		{4, 5},
	}
	got := Flatten(in)
	want := []int{1, 2, 3, 4, 5}

	if !slices.Equal(got, want) {
		t.Fatalf("Flatten() = %v, want %v", got, want)
	}
}

func TestUnique(t *testing.T) {
	t.Run("deduplicates and preserves order", func(t *testing.T) {
		in := []int{1, 2, 1, 3, 2, 4, 4}
		got := Unique(in)
		want := []int{1, 2, 3, 4}
		if !slices.Equal(got, want) {
			t.Fatalf("Unique() = %v, want %v", got, want)
		}
	})

	t.Run("already unique", func(t *testing.T) {
		in := []int{1, 2, 3}
		got := Unique(in)
		if !slices.Equal(got, in) {
			t.Fatalf("Unique() = %v, want %v", got, in)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		var in []int
		got := Unique(in)
		if len(got) != 0 {
			t.Fatalf("Unique() = %v, want empty slice", got)
		}
	})
}

func TestUniqueBy(t *testing.T) {
	type user struct {
		Name  string
		Email string
	}

	in := []user{
		{Name: "Alice", Email: "a@example.com"},
		{Name: "Bob", Email: "b@example.com"},
		{Name: "Alice Duplicate", Email: "a@example.com"},
		{Name: "Charlie", Email: "c@example.com"},
	}

	got := UniqueBy(in, func(u user) string { return u.Email })
	want := []user{
		{Name: "Alice", Email: "a@example.com"},
		{Name: "Bob", Email: "b@example.com"},
		{Name: "Charlie", Email: "c@example.com"},
	}

	if len(got) != len(want) {
		t.Fatalf("UniqueBy() len = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("UniqueBy()[%d] = %#v, want %#v", i, got[i], want[i])
		}
	}
}

func TestChunk(t *testing.T) {
	t.Run("exact division", func(t *testing.T) {
		in := []int{1, 2, 3, 4}
		got := Chunk(in, 2)
		want := [][]int{{1, 2}, {3, 4}}

		if len(got) != len(want) {
			t.Fatalf("Chunk() len = %d, want %d", len(got), len(want))
		}
		for i := range want {
			if !slices.Equal(got[i], want[i]) {
				t.Fatalf("Chunk()[%d] = %v, want %v", i, got[i], want[i])
			}
		}
	})

	t.Run("last chunk smaller", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5}
		got := Chunk(in, 2)
		want := [][]int{{1, 2}, {3, 4}, {5}}

		if len(got) != len(want) {
			t.Fatalf("Chunk() len = %d, want %d", len(got), len(want))
		}
		for i := range want {
			if !slices.Equal(got[i], want[i]) {
				t.Fatalf("Chunk()[%d] = %v, want %v", i, got[i], want[i])
			}
		}
	})

	t.Run("size larger than slice", func(t *testing.T) {
		in := []int{1, 2, 3}
		got := Chunk(in, 10)
		want := [][]int{{1, 2, 3}}

		if len(got) != len(want) {
			t.Fatalf("Chunk() len = %d, want %d", len(got), len(want))
		}
		if !slices.Equal(got[0], want[0]) {
			t.Fatalf("Chunk()[0] = %v, want %v", got[0], want[0])
		}
	})
}

func TestChunkPrecondition(t *testing.T) {
	// The docs say: caller must ensure size > 0.
	// Here we just assert that using a valid size works; we do NOT test size <= 0
	// because that would loop indefinitely.
	in := []int{1, 2, 3}
	_ = Chunk(in, 1)
}

func TestMinMax(t *testing.T) {
	t.Run("ints", func(t *testing.T) {
		in := []int{5, 2, 9, 1, 7}
		min, max := MinMax(in)

		if min != 1 || max != 9 {
			t.Fatalf("MinMax() = (%d, %d), want (1, 9)", min, max)
		}
	})

	t.Run("floats", func(t *testing.T) {
		in := []float64{1.5, 2.5, -0.5, 3.0}
		min, max := MinMax(in)

		if min != -0.5 || max != 3.0 {
			t.Fatalf("MinMax() = (%f, %f), want (-0.5, 3.0)", min, max)
		}
	})
}

func TestMinMaxEmptyPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("MinMax() did not panic on empty slice")
		}
	}()

	var in []int
	MinMax(in)
}

func TestMinMaxFunc(t *testing.T) {
	type person struct {
		Name string
		Age  int
	}

	in := []person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 40},
	}

	less := func(a, b person) int {
		return cmp.Compare(a.Age, b.Age)
	}

	min, max := MinMaxFunc(in, less)

	if min.Name != "Bob" || max.Name != "Charlie" {
		t.Fatalf("MinMaxFunc() = (%v, %v), want (Bob, Charlie)", min, max)
	}
}

func TestMinMaxFuncEmptyPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("MinMaxFunc() did not panic on empty slice")
		}
	}()

	var in []int
	MinMaxFunc(in, func(a, b int) int {
		return cmp.Compare(a, b)
	})
}
