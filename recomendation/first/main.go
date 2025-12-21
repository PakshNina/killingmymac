package main

type Node struct {
	value int
	next  *Node
}

func badLinkedList(n int) *Node {
	var head Node // первый объект стека
	current := &head

	for i := 1; i < n; i++ {
		next := Node{value: i} // создание i-го объекта стека
		current.next = &next
		current = &next
	}

	return &head
}

func goodArray(n int) []Node {
	nodes := make([]Node, n) // хедер слайса - первый и единственный объект стека.

	for i := range nodes {
		nodes[i].value = i
	}

	return nodes
}
