package main

import (
	"fmt"
	"github.com/blevesearch/bleve/v2"
)

func main() {
	index, _ := bleve.Open("example.bleve")
	query := bleve.NewQueryStringQuery("indexing")
	searchRequest := bleve.NewSearchRequest(query)
	searchResult, _ := index.Search(searchRequest)
	fmt.Println(searchResult)
}
