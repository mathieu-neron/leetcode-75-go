package solution

import "testing"

func TestRecentCounter(t *testing.T) {
	t.Skip("TODO: implement RecentCounter")
	obj := Constructor()
	if got := obj.Ping(1); got != 1 {
		t.Errorf("Ping(%v) = %v, want %v", []any{1}, got, 1)
	}
	if got := obj.Ping(100); got != 2 {
		t.Errorf("Ping(%v) = %v, want %v", []any{100}, got, 2)
	}
	if got := obj.Ping(3001); got != 3 {
		t.Errorf("Ping(%v) = %v, want %v", []any{3001}, got, 3)
	}
	if got := obj.Ping(3002); got != 3 {
		t.Errorf("Ping(%v) = %v, want %v", []any{3002}, got, 3)
	}
}
