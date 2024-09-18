package main

import (
	"bytes"
	"fmt"
	"github.com/bits-and-blooms/bitset"
)

func main() {
	const length = 9585
	const oneEvery = 97
	bs := bitset.New(length)
	// Add some bits
	for i := uint(0); i < length; i += oneEvery {
		bs = bs.Set(i)
	}

	var buf bytes.Buffer
	_, err := bs.WriteTo(&buf)
	if err != nil {
		// failure
	}
	// Here n == buf.Len()
	// Read back from buf
	bs = bitset.New(length)
	_, err = bs.ReadFrom(&buf)
	if err != nil {
		// error
	}
	fmt.Println(bs.String())
}
