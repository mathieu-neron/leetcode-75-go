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

func TestDeleteMiddle(t *testing.T) {
	t.Skip("TODO: implement deleteMiddle")
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{1, 3, 4, 7, 1, 2, 6}, []int{1, 3, 4, 1, 2, 6}},
		{[]int{1, 2, 3, 4}, []int{1, 2, 4}},
		{[]int{2, 1}, []int{2}},
	}
	for _, tc := range tests {
		got := deleteMiddle(buildList(tc.input))
		if !reflect.DeepEqual(listToArr(got), tc.want) {
			t.Errorf("deleteMiddle(%v) = %v, want %v", tc.input, listToArr(got), tc.want)
		}
	}
}
