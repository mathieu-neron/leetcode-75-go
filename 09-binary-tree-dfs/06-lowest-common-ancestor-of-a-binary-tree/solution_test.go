package solution

import "testing"

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

func TestLowestCommonAncestor(t *testing.T) {
	t.Skip("TODO: implement lowestCommonAncestor")
	tests := []struct {
		tree []any
		p, q int
		want int
	}{
		{[]any{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}, 5, 1, 3},
		{[]any{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}, 5, 4, 5},
		{[]any{1, 2}, 1, 2, 1},
	}
	for _, tc := range tests {
		root := buildTree(tc.tree)
		p := findNode(root, tc.p)
		q := findNode(root, tc.q)
		got := lowestCommonAncestor(root, p, q)
		if got == nil || got.Val != tc.want {
			var gv any = nil
			if got != nil { gv = got.Val }
			t.Errorf("lowestCommonAncestor got %v, want %v", gv, tc.want)
		}
	}
}
