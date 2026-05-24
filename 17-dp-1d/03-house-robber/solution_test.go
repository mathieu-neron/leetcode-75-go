package solution

import (
	"testing"
)

func TestRob(t *testing.T) {
	t.Skip("TODO: implement rob")
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 1}, 4},
		{[]int{2, 7, 9, 3, 1}, 12},
	}
	for _, tc := range tests {
		got := rob(tc.nums)
		if got != tc.want {
			t.Errorf("rob(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
