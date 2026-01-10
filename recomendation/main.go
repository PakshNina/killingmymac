package main

// Node объект, хранящий указатель на следующий узел и некую нагрузку.
type Node struct {
	next    *Node
	payload [64]byte
}

func badManyStackObjects(n int) interface{} {
	var head Node
	current := &head

	for i := 1; i < n; i++ {
		next := Node{}
		current.next = &next
		current = &next
	}

	return &head
}

func goodOneStackObject(n int) interface{} {
	nodes := make([]Node, n)

	for i := range nodes {
		if i > 0 {
			nodes[i].next = &nodes[i-1]
		}
	}

	return nodes
}
