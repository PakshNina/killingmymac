package recomendation

// Node объект, хранящий указатель на следующий узел и некую нагрузку.
type Node struct {
	next    *Node
	payload [64]byte
}

// badManyStackObjects создает много объектов стека.
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

// goodOneStackObject: создает один объект (слайс) на стеке.
func goodOneStackObject(n int) interface{} {
	nodes := make([]Node, n) // хедер слайса - первый и единственный объект стека

	for i := range nodes {
		if i > 0 {
			nodes[i].next = &nodes[i-1]
		}
	}

	return nodes
}
