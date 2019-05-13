package kit

import (
	"errors"
	"sync"
)

/*definition of tree node*/
type AVLNode struct {
	key         string
	val         interface{}
	left, right *AVLNode
	height      int
}

/*tree root node with read-write lock*/
type AVLTree struct {
	lock sync.RWMutex
	root *AVLNode
}

const (
	NODE_NOT_FOUND = "key not found"
)

/*find node with key k in tree under root n*/
func (n *AVLNode) value(k string) (interface{}, error) {
	for n != nil {
		if k < n.key {
			n = n.left
		} else if k > n.key {
			n = n.right
		} else {
			return n.val, nil
		}
	}

	return nil, errors.New(NODE_NOT_FOUND)
}

/*get height of node*/
func (root *AVLNode) getHeight() int {
	if root == nil {
		return 0
	}
	return root.height
}

/*get balance factor of node*/
func (root *AVLNode) getBalance() int {
	if root == nil {
		return 0
	}
	return root.left.getHeight() - root.right.getHeight()
}

/*create new tree node*/
func newAVLNode(k string, v interface{}) (n *AVLNode) {
	n = &AVLNode{key: k, val: v, height: 1}
	return
}

/*
     y            x
    / \    (R)   / \
   x   T3  =>   T1  y
  / \      <=      / \
 T1	 T2	   (L) 	 (T2) T3
    do right/left roation and change root node,
	for right roation, T2 will be moved from x.Right to y.Left, and root node will change from y to x
	for left roation, T2 will be moved from y.Left to x.Right, root node
	  will be changed from x to y
*/
func rightRotate(y *AVLNode) *AVLNode {
	if y == nil || y.left == nil {
		return y
	}
	x := y.left
	T2 := x.right

	x.right = y
	y.left = T2

	y.height = Max(y.left.getHeight(), y.right.getHeight()) + 1
	x.height = Max(x.left.getHeight(), x.right.getHeight()) + 1

	return x
}
func leftRotate(x *AVLNode) *AVLNode {
	if x == nil || x.right == nil {
		return x
	}
	y := x.right
	T2 := y.left

	y.left = x
	x.right = T2

	x.height = Max(x.left.getHeight(), x.right.getHeight()) + 1
	y.height = Max(y.left.getHeight(), y.right.getHeight()) + 1

	return y
}

/*rebalance the root node*/
func (root *AVLNode) reBalance(key string, balance int) *AVLNode {
	if root.left != nil && balance > 1 {
		if key < root.left.key { //left left case
			/*
			    root 'z' is the first unbalanced parent after insert
			   		*z					  (y)
			   	   /  \					/    \
			         *y   T4			   x      (z)
			        / \				  / \     / \
			   	x  *T3        =>    T1   T2 (T3) T4
			      / \
			     T1 T2
			     during right rotation, T2 will change its parent node
			*/
			return rightRotate(root)
		} else if key > root.left.key { //left right case
			/*
			   	   z					 *z				   (x)
			   	   / \					/  \			   / \
			        *y   T3			 *(x)   T3			  y	  (z)
			        / \				  /					 / \   \
			   	T1 *x        =>		(y)			=>		T1 T2	T3
			   	   /				/ \
			         *T2               T1 (T2)
			   	firstly, left rotate the the left node (y) of root (z)
			   	secondly, right rotate the root (z)
			*/
			root.left = leftRotate(root.left)
			return rightRotate(root)
		}
	} else if root.right != nil && balance < -1 {
		if key > root.right.key { //right right case
			/*
			       *z				 (y)
			      / \				/    \
			     T1  *y		  (z)     x
			        / \    =>	  / \    / \
			       *T2  x		 T1(T2) T3 T4
			   	   / \
			         T3 T4
			*/
			return leftRotate(root)
		} else if key < root.right.key { //right left case
			/*
			   	z					*z						(x)
			      / \				   / \					   /   \
			     T1 *y				  T1  *(x)				 (z)    (y)
			        / \     =>			 /  \		=>		 / \   / \
			       *x   T4			   *T2  (y)				T1 T2 T3 T4
			      / \					   /   \
			     T2  *T3			      (T3)  T4
			*/
			root.right = rightRotate(root.right)
			return leftRotate(root)
		}
	}

	return root
}

/*insert key-value pair into tree*/
func (root *AVLNode) insertKv(k string, v interface{}) *AVLNode {
	if root == nil {
		return newAVLNode(k, v)
	}

	if k < root.key {
		root.left = root.left.insertKv(k, v)
	} else if k > root.key {
		root.right = root.right.insertKv(k, v)
	}

	//update height of root node
	root.height = Max(root.left.getHeight(), root.right.getHeight()) + 1
	//balance the root and return new root node
	balance := root.getBalance()
	return root.reBalance(root.key, balance)
}

/*iterative in-order traversal*/
func (root *AVLNode) inorder() (out []Item) {
	if root == nil {
		return
	}

	var stack []*AVLNode
	curr := root

	for true {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.left
		}

		if len(stack) > 0 {
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			out = append(out, Item{curr.key, curr.val})
		}

		curr = curr.right
		if curr == nil && len(stack) == 0 {
			break
		}
	}
	return
}

/*iterative pre-order traversal*/
func (root *AVLNode) preorder() (out []Item) {
	if root == nil {
		return
	}

	queue := []*AVLNode{root}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		out = append(out, Item{curr.key, curr.val})

		if curr.left != nil {
			queue = append(queue, curr.left)
		}
		if curr.right != nil {
			queue = append(queue, curr.right)
		}
	}

	return
}

/*iterative post-order traversal with two stacks*/
func (root *AVLNode) postorder() (out []Item) {
	if root == nil {
		return
	}

	var stack1, stack2 []*AVLNode
	stack1 = append(stack1, root)
	for len(stack1) > 0 {
		curr := stack1[len(stack1)-1]
		stack1 = stack1[:len(stack1)-1]
		stack2 = append(stack2, curr)

		if curr.left != nil {
			stack1 = append(stack1, curr.left)
		}
		if curr.right != nil {
			stack1 = append(stack1, curr.right)
		}
	}

	for len(stack2) > 0 {
		curr := stack2[len(stack2)-1]
		stack2 = stack2[:len(stack2)-1]
		out = append(out, Item{curr.key, curr.val})
	}

	return
}

/*iterative level-order traversal with BFS*/
func (root *AVLNode) levelorder() (out [][]Item) {
	if root == nil {
		return
	}

	queue := []*AVLNode{root}
	level := 0
	for len(queue) > 0 {
		size := len(queue)
		if size > 0 {
			out = append(out, []Item{})
		} //allocate array to store next level
		for ; size > 0; size-- {
			curr := queue[0]
			queue = queue[1:]
			//travers node of current level
			out[level] = append(out[level], Item{curr.key, curr.val})

			//put next level nodes into queue
			if curr.left != nil {
				queue = append(queue, curr.left)
			}
			if curr.right != nil {
				queue = append(queue, curr.right)
			}
		}
		level++ //increase level
	}

	return
}

/*check links to node*/
func (root *AVLNode) hasLeft() bool {
	return root.left != nil
}
func (root *AVLNode) hasRight() bool {
	return root.right != nil
}
func (root *AVLNode) isLeaf() bool {
	return !root.hasLeft() && !root.hasRight()
}

/*get the left-most node which preserve the min key*/
func (root *AVLNode) min() *AVLNode {
	for ; root.hasLeft(); root = root.left {
	}
	return root
}

/*delete a node from AVL tree by key*/
func (root *AVLNode) delete(k string) (*AVLNode, error) {
	var err error
	if root == nil {
		return nil, errors.New(NODE_NOT_FOUND)
	}
	//recursively find the parent node of to-delete node
	if k < root.key {
		root.left, err = root.left.delete(k)
		return root, err
	} else if k > root.key {
		root.right, err = root.right.delete(k)
		return root, err
	}

	if root.isLeaf() {
		return nil, nil //if the to-delete node is leaf, return it as nil to remove link
	} else if root.hasLeft() && !root.hasRight() {
		return root.left, nil
	} else if !root.hasLeft() && root.hasRight() {
		return root.right, nil
	} //if the to-delete node has only one child, replace the position of root
	//with it child to remove the link

	//exchange value of to-delete node with the left most node of its right
	// node, which is the successor of to-delete node
	min := root.right.min()
	root.key, root.val = min.key, min.val
	root.right, err = root.right.delete(min.key)
	return root, err
}

/*inorder iterator*/
//refer to https://danrl.com/blog/2018/basic-data-structures-binary-tree/
type Item struct {
	Key string
	Val interface{}
}

/*inorder traversal through channel*/
func (root *AVLNode) iter(ch chan<- Item) {
	if root == nil {
		return
	}
	root.left.iter(ch)
	ch <- Item{Key: root.key, Val: root.val}
	root.right.iter(ch)
}

//public function guarante concurrency-safty
func NewAVLTree() (t *AVLTree) {
	t = &AVLTree{}
	return t
}

/*search value by key in tree*/
func (t *AVLTree) Value(k string) (interface{}, error) {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return t.root.value(k)
}

/*insert key-value pair to tree with tree gurannted to be balanced*/
func (t *AVLTree) BalancedInsert(key string, val interface{}) {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.root = t.root.insertKv(key, val)
}

/*get the height of tree*/
func (t *AVLTree) GetHeight() int {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return t.root.getHeight()
}

/*check if tree is balanced*/
func (t *AVLTree) IsBalanced() bool {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return Abs(t.root.getBalance()) <= 1
}

/*iterative in-order traversal*/
func (t *AVLTree) Inorder() (out []Item) {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return t.root.inorder()
}

/*iterative pre-order traversal*/
func (t *AVLTree) Preorder() (out []Item) {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return t.root.preorder()
}

/*iterative post-order traversal*/
func (t *AVLTree) Postorder() (out []Item) {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return t.root.postorder()
}

/*iterative level-order traversal with BFS*/
func (t *AVLTree) LevelOrder() (out [][]Item) {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return t.root.levelorder()
}

/*inorder traversal through channel*/
func (t *AVLTree) Iter() <-chan Item {
	ch := make(chan Item)
	t.lock.RLock()
	//delegate all tasks to a goroutine,
	//  keep program from blocking
	go func() {
		t.root.iter(ch)
		t.lock.RUnlock() //unlock the mutex once
		close(ch)
	}()
	return ch
}

/*deletion*/
func (t *AVLTree) Delete(k string) error {
	var err error
	t.lock.Lock()
	defer t.lock.Unlock()
	t.root, err = t.root.delete(k)
	return err
}
