package solution

import "testing"

func TestSmallestInfiniteSet(t *testing.T) {
	t.Skip("TODO: implement SmallestInfiniteSet")
	obj := Constructor()
	obj.AddBack(2)
	if got := obj.PopSmallest(); got != 1 {
		t.Errorf("PopSmallest(%v) = %v, want %v", []any{}, got, 1)
	}
	if got := obj.PopSmallest(); got != 2 {
		t.Errorf("PopSmallest(%v) = %v, want %v", []any{}, got, 2)
	}
	if got := obj.PopSmallest(); got != 3 {
		t.Errorf("PopSmallest(%v) = %v, want %v", []any{}, got, 3)
	}
	obj.AddBack(1)
	if got := obj.PopSmallest(); got != 1 {
		t.Errorf("PopSmallest(%v) = %v, want %v", []any{}, got, 1)
	}
	if got := obj.PopSmallest(); got != 4 {
		t.Errorf("PopSmallest(%v) = %v, want %v", []any{}, got, 4)
	}
	if got := obj.PopSmallest(); got != 5 {
		t.Errorf("PopSmallest(%v) = %v, want %v", []any{}, got, 5)
	}
}
