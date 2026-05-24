package solution

import (
	"reflect"
	"sort"
	"testing"
)

func normalize2D(xs [][]int) [][]int {
	out := make([][]int, len(xs))
	for i, x := range xs {
		cp := append([]int{}, x...)
		sort.Ints(cp)
		out[i] = cp
	}
	sort.Slice(out, func(i, j int) bool {
		for k := 0; k < len(out[i]) && k < len(out[j]); k++ {
			if out[i][k] != out[j][k] { return out[i][k] < out[j][k] }
		}
		return len(out[i]) < len(out[j])
	})
	return out
}

func TestCombinationSum3(t *testing.T) {
	t.Skip("TODO: implement combinationSum3")
	tests := []struct {
		k int
		n int
		want [][]int
	}{
		{3, 7, [][]int{[]int{1, 2, 4}}},
		{3, 9, [][]int{[]int{1, 2, 6}, []int{1, 3, 5}, []int{2, 3, 4}}},
		{4, 1, [][]int{}},
	}
	for _, tc := range tests {
		got := combinationSum3(tc.k, tc.n)
		if !reflect.DeepEqual(normalize2D(got), normalize2D(tc.want)) {
			t.Errorf("combinationSum3(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
