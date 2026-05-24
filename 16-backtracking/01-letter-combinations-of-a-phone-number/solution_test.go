package solution

import (
	"reflect"
	"sort"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	t.Skip("TODO: implement letterCombinations")
	tests := []struct {
		digits string
		want []string
	}{
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"", []string{}},
		{"2", []string{"a", "b", "c"}},
	}
	for _, tc := range tests {
		got := letterCombinations(tc.digits)
		sort.Strings(got)
		sort.Strings(tc.want)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("letterCombinations(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
