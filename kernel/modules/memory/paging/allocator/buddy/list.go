package buddy

type list struct {
	head *node
}

func (l *list) Head() *node {
	if l == nil {
		return nil
	}
	return l.head
}

func (l *list) SetHead(head *node) {
	l.head = head
}

func (l *list) IsEmpty() bool {
	return l == nil || l.head == nil
}

func (l *list) find(value uint32) (curr *node, prev *node, ok bool) {
	curr = l.Head()

	for curr != nil {
		if curr.Value() == value {
			break
		}
		prev = curr
		curr = curr.Next()
	}

	if curr != nil {
		ok = true
	}
	return
}

func (l *list) Delete(value uint32) {
	if l == nil {
		return
	}

	curr := l.Head()
	var prev *node

	curr, prev, ok := l.find(value)

	// Already deleted, rejoice
	if !ok {
		return
	}

	if curr == l.Head() {
		// remove head
		l.SetHead(nil)
	} else {
		// point prev to next of curr
		prev.SetNext(curr.Next())
	}

	// free curr
	releaseNode(curr)
}

func (l *list) Append(value uint32) (ok bool) {
	node, ok := newNode()
	if !ok {
		// OMG, we cannot allocate a node for list, something is wrong with node
		// pool initialization
		return false
	}
	node.SetValue(value)
	node.SetNext(l.Head())
	l.SetHead(node)
	return true
}
