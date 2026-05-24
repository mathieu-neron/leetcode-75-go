package solution

import (
	"testing"
)

func TestMinDistance(t *testing.T) {
	t.Skip("TODO: implement minDistance")
	tests := []struct {
		word1 string
		word2 string
		want int
	}{
		{"horse", "ros", 3},
		{"intention", "execution", 5},
	}
	for _, tc := range tests {
		got := minDistance(tc.word1, tc.word2)
		if got != tc.want {
			t.Errorf("minDistance(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
