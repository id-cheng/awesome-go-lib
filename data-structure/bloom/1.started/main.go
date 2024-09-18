package main

import (
	"encoding/binary"
	"fmt"
	"github.com/bits-and-blooms/bloom/v3"
)

// You may construct the Bloom filter capable of receiving 1 million elements
// with a false-positive rate of 1% in the following manner.
func main() {
	filter := bloom.NewWithEstimates(1000000, 0.01)
	filter.Add([]byte("Love"))
	if filter.Test([]byte("Love")) {
		fmt.Println("love")
	}

	n1 := make([]byte, 4)
	binary.BigEndian.PutUint32(n1, uint32(100))
	fmt.Println(string(n1))
	filter.Add(n1)
}
