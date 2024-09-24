package main

import (
	"fmt"
	"github.com/ahmetb/go-linq/v3"
)

func main() {
	Range()
	Repeat()
	Query()
	Aggregate()
	Distinct()
	DistinctBy()
	ThenBy()
	Any()
	Average()
	Append()
	Count()
	Contains()
}

func Contains() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	has5 := linq.From(slice).Contains(5)

	fmt.Printf("Does the slice contains 5? %t", has5)
}

func Count() {
	fruits := []string{"apple", "banana", "mango", "orange", "grape", "pear"}
	numberOfFruits := linq.From(fruits).Count()

	fmt.Println(numberOfFruits)
}

func Append() {
	input := []int{1, 2, 3, 4}
	q := linq.From(input).Append(5)
	last := q.Last()
	fmt.Println(last)
}

func Average() {
	grades := []int{78, 92, 100, 37, 81}
	average := linq.From(grades).Average()

	fmt.Println(average)
}

func Any() {
	numbers := []int{1, 2}
	hasElements := linq.From(numbers).Any()
	fmt.Printf("Are there any element in the list? %t", hasElements)
}

func ThenBy() {
	fruits := []string{"grape", "banana", "mango", "orange", "raspberry", "apple", "blueberry"}

	// Sort the strings first by their length and then
	//alphabetically by passing the identity selector function.
	var query []string
	linq.From(fruits).
		OrderBy(
			func(fruit interface{}) interface{} { return len(fruit.(string)) },
		).
		ThenBy(
			func(fruit interface{}) interface{} { return fruit },
		).
		ToSlice(&query)

	for _, fruit := range query {
		fmt.Println(fruit)
	}
}

func DistinctBy() {
	type Product struct {
		Name string
		Code int
	}

	products := []Product{
		{Name: "orange", Code: 4},
		{Name: "apple", Code: 9},
		{Name: "lemon", Code: 12},
		{Name: "apple", Code: 9},
	}

	//Order and exclude duplicates.
	var list []Product
	linq.From(products).
		OrderBy(
			func(item interface{}) interface{} { return item.(Product).Name },
		).
		DistinctBy(
			func(item interface{}) interface{} { return item.(Product).Code },
		).
		ToSlice(&list)

	for _, product := range list {
		fmt.Printf("%s %d\n", product.Name, product.Code)
	}
}

func Distinct() {
	ages := []int{21, 46, 46, 55, 17, 21, 55, 55}

	var distinctAges []int
	linq.From(ages).
		OrderBy(
			func(item any) any { return item },
		).
		Distinct().
		ToSlice(&distinctAges)

	fmt.Println(distinctAges)
}

func Aggregate() {
	fruits := []string{"apple", "mango", "orange", "watermelon", "grape"}

	// Determine which string in the slice is the longest.
	longestName := linq.From(fruits).
		Aggregate(
			func(r any, i any) any {
				if len(r.(string)) > len(i.(string)) {
					return r
				}
				return i
			},
		)

	fmt.Println(longestName)
}

func Query() {
	query := linq.From([]int{1, 2, 3, 4, 5}).Where(func(i any) bool {
		return i.(int) <= 3
	})
	query.ForEach(func(i any) {
		fmt.Println(i.(int))
	})
}

func Repeat() {
	var slice []string
	linq.Repeat("I like programming.", 5).ToSlice(&slice)

	for _, str := range slice {
		fmt.Println(str)
	}
}

func Range() {
	// Generate a slice of integers from 1 to 10
	// and then select their squares.
	var squares []int
	linq.Range(1, 10).
		SelectT(
			func(x int) int { return x * x },
		).
		ToSlice(&squares)

	for _, num := range squares {
		fmt.Println(num)
	}
}
