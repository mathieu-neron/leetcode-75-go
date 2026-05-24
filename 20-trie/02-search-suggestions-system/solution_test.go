package solution

import (
	"testing"
	"reflect"
)

func TestSuggestedProducts(t *testing.T) {
	t.Skip("TODO: implement suggestedProducts")
	tests := []struct {
		products []string
		searchWord string
		want [][]string
	}{
		{[]string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}, "mouse", [][]string{[]string{"mobile", "moneypot", "monitor"}, []string{"mobile", "moneypot", "monitor"}, []string{"mouse", "mousepad"}, []string{"mouse", "mousepad"}, []string{"mouse", "mousepad"}}},
		{[]string{"havana"}, "havana", [][]string{[]string{"havana"}, []string{"havana"}, []string{"havana"}, []string{"havana"}, []string{"havana"}, []string{"havana"}}},
	}
	for _, tc := range tests {
		got := suggestedProducts(tc.products, tc.searchWord)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("suggestedProducts(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
