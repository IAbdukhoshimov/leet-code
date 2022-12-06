package main

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var max func(left, right int) int
	max = func(left, right int) int {
		if left > right {
			return left
		}
		return right
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
