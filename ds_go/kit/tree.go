package kit

import "fmt"

type TreeNode struct {
	Left, Right *TreeNode
	Key         int
}

type Tree struct {
	Root *TreeNode
	Num  int
}

func NewTree() (t *Tree) {
	t = &Tree{}
	return
}

/*iterative insert new key into tree*/
func (t *Tree) Insert(val int) {
	node := &TreeNode{Key: val}
	if t.Root == nil {
		t.Root = node
		t.Num++
		return
	} //create root

	//find the parent node to host new one
	var cur, parent *TreeNode = t.Root, nil
	for cur != nil {
		parent = cur
		if val < cur.Key {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}

	//link the new node to the tree
	if parent == nil {
		panic("corrupt tree") //parent = node
	} else if val < parent.Key {
		parent.Left = node
	} else if val > parent.Key {
		parent.Right = node
	}
	t.Num++
}

/*recursively insert new key into tree*/
func (t *Tree) InsertRecur(val int) {
	t.Root = insert(t.Root, val)
}

func insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Key: val}
	}

	if val < root.Key {
		root.Left = insert(root.Left, val)
	} else if val > root.Key {
		root.Right = insert(root.Right, val)
	}
	return root
}

/*iterative in-order traversal*/
func (t *Tree) Inorder() (out []int) {
	if t.Root == nil {
		return
	}

	var stack []*TreeNode
	curr := t.Root

	for true {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		if len(stack) > 0 {
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			out = append(out, curr.Key)
		}

		curr = curr.Right
		if curr == nil && len(stack) == 0 {
			break
		}
	}
	return
}

/*iterative pre-order traversal*/
func (t *Tree) Preorder() (out []int) {
	if t.Root == nil {
		return
	}

	queue := []*TreeNode{t.Root}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		out = append(out, curr.Key)

		if curr.Left != nil {
			queue = append(queue, curr.Left)
		}
		if curr.Right != nil {
			queue = append(queue, curr.Right)
		}
	}

	return
}

/*iterative post-order traversal with two stacks*/
func (t *Tree) Postorder() (out []int) {
	if t.Root == nil {
		return
	}

	var stack1, stack2 []*TreeNode
	stack1 = append(stack1, t.Root)
	for len(stack1) > 0 {
		curr := stack1[len(stack1)-1]
		stack1 = stack1[:len(stack1)-1]
		stack2 = append(stack2, curr)

		if curr.Left != nil {
			stack1 = append(stack1, curr.Left)
		}
		if curr.Right != nil {
			stack1 = append(stack1, curr.Right)
		}
	}

	for len(stack2) > 0 {
		curr := stack2[len(stack2)-1]
		stack2 = stack2[:len(stack2)-1]
		out = append(out, curr.Key)
	}

	return
}

/*iterative level-order traversal with BFS*/
func (t *Tree) LevelOrder() (out [][]int) {
	if t.Root == nil {
		return
	}

	queue := []*TreeNode{t.Root}
	level := 0
	for len(queue) > 0 {
		size := len(queue)
		if size > 0 {
			out = append(out, []int{})
		} //allocate array to store next level
		for ; size > 0; size-- {
			curr := queue[0]
			queue = queue[1:]
			out[level] = append(out[level], curr.Key) //travers node of current level

			//put next level nodes into queue
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
		level++ //increase level
	}

	return
}

/*get the height of tree in recursively manner*/
func getHeightRecur(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	return Max(getHeightRecur(root.Left), getHeightRecur(root.Right)) + 1
}

/*check if tree is balanced in recursively manner*/
func (t *Tree) IsBalancedRecur() bool {
	if t.Root == nil {
		return true
	}
	return Abs(getHeightRecur(t.Root.Left)-getHeightRecur(t.Root.Right)) < 1
}

/*create and link tree node for a balanced tree*/
func buildBalancedTree(keys []int, start, end int) (root *TreeNode) {
	if start < end {
		return
	}

	mid := start + (end-start)/2
	root = &TreeNode{Key: keys[mid]}
	root.Left = buildBalancedTree(keys, start, mid-1)
	root.Right = buildBalancedTree(keys, mid+1, end)
	return
}

/*rebuild the tree to make it balanced*/
func (t *Tree) MakeBalancedRecur() {
	keys := t.Inorder()
	t.Root = buildBalancedTree(keys, 0, len(keys)-1)
}

func TreeTest(n int) {
	nums := []int{3, 4, 6, 2, 1, 5}
	tree := NewTree()
	for _, n := range nums {
		tree.Insert(n)
	}
	//visit := tree.Inorder()
	fmt.Println("#", n, "TreeTest", tree.Num)
	fmt.Println("pre-order", tree.Preorder())
	fmt.Println("in-order", tree.Inorder())
	fmt.Println("post-order", tree.Postorder())
	fmt.Println("level-order", tree.LevelOrder())
}
