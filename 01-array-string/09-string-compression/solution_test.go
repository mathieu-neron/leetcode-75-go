package solution

import (
	"reflect"
	"testing"
)

func TestCompress(t *testing.T) {
	t.Skip("TODO: implement compress")
	tests := []struct {
		input      []byte
		wantLen    int
		wantPrefix []byte
	}{
		{input: []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}, wantLen: 6, wantPrefix: []byte{'a', '2', 'b', '2', 'c', '3'}},
		{input: []byte{'a'}, wantLen: 1, wantPrefix: []byte{'a'}},
		{input: []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}, wantLen: 4, wantPrefix: []byte{'a', 'b', '1', '2'}},
	}
	for _, tc := range tests {
		gotLen := compress(tc.input)
		if gotLen != tc.wantLen {
			t.Errorf("len: got %d, want %d", gotLen, tc.wantLen)
		}
		if !reflect.DeepEqual(tc.input[:gotLen], tc.wantPrefix) {
			t.Errorf("prefix: got %v, want %v", tc.input[:gotLen], tc.wantPrefix)
		}
	}
}
