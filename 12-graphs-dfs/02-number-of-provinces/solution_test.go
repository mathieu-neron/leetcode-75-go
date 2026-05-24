package solution

import (
	"testing"
)

func TestFindCircleNum(t *testing.T) {
	t.Skip("TODO: implement findCircleNum")
	tests := []struct {
		isConnected [][]int
		want int
	}{
		{[][]int{[]int{1, 1, 0}, []int{1, 1, 0}, []int{0, 0, 1}}, 2},
		{[][]int{[]int{1, 0, 0}, []int{0, 1, 0}, []int{0, 0, 1}}, 3},
	}
	for _, tc := range tests {
		got := findCircleNum(tc.isConnected)
		if got != tc.want {
			t.Errorf("findCircleNum(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
