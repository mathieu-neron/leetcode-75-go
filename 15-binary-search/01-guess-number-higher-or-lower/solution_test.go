package solution

import "testing"

func TestGuessNumber(t *testing.T) {
	t.Skip("TODO: implement guessNumber (and stub guess(pick))")
	tests := []struct{ n, pick int }{
		{10, 6},
		{1, 1},
		{2, 1},
	}
	for _, tc := range tests {
		// Replace the package-level guess with a closure for the test.
		// Idiomatic LeetCode harness; you may need to adjust how guess is wired.
		_ = tc
		// got := guessNumber(tc.n)
		// if got != tc.pick { t.Errorf("got %d, want %d", got, tc.pick) }
	}
}
