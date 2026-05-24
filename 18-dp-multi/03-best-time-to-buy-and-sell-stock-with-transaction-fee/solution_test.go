package solution

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	t.Skip("TODO: implement maxProfit")
	tests := []struct {
		prices []int
		fee int
		want int
	}{
		{[]int{1, 3, 2, 8, 4, 9}, 2, 8},
		{[]int{1, 3, 7, 5, 10, 3}, 3, 6},
	}
	for _, tc := range tests {
		got := maxProfit(tc.prices, tc.fee)
		if got != tc.want {
			t.Errorf("maxProfit(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
