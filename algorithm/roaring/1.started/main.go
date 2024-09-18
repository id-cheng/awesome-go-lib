package main

import (
	"bytes"
	"fmt"
	"github.com/RoaringBitmap/roaring/v2"
)

func main() {
	// example inspired by https://github.com/fzandona/goroar
	fmt.Println("==roaring==")
	rb1 := roaring.BitmapOf(1, 2, 3, 4, 5, 100, 1000)
	fmt.Println(rb1.String())

	rb2 := roaring.BitmapOf(3, 4, 1000)
	fmt.Println(rb2.String())

	rb3 := roaring.New()
	fmt.Println(rb3.String())

	fmt.Println("Cardinality: ", rb1.GetCardinality())

	fmt.Println("Contains 3? ", rb1.Contains(3))

	rb1.And(rb2)
	fmt.Println("rb1:", rb1.String())
	rb3.Add(1)
	rb3.Add(5)

	rb3.Or(rb1)
	fmt.Println("rb3:", rb3.String())
	// prints 1, 3, 4, 5, 1000
	i := rb3.Iterator()
	for i.HasNext() {
		fmt.Println(i.Next())
	}
	fmt.Println()

	// computes union of the three bitmaps in parallel using 4 workers
	fmt.Println(roaring.ParOr(4, rb1, rb2, rb3))
	// computes intersection of the three bitmaps in parallel using 4 workers
	fmt.Println(roaring.ParAnd(4, rb1, rb2, rb3))

	// next we include an example of serialization
	buf := new(bytes.Buffer)
	rb1.WriteTo(buf) // we omit error handling
	newRb := roaring.New()
	newRb.ReadFrom(buf)
	if newRb.Validate() != nil {
		fmt.Println("Failed validation") // return or panic
	}
	if rb1.Equals(newRb) {
		fmt.Println("I wrote the content to a byte stream and read it back.")
	}
	fmt.Println("new:", newRb.String())
	// you can iterate over bitmaps using ReverseIterator(), Iterator, ManyIterator()
}
