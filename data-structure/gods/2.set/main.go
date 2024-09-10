package main

import (
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/emirpasic/gods/sets/linkedhashset"
	"github.com/emirpasic/gods/sets/treeset"
)

func main() {
	newHashSet()
	newTreeSet()
	newLinkedHashSet()
}

func newLinkedHashSet() {
	set := linkedhashset.New()
	set.Add(5)
	set.Add(4, 4, 3, 2, 1)
	set.Add(4)
	set.Remove(4)
	set.Remove(2, 3)
	set.Contains(1)
	set.Contains(1, 5)
	set.Contains(1, 6)
	_ = set.Values()
	set.Clear()
	set.Empty()
	set.Size()
}

func newTreeSet() {
	set := treeset.NewWithIntComparator()
	set.Add(1)
	set.Add(2, 2, 3, 4, 5)
	set.Remove(4)
	set.Remove(2, 3)
	set.Contains(1)
	set.Contains(1, 5)
	set.Contains(1, 6)
	_ = set.Values()
	set.Clear()
	set.Empty()
	set.Size()
}

func newHashSet() {
	set := hashset.New()
	set.Add(1)
	set.Add(2, 2, 3, 4, 5)
	set.Remove(1)
	set.Remove(2, 3)
	set.Contains(1)
	set.Contains(1, 5)
	set.Contains(1, 6)
	_ = set.Values()
	set.Clear()
	set.Empty()
	set.Size()
}
