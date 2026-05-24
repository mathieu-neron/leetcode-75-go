package solution

import (
	"testing"
	"reflect"
)

func TestFindDifference(t *testing.T) {
	t.Skip("TODO: implement findDifference")
	tests := []struct {
		nums1 []int
		nums2 []int
		want [][]int
	}{
		{[]int{1, 2, 3}, []int{2, 4, 6}, [][]int{[]int{1, 3}, []int{4, 6}}},
		{[]int{1, 2, 3, 3}, []int{1, 1, 2, 2}, [][]int{[]int{3}, []int{}}},
	}
	for _, tc := range tests {
		got := findDifference(tc.nums1, tc.nums2)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("findDifference(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
