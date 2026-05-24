package solution

import (
	"testing"
)

func TestUniqueOccurrences(t *testing.T) {
	t.Skip("TODO: implement uniqueOccurrences")
	tests := []struct {
		arr []int
		want bool
	}{
		{[]int{1, 2, 2, 1, 1, 3}, true},
		{[]int{1, 2}, false},
		{[]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}, true},
	}
	for _, tc := range tests {
		got := uniqueOccurrences(tc.arr)
		if got != tc.want {
			t.Errorf("uniqueOccurrences(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
