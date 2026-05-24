package solution

import (
	"testing"
)

func TestOrangesRotting(t *testing.T) {
	t.Skip("TODO: implement orangesRotting")
	tests := []struct {
		grid [][]int
		want int
	}{
		{[][]int{[]int{2, 1, 1}, []int{1, 1, 0}, []int{0, 1, 1}}, 4},
		{[][]int{[]int{2, 1, 1}, []int{0, 1, 1}, []int{1, 0, 1}}, -1},
		{[][]int{[]int{0, 2}}, 0},
	}
	for _, tc := range tests {
		got := orangesRotting(tc.grid)
		if got != tc.want {
			t.Errorf("orangesRotting(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
