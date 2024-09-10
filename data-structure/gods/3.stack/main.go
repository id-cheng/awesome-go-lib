package main

import (
	"github.com/emirpasic/gods/stacks/arraystack"
	"github.com/emirpasic/gods/stacks/linkedliststack"
)

func main() {
	newLinkedListStack()
	newArrayStack()
}

func newArrayStack() {
	stack := arraystack.New()
	stack.Push(1)
	stack.Push(2)
	stack.Values()
	_, _ = stack.Peek()
	_, _ = stack.Pop()
	_, _ = stack.Pop()
	_, _ = stack.Pop()
	stack.Push(1)
	stack.Clear()
	stack.Empty()
	stack.Size()
}

func newLinkedListStack() {
	stack := linkedliststack.New()
	stack.Push(1)
	stack.Push(2)
	stack.Values()
	_, _ = stack.Peek()
	_, _ = stack.Pop()
	_, _ = stack.Pop()
	_, _ = stack.Pop()
	stack.Push(1)
	stack.Clear()
	stack.Empty()
	stack.Size()
}
