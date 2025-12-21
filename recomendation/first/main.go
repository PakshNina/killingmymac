package main

type Node struct {
	value int
	next  *Node
}

func badLinkedList(n int) *Node {
	var head Node // ← stack object #1
	current := &head

	for i := 1; i < n; i++ {
		next := Node{value: i} // ← stack object #i!
		current.next = &next   // ← putPtr(&next)!
		current = &next
	}

	return &head
}

func goodArray(n int) []Node {
	nodes := make([]Node, n) // ← slice header = 1 stack object!

	for i := range nodes {
		nodes[i].value = i
	}

	return nodes
}
