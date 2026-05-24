package solution

import (
	"testing"
)

func TestPredictPartyVictory(t *testing.T) {
	t.Skip("TODO: implement predictPartyVictory")
	tests := []struct {
		senate string
		want string
	}{
		{"RD", "Radiant"},
		{"RDD", "Dire"},
	}
	for _, tc := range tests {
		got := predictPartyVictory(tc.senate)
		if got != tc.want {
			t.Errorf("predictPartyVictory(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
