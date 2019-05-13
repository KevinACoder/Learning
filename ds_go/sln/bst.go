package sln

//230. Kth Smallest Element in a BST
func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return -1
	}

	var stack []*TreeNode
	curr := root

	for true {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		if len(stack) > 0 {
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if k <= 0 {
				return curr.key
			}
			k--
		}

		curr = curr.Right
		if curr == nil && len(stack) == 0 {
			break
		}
	}
	return -1
}
