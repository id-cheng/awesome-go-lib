package main

import (
	"fmt"
	avl "github.com/emirpasic/gods/trees/avltree"
	"github.com/emirpasic/gods/trees/binaryheap"
	"github.com/emirpasic/gods/trees/btree"
	rbt "github.com/emirpasic/gods/trees/redblacktree"
	"github.com/emirpasic/gods/utils"
)

func main() {
	newRbt()
	newAvl()
	newBtree()
	newBinaryHeap()
}

func newBtree() {
	tree := btree.NewWithIntComparator(3) // empty (keys are of type int)

	tree.Put(1, "x") // 1->x
	tree.Put(2, "b") // 1->x, 2->b (in order)
	tree.Put(1, "a") // 1->a, 2->b (in order, replacement)
	tree.Put(3, "c") // 1->a, 2->b, 3->c (in order)
	tree.Put(4, "d") // 1->a, 2->b, 3->c, 4->d (in order)
	tree.Put(5, "e") // 1->a, 2->b, 3->c, 4->d, 5->e (in order)
	tree.Put(6, "f") // 1->a, 2->b, 3->c, 4->d, 5->e, 6->f (in order)
	tree.Put(7, "g") // 1->a, 2->b, 3->c, 4->d, 5->e, 6->f, 7->g (in order)

	fmt.Println(tree)
	// BTree
	//         1
	//     2
	//         3
	// 4
	//         5
	//     6
	//         7

	_ = tree.Values() // []interface {}{"a", "b", "c", "d", "e", "f", "g"} (in order)
	_ = tree.Keys()   // []interface {}{1, 2, 3, 4, 5, 6, 7} (in order)

	tree.Remove(2) // 1->a, 3->c, 4->d, 5->e, 6->f, 7->g (in order)
	fmt.Println(tree)
	// BTree
	//     1
	//     3
	// 4
	//     5
	// 6
	//     7

	tree.Clear() // empty
	tree.Empty() // true
	tree.Size()  // 0

	// Other:
	tree.Height()     // gets the height of the tree
	tree.Left()       // gets the left-most (min) node
	tree.LeftKey()    // get the left-most (min) node's key
	tree.LeftValue()  // get the left-most (min) node's value
	tree.Right()      // get the right-most (max) node
	tree.RightKey()   // get the right-most (max) node's key
	tree.RightValue() // get the right-most (max) node's value
}

func newAvl() {
	tree := avl.NewWithIntComparator() // empty(keys are of type int)

	tree.Put(1, "x") // 1->x
	tree.Put(2, "b") // 1->x, 2->b (in order)
	tree.Put(1, "a") // 1->a, 2->b (in order, replacement)
	tree.Put(3, "c") // 1->a, 2->b, 3->c (in order)
	tree.Put(4, "d") // 1->a, 2->b, 3->c, 4->d (in order)
	tree.Put(5, "e") // 1->a, 2->b, 3->c, 4->d, 5->e (in order)
	tree.Put(6, "f") // 1->a, 2->b, 3->c, 4->d, 5->e, 6->f (in order)

	fmt.Println(tree)
	//
	//  AVLTree
	//  │       ┌── 6
	//  │   ┌── 5
	//  └── 4
	//      │   ┌── 3
	//      └── 2
	//          └── 1

	_ = tree.Values() // []interface {}{"a", "b", "c", "d", "e", "f"} (in order)
	_ = tree.Keys()   // []interface {}{1, 2, 3, 4, 5, 6} (in order)

	tree.Remove(2) // 1->a, 3->c, 4->d, 5->e, 6->f (in order)
	fmt.Println(tree)
	//
	//  AVLTree
	//  │       ┌── 6
	//  │   ┌── 5
	//  └── 4
	//      └── 3
	//          └── 1

	tree.Clear() // empty
	tree.Empty() // true
	tree.Size()  // 0
}

func newRbt() {
	tree := rbt.NewWithIntComparator() // empty (keys are of type int)

	tree.Put(1, "x") // 1->x
	tree.Put(2, "b") // 1->x, 2->b (in order)
	tree.Put(1, "a") // 1->a, 2->b (in order, replacement)
	tree.Put(3, "c") // 1->a, 2->b, 3->c (in order)
	tree.Put(4, "d") // 1->a, 2->b, 3->c, 4->d (in order)
	tree.Put(5, "e") // 1->a, 2->b, 3->c, 4->d, 5->e (in order)
	tree.Put(6, "f") // 1->a, 2->b, 3->c, 4->d, 5->e, 6->f (in order)

	fmt.Println(tree)
	//
	//  RedBlackTree
	//  │           ┌── 6
	//	│       ┌── 5
	//	│   ┌── 4
	//	│   │   └── 3
	//	└── 2
	//		└── 1

	_ = tree.Values() // []interface {}{"a", "b", "c", "d", "e", "f"} (in order)
	_ = tree.Keys()   // []interface {}{1, 2, 3, 4, 5, 6} (in order)

	tree.Remove(2) // 1->a, 3->c, 4->d, 5->e, 6->f (in order)
	fmt.Println(tree)
	//
	//  RedBlackTree
	//  │       ┌── 6
	//  │   ┌── 5
	//  └── 4
	//      │   ┌── 3
	//      └── 1

	tree.Clear() // empty
	tree.Empty() // true
	tree.Size()  // 0

	// Other:
	tree.Left()     // gets the left-most (min) node
	tree.Right()    // get the right-most (max) node
	tree.Floor(1)   // get the floor node
	tree.Ceiling(1) // get the ceiling node
}

func newBinaryHeap() {
	// Min-heap
	heap := binaryheap.NewWithIntComparator() // empty (min-heap)
	heap.Push(2)                              // 2
	heap.Push(3)                              // 2, 3
	heap.Push(1)                              // 1, 3, 2
	heap.Values()                             // 1, 3, 2
	_, _ = heap.Peek()                        // 1,true
	_, _ = heap.Pop()                         // 1, true
	_, _ = heap.Pop()                         // 2, true
	_, _ = heap.Pop()                         // 3, true
	_, _ = heap.Pop()                         // nil, false (nothing to pop)
	heap.Push(1)                              // 1
	heap.Clear()                              // empty
	heap.Empty()                              // true
	heap.Size()                               // 0

	// Max-heap
	inverseIntComparator := func(a, b interface{}) int {
		return -utils.IntComparator(a, b)
	}
	heap = binaryheap.NewWith(inverseIntComparator) // empty (min-heap)
	heap.Push(2, 3, 1)                              // 3, 2, 1 (bulk optimized)
	heap.Values()                                   // 3, 2, 1
}
