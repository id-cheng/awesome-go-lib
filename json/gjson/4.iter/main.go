package main

import "github.com/tidwall/gjson"

func main() {
	json := `
		{"name": "Gilbert", "age": 60}
		{"name": "Alexa", "age": 35}
		{"name": "May", "age": 56}
		{"name": "cheng", "age": 18}
   `
	json2 := `
	{
	  "programmers": [
		{
		  "firstName": "Tom", 
		  "lastName": "Mc", 
		}, {
		  "firstName": "Ming", 
		  "lastName": "Hunter", 
		}, {
		  "firstName": "Jason", 
		  "lastName": "Harold", 
		}
	  ]
	}
`
	gjson.ForEachLine(json, func(line gjson.Result) bool {
		println(line.String())
		return true
	})

	result := gjson.Get(json2, "programmers.#.lastName")
	for i, name := range result.Array() {
		println(i)
		println(name.String())
	}

	result = gjson.Get(json2, "programmers")
	result.ForEach(func(key, value gjson.Result) bool {
		println(value.String())
		return true
	})
}
