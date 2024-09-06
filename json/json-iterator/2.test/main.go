package main

import (
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"testing"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var jsonData = []byte(`{"name":"cheng","age":30}`)

func BenchmarkEncodingJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var p Person
		err := json.Unmarshal(jsonData, &p)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkJSONIterator(b *testing.B) {
	var jsonIter = jsoniter.ConfigCompatibleWithStandardLibrary
	for i := 0; i < b.N; i++ {
		var p Person
		err := jsonIter.Unmarshal(jsonData, &p)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func main() {
	fmt.Println("Running benchmarks...")
	// 性能测试
	result1 := testing.Benchmark(BenchmarkEncodingJSON)
	result2 := testing.Benchmark(BenchmarkJSONIterator)

	fmt.Printf("encoding/json: %s\n", result1)
	fmt.Printf("json-iterator: %s\n", result2)
}
