package solution

import (
	"reflect"
	"testing"
)

func buildList(arr []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range arr {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func listToArr(head *ListNode) []int {
	var out []int
	for n := head; n != nil; n = n.Next {
		out = append(out, n.Val)
	}
	return out
}

func TestReverseList(t *testing.T) {
	t.Skip("TODO: implement reverseList")
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{}, []int{}},
	}
	for _, tc := range tests {
		got := reverseList(buildList(tc.input))
		if !reflect.DeepEqual(listToArr(got), tc.want) {
			t.Errorf("reverseList(%v) = %v, want %v", tc.input, listToArr(got), tc.want)
		}
	}
}
