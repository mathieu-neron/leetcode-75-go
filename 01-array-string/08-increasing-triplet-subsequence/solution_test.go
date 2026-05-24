package solution

import (
	"testing"
)

func TestIncreasingTriplet(t *testing.T) {
	t.Skip("TODO: implement increasingTriplet")
	tests := []struct {
		nums []int
		want bool
	}{
		{[]int{1, 2, 3, 4, 5}, true},
		{[]int{5, 4, 3, 2, 1}, false},
		{[]int{2, 1, 5, 0, 4, 6}, true},
	}
	for _, tc := range tests {
		got := increasingTriplet(tc.nums)
		if got != tc.want {
			t.Errorf("increasingTriplet(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
