package main

// BigWithPtr структура с указателем.
type BigWithPtr struct {
	ID   int
	Next *BigWithPtr // указатель внутри
	Data [1024]byte  // что-то тяжёлое
}

// BigNoPtr структура без указателя.
type BigNoPtr struct {
	ID   int
	Data [1024]byte // тот же размер, но без указателя
}

func useBigWithPtr(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		x := BigWithPtr{ID: i}
		if x.Next != nil { // доступ к pointer‑полю
			sum += x.Next.ID
		} else {
			sum += x.ID
		}
	}
	return sum
}

func useBigNoPtr(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		x := BigNoPtr{ID: i}
		sum += x.ID
	}
	return sum
}
