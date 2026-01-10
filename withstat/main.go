package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

type Node struct {
	next    *Node
	payload [64]byte
}

func main() {
	signal := make(chan os.Signal, 1)

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	lastGC := m.NumGC

	curr := &Node{}
	var n int
	for {
		if n >= 1000000 {
			runtime.ReadMemStats(&m)
			if m.NumGC > lastGC {
				fmt.Printf("Количество сборок мусора: %d\n", m.NumGC)
				lastGC = m.NumGC
			}
			fmt.Printf("Куча: %.1fMB, стэк:%.1fMB\n",
				float64(m.HeapAlloc)/1024/1024,
				float64(m.StackInuse)/1024/1024)
			n = 0
			time.Sleep(500 * time.Millisecond)
			if float64(m.HeapAlloc)/1024/1024 > 2000 {
				fmt.Printf("Останавливаемся\n")
				break
			}
		}
		curr.next = new(Node)
		curr = curr.next
		n++
	}
	<-signal
}
