package solution

import (
	"testing"
)

func TestLongestSubarray(t *testing.T) {
	t.Skip("TODO: implement longestSubarray")
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 1, 0, 1}, 3},
		{[]int{0, 1, 1, 1, 0, 1, 1, 0, 1}, 5},
		{[]int{1, 1, 1}, 2},
	}
	for _, tc := range tests {
		got := longestSubarray(tc.nums)
		if got != tc.want {
			t.Errorf("longestSubarray(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
