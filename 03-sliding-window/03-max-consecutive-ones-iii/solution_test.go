package solution

import (
	"testing"
)

func TestLongestOnes(t *testing.T) {
	t.Skip("TODO: implement longestOnes")
	tests := []struct {
		nums []int
		k int
		want int
	}{
		{[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2, 6},
		{[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3, 10},
	}
	for _, tc := range tests {
		got := longestOnes(tc.nums, tc.k)
		if got != tc.want {
			t.Errorf("longestOnes(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
