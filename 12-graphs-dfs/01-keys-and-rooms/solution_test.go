package solution

import (
	"testing"
)

func TestCanVisitAllRooms(t *testing.T) {
	t.Skip("TODO: implement canVisitAllRooms")
	tests := []struct {
		rooms [][]int
		want bool
	}{
		{[][]int{[]int{1}, []int{2}, []int{3}, []int{}}, true},
		{[][]int{[]int{1, 3}, []int{3, 0, 1}, []int{2}, []int{0}}, false},
	}
	for _, tc := range tests {
		got := canVisitAllRooms(tc.rooms)
		if got != tc.want {
			t.Errorf("canVisitAllRooms(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
