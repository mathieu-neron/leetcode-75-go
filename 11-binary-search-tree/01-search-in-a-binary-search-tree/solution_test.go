package solution

import (
	"reflect"
	"testing"
)

func buildTree(arr []any) *TreeNode {
	if len(arr) == 0 || arr[0] == nil {
		return nil
	}
	root := &TreeNode{Val: arr[0].(int)}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(arr) {
		node := queue[0]
		queue = queue[1:]
		if i < len(arr) && arr[i] != nil {
			node.Left = &TreeNode{Val: arr[i].(int)}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(arr) && arr[i] != nil {
			node.Right = &TreeNode{Val: arr[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

func levelOrder(root *TreeNode) []any {
	if root == nil {
		return nil
	}
	var out []any
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		if n == nil {
			out = append(out, nil)
			continue
		}
		out = append(out, n.Val)
		queue = append(queue, n.Left, n.Right)
	}
	for len(out) > 0 && out[len(out)-1] == nil {
		out = out[:len(out)-1]
	}
	return out
}

func findNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if n := findNode(root.Left, val); n != nil {
		return n
	}
	return findNode(root.Right, val)
}

func inorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	out := inorder(root.Left)
	out = append(out, root.Val)
	out = append(out, inorder(root.Right)...)
	return out
}

func TestSearchBST(t *testing.T) {
	t.Skip("TODO: implement searchBST")
	tests := []struct {
		tree []any
		val  int
		want []int
	}{
		{[]any{4, 2, 7, 1, 3}, 2, []int{2, 1, 3}},
		{[]any{4, 2, 7, 1, 3}, 5, []int{}},
	}
	for _, tc := range tests {
		got := searchBST(buildTree(tc.tree), tc.val)
		var gotInorder []int
		if got != nil { gotInorder = inorder(got) }
		if !reflect.DeepEqual(gotInorder, tc.want) && !(len(gotInorder) == 0 && len(tc.want) == 0) {
			t.Errorf("searchBST got %v, want %v", gotInorder, tc.want)
		}
	}
}
