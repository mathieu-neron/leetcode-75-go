package solution

import (
	"testing"
)

func TestEqualPairs(t *testing.T) {
	t.Skip("TODO: implement equalPairs")
	tests := []struct {
		grid [][]int
		want int
	}{
		{[][]int{[]int{3, 2, 1}, []int{1, 7, 6}, []int{2, 7, 7}}, 1},
		{[][]int{[]int{3, 1, 2, 2}, []int{1, 4, 4, 5}, []int{2, 4, 2, 2}, []int{2, 4, 2, 2}}, 3},
	}
	for _, tc := range tests {
		got := equalPairs(tc.grid)
		if got != tc.want {
			t.Errorf("equalPairs(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
