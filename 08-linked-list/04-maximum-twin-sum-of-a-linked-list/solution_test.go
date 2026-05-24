package solution

import "testing"

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

func TestPairSum(t *testing.T) {
	t.Skip("TODO: implement pairSum")
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{5, 4, 2, 1}, 6},
		{[]int{4, 2, 2, 3}, 7},
		{[]int{1, 100000}, 100001},
	}
	for _, tc := range tests {
		got := pairSum(buildList(tc.input))
		if got != tc.want {
			t.Errorf("pairSum(%v) = %v, want %v", tc.input, got, tc.want)
		}
	}
}
