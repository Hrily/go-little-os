package buddy

// node in a linked list
type node struct {
	value uint32
	next  *node
}

// Value of current node
func (n *node) Value() uint32 {
	if n == nil {
		return 0
	}
	return n.value
}

// SetValue of current node
func (n *node) SetValue(value uint32) {
	n.value = value
}

// Next returns next node
func (n *node) Next() *node {
	return n.next
}

// SetNext node
func (n *node) SetNext(next *node) {
	n.next = next
}

// HasNext tells if current node has a next
func (n *node) HasNext() bool {
	return n.next != nil
}
