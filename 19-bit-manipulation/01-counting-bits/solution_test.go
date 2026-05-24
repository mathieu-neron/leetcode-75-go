package solution

import (
	"testing"
	"reflect"
)

func TestCountBits(t *testing.T) {
	t.Skip("TODO: implement countBits")
	tests := []struct {
		n int
		want []int
	}{
		{2, []int{0, 1, 1}},
		{5, []int{0, 1, 1, 2, 1, 2}},
	}
	for _, tc := range tests {
		got := countBits(tc.n)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("countBits(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
