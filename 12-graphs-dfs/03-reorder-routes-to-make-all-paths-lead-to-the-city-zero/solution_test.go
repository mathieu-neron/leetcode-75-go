package solution

import (
	"testing"
)

func TestMinReorder(t *testing.T) {
	t.Skip("TODO: implement minReorder")
	tests := []struct {
		n int
		connections [][]int
		want int
	}{
		{6, [][]int{[]int{0, 1}, []int{1, 3}, []int{2, 3}, []int{4, 0}, []int{4, 5}}, 3},
		{5, [][]int{[]int{1, 0}, []int{1, 2}, []int{3, 2}, []int{3, 4}}, 2},
		{3, [][]int{[]int{1, 0}, []int{2, 0}}, 0},
	}
	for _, tc := range tests {
		got := minReorder(tc.n, tc.connections)
		if got != tc.want {
			t.Errorf("minReorder(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
