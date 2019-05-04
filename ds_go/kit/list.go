package kit

import (
	"fmt"
	"math"
)

type Val_t int

type ListNode struct {
	Val  Val_t
	Next *ListNode
}

type List struct {
	Head *ListNode
	Tail *ListNode
	Len  int
}

func NewList() (lst *List) {
	lst = &List{} //escpae once return
	return
}

func (lst *List) Ints2Lst(nums []Val_t) {
	if len(nums) <= 0 {
		return
	}

	//create a dummy node
	dummy := &ListNode{Val: math.MaxInt32}
	prevNode := dummy
	lst.Len = 0
	//create new node and link as tail
	for _, n := range nums {
		newNode := &ListNode{Val: n}
		prevNode.Next = newNode
		lst.Len++
		prevNode = newNode
	}

	//update head and tail
	if lst.Len > 0 {
		lst.Head = dummy.Next
		lst.Tail = prevNode
	}
}

func (lst *List) Print() {
	curr := lst.Head
	if lst.Len > 0 {
		fmt.Println("[len]", lst.Len, "[head]", lst.Head.Val,
			"[tail]", lst.Tail.Val)
	} else {
		fmt.Println("len", lst.Len)
		return
	}

	for curr != nil {
		fmt.Printf("%v", curr.Val)
		if curr != lst.Tail {
			fmt.Printf("->")
		} else {
			fmt.Println()
		}
		curr = curr.Next
	}
}

/*ret node at index idx*/
func (lst *List) GetAt(idx int) *ListNode {
	if idx < 0 || idx > lst.Len {
		panic("index out of boundary")
	}
	curr := lst.Head
	for idx > 0 {
		curr = curr.Next
		idx--
	}
	return curr
}

/*append node as tail*/
func (lst *List) PushBack(item Val_t) {
	newNode := &ListNode{Val: item}
	if lst.Len <= 0 {
		lst.Head, lst.Tail = newNode, newNode
	} else {
		lst.Tail.Next = newNode
		lst.Tail = newNode
	}
	lst.Len++
}

/*preppend node as head*/
func (lst *List) PushFront(item Val_t) {
	newNode := &ListNode{Val: item}
	if lst.Len <= 0 {
		lst.Head, lst.Tail = newNode, newNode
	} else {
		newNode.Next = lst.Head
		lst.Head = newNode
	}
	lst.Len++
}

/*insert node at index idx*/
func (lst *List) InsertAt(idx int, item Val_t) {
	if idx < 0 || idx > lst.Len {
		panic("index out of boundary")
	}

	newNode := &ListNode{Val: item}

	//insert a node after the idx-th node
	dummy := &ListNode{Val: math.MaxInt32}
	dummy.Next = lst.Head
	prev, next := dummy, dummy.Next
	for idx > 0 {
		prev = next
		next = next.Next
		idx--
	}
	prev.Next = newNode
	newNode.Next = next

	//update the reference of head and tail
	lst.Len++
	lst.Head = dummy.Next
	curr := lst.Head
	for curr.Next != nil {
		curr = curr.Next
	}
	lst.Tail = curr
}

/*remove the tail node*/
func (lst *List) PopBack() (ret Val_t) {
	if lst.Len < 1 {
		panic("no item to pop")
	}
	ret = lst.Tail.Val
	if lst.Len == 1 {
		lst.Head, lst.Tail = nil, nil
	} else {
		prevTail := lst.Head
		//move to the tail
		for prevTail.Next != lst.Tail {
			prevTail = prevTail.Next
		}
		//unlink tail node, it will be re-collected
		lst.Tail = prevTail
	}
	lst.Len--
	return
}

/*remove the head node*/
func (lst *List) PopFront() (ret Val_t) {
	if lst.Len < 1 {
		panic("no item to pop")
	}
	ret = lst.Head.Val
	if lst.Len == 1 {
		lst.Head, lst.Tail = nil, nil
	} else {
		lst.Head = lst.Head.Next
	}
	lst.Len--
	return
}

/*remove the node at index idx*/
func (lst *List) RemoveAt(idx int) (ret Val_t) {
	if idx < 0 || idx > lst.Len {
		panic("index out of boundary")
	}
	curr := lst.GetAt(idx)
	ret = curr.Val
	if idx == 0 {
		return lst.PopFront()
	} else if idx == lst.Len-1 {
		return lst.PopBack()
	} else {
		prev, next := lst.GetAt(idx-1), lst.GetAt(idx+1)
		prev.Next = next
		lst.Len--
	}
	return
}

/*reverse list start from head*/
func reverse(head *ListNode) *ListNode {
	var prev, curr *ListNode = nil, head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

/*reverse the list between from and to th node*/
func (lst *List) Reverse(from, to int) {
	if from < 0 || from >= lst.Len || to < 0 ||
		to >= lst.Len || from > to {
		panic("index out of boundary")
	}

	if from == to {
		return
	}
	var start, end, before, next *ListNode
	curr := lst.Head

	//fint the four point before reverse
	for cnt := 0; cnt <= to; cnt++ {
		if cnt < from {
			before = curr
		}
		if cnt == from {
			start = curr
		}
		if cnt == to {
			end = curr
		}
		curr = curr.Next
	}
	next = end.Next
	end.Next = nil //break the from-to list

	//re-link from-to list head
	if before != nil {
		before.Next = reverse(start)
	} else {
		lst.Head = reverse(start)
	}
	//re-link from-to list to remain part
	start.Next = next
	lst.Tail = start

	//update list tail
	for lst.Tail.Next != nil {
		lst.Tail = lst.Tail.Next
	}
}

func LstTest() {
	nums := []Val_t{3, 5, 2, 4, 6}
	lst := NewList()
	lst.Ints2Lst(nums)
	lst.Print()
	lst.Reverse(0, lst.Len-1)
	lst.Print()
	/*lst.PushBack(7)
	lst.PushFront(1)
	lst.Print()
	lst.InsertAt(4, 3)
	lst.Print()
	fmt.Printf("3rd item: %v\n", lst.GetAt(3))

	lst.PopBack()
	lst.PopFront()
	fmt.Printf("remove 3rd item: %v \n", lst.RemoveAt(3))
	lst.Print()*/
}
