package main

import (
	"fmt"
	"github.com/bits-and-blooms/bitset"
	"math/rand"
)

func main() {
	fmt.Printf("Hello from BitSet!\n")
	var b bitset.BitSet
	// play some Go Fish
	for i := 0; i < 100; i++ {
		card1 := uint(rand.Intn(52))
		card2 := uint(rand.Intn(52))
		b.Set(card1)
		if b.Test(card2) {
			fmt.Println("Go Fish!")
		}
		b.Clear(card1)
	}

	// Chaining
	b.Set(10).Set(11)
	fmt.Println(b.String())
	for i, e := b.NextSet(0); e; i, e = b.NextSet(i + 1) {
		fmt.Println("The following bit is set:", i)
	}
	fmt.Println()
	if b.Intersection(bitset.New(100).Set(10)).Count() == 1 {
		fmt.Println("Intersection works.")
	} else {
		fmt.Println("Intersection doesn't work???")
	}
}
