package solution

import "testing"

func TestTrie(t *testing.T) {
	t.Skip("TODO: implement Trie")
	obj := Constructor()
	obj.Insert("apple")
	if got := obj.Search("apple"); got != true {
		t.Errorf("Search(%v) = %v, want %v", []any{"apple"}, got, true)
	}
	if got := obj.Search("app"); got != false {
		t.Errorf("Search(%v) = %v, want %v", []any{"app"}, got, false)
	}
	if got := obj.StartsWith("app"); got != true {
		t.Errorf("StartsWith(%v) = %v, want %v", []any{"app"}, got, true)
	}
	obj.Insert("app")
	if got := obj.Search("app"); got != true {
		t.Errorf("Search(%v) = %v, want %v", []any{"app"}, got, true)
	}
}
