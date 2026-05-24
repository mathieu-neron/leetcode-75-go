package solution

import "testing"

func TestStockSpanner(t *testing.T) {
	t.Skip("TODO: implement StockSpanner")
	obj := Constructor()
	if got := obj.Next(100); got != 1 {
		t.Errorf("Next(%v) = %v, want %v", []any{100}, got, 1)
	}
	if got := obj.Next(80); got != 1 {
		t.Errorf("Next(%v) = %v, want %v", []any{80}, got, 1)
	}
	if got := obj.Next(60); got != 1 {
		t.Errorf("Next(%v) = %v, want %v", []any{60}, got, 1)
	}
	if got := obj.Next(70); got != 2 {
		t.Errorf("Next(%v) = %v, want %v", []any{70}, got, 2)
	}
	if got := obj.Next(60); got != 1 {
		t.Errorf("Next(%v) = %v, want %v", []any{60}, got, 1)
	}
	if got := obj.Next(75); got != 4 {
		t.Errorf("Next(%v) = %v, want %v", []any{75}, got, 4)
	}
	if got := obj.Next(85); got != 6 {
		t.Errorf("Next(%v) = %v, want %v", []any{85}, got, 6)
	}
}
