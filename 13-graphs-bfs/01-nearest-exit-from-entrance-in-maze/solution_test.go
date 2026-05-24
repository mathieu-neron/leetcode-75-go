package solution

import (
	"testing"
)

func TestNearestExit(t *testing.T) {
	t.Skip("TODO: implement nearestExit")
	tests := []struct {
		maze [][]byte
		entrance []int
		want int
	}{
		{[][]byte{[]byte{'+', '+', '.', '+'}, []byte{'.', '.', '.', '+'}, []byte{'+', '+', '+', '.'}}, []int{1, 2}, 1},
		{[][]byte{[]byte{'+', '+', '+'}, []byte{'.', '.', '.'}, []byte{'+', '+', '+'}}, []int{1, 0}, 2},
		{[][]byte{[]byte{'.', '+'}}, []int{0, 0}, -1},
	}
	for _, tc := range tests {
		got := nearestExit(tc.maze, tc.entrance)
		if got != tc.want {
			t.Errorf("nearestExit(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
