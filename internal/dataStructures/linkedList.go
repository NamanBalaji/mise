package dataStructures

type ListNode struct {
	prev  *ListNode
	next  *ListNode
	value interface{}
}

type LinkedList struct {
	head *ListNode
	tail *ListNode
	len  uint32
}

// Create a new doubly linked list
func NewLinkedList() *LinkedList {
	return &LinkedList{
		head: nil,
		tail: nil,
		len:  0,
	}
}

func (r *LinkedList) Len() uint32 {
	return r.len
}

func (r *ListNode) Value() interface{} {
	return r.value
}

// last node
func (r *LinkedList) GetLast() *ListNode {
	return r.tail
}

// first node
func (r *LinkedList) GetFirst() *ListNode {
	return r.head
}

// add to the header
func (r *LinkedList) AddHead(value interface{}) {
	node := &ListNode{value: value}

	if r.head == nil && r.len == 0 {
		r.head = node
		r.tail = node
		node.prev = nil
		node.next = nil
		r.len++
		return
	}

	node.next = r.head.next
	node.prev = nil
	r.head.next.prev = node
	r.head = node
	r.len++
}

// add to the end of the list
func (r *LinkedList) AddTail(value interface{}) {
	node := &ListNode{value: value}

	if r.tail == nil && r.len == 0 {
		r.head = node
		r.tail = node
		node.prev = nil
		node.next = nil
		r.len++
		return
	}

	r.tail.next = node
	node.prev = r.tail
	node.next = nil
	r.tail = node
	r.len++
}

// delete the first node
func (r *LinkedList) DelHead() *ListNode {
	if r.len == 0 {
		return nil
	}

	node := r.head

	r.head.next.prev = nil
	r.head = r.head.next
	r.len--

	return node
}

// delete the last node
func (r *LinkedList) DelTail() *ListNode {
	if r.len == 0 {
		return nil
	}

	node := r.tail

	r.tail.prev.next = nil
	r.tail = r.tail.prev
	r.tail.next = nil
	r.len--
	return node

}
