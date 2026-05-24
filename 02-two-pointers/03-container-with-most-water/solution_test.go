package solution

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
	t.Skip("TODO: implement maxArea")
	tests := []struct {
		height []int
		want int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
	}
	for _, tc := range tests {
		got := maxArea(tc.height)
		if got != tc.want {
			t.Errorf("maxArea(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
