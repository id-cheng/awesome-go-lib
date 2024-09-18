package main

import (
	"bytes"
	"fmt"
	"github.com/bits-and-blooms/bloom/v3"
)

func main() {
	f := bloom.New(1000, 4)
	var buf bytes.Buffer
	bytesWritten, _ := f.WriteTo(&buf)
	var g bloom.BloomFilter
	bytesRead, _ := g.ReadFrom(&buf)
	if bytesRead == bytesWritten {
		fmt.Println("ok")
	}
}
