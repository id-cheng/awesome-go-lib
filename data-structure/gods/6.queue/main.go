package main

import (
	"github.com/emirpasic/gods/queues/arrayqueue"
	"github.com/emirpasic/gods/queues/linkedlistqueue"
)

func main() {
	newArrayQueue()
	newLinkedListQueue()
}

func newArrayQueue() {
	queue := arrayqueue.New() // empty
	queue.Enqueue(1)          // 1
	queue.Enqueue(2)          // 1, 2
	_ = queue.Values()        // 1, 2 (FIFO order)
	_, _ = queue.Peek()       // 1,true
	_, _ = queue.Dequeue()    // 1, true
	_, _ = queue.Dequeue()    // 2, true
	_, _ = queue.Dequeue()    // nil, false (nothing to deque)
	queue.Enqueue(1)          // 1
	queue.Clear()             // empty
	queue.Empty()             // true
	_ = queue.Size()          // 0
}

func newLinkedListQueue() {
	queue := linkedlistqueue.New() // empty
	queue.Enqueue(1)               // 1
	queue.Enqueue(2)               // 1, 2
	_ = queue.Values()             // 1, 2 (FIFO order)
	_, _ = queue.Peek()            // 1,true
	_, _ = queue.Dequeue()         // 1, true
	_, _ = queue.Dequeue()         // 2, true
	_, _ = queue.Dequeue()         // nil, false (nothing to deque)
	queue.Enqueue(1)               // 1
	queue.Clear()                  // empty
	queue.Empty()                  // true
	_ = queue.Size()               // 0
}
