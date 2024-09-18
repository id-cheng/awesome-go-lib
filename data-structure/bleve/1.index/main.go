package main

import "github.com/blevesearch/bleve/v2"

func main() {
	message := struct {
		Id   string
		From string
		Body string
	}{
		Id:   "example",
		From: "marty.schoch@gmail.com",
		Body: "bleve indexing is easy",
	}

	mapping := bleve.NewIndexMapping()
	index, err := bleve.New("example.bleve", mapping)
	if err != nil {
		panic(err)
	}
	_ = index.Index(message.Id, message)
}
