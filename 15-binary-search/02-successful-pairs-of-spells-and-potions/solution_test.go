package solution

import (
	"testing"
	"reflect"
)

func TestSuccessfulPairs(t *testing.T) {
	t.Skip("TODO: implement successfulPairs")
	tests := []struct {
		spells []int
		potions []int
		success int64
		want []int
	}{
		{[]int{5, 1, 3}, []int{1, 2, 3, 4, 5}, 7, []int{4, 0, 3}},
		{[]int{3, 1, 2}, []int{8, 5, 8}, 16, []int{2, 0, 2}},
	}
	for _, tc := range tests {
		got := successfulPairs(tc.spells, tc.potions, tc.success)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("successfulPairs(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
