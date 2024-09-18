package main

import (
	"fmt"
	"github.com/bits-and-blooms/bloom/v3"
)

var n uint = 1000000
var fp = 0.001

// 具有 m 位和 k 个哈希函数的bloom过滤器，用于一组大小为 n 的集合
func main() {
	m, k := bloom.EstimateParameters(n, fp)
	ActualfpRate := bloom.EstimateFalsePositiveRate(m, k, n)
	fmt.Println(ActualfpRate)

	f := bloom.NewWithEstimates(n, fp)
	ActualfpRate = bloom.EstimateFalsePositiveRate(f.Cap(), f.K(), n)
	fmt.Println(ActualfpRate)
}
