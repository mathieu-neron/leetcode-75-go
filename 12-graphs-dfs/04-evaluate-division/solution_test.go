package solution

import (
	"testing"
	"reflect"
)

func TestCalcEquation(t *testing.T) {
	t.Skip("TODO: implement calcEquation")
	tests := []struct {
		equations [][]string
		values []float64
		queries [][]string
		want []float64
	}{
		{[][]string{[]string{"a", "b"}, []string{"b", "c"}}, []float64{2.0, 3.0}, [][]string{[]string{"a", "c"}, []string{"b", "a"}, []string{"a", "e"}, []string{"a", "a"}, []string{"x", "x"}}, []float64{6.0, 0.5, -1.0, 1.0, -1.0}},
	}
	for _, tc := range tests {
		got := calcEquation(tc.equations, tc.values, tc.queries)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("calcEquation(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
