package solution

import (
	"testing"
	"reflect"
)

func TestKidsWithCandies(t *testing.T) {
	t.Skip("TODO: implement kidsWithCandies")
	tests := []struct {
		candies []int
		extraCandies int
		want []bool
	}{
		{[]int{2, 3, 5, 1, 3}, 3, []bool{true, true, true, false, true}},
		{[]int{4, 2, 1, 1, 2}, 1, []bool{true, false, false, false, false}},
		{[]int{12, 1, 12}, 10, []bool{true, false, true}},
	}
	for _, tc := range tests {
		got := kidsWithCandies(tc.candies, tc.extraCandies)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("kidsWithCandies(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
