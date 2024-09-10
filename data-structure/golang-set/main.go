package main

import (
	"fmt"
	"github.com/deckarep/golang-set/v2"
)

func main() {
	// Create a string-based set of required classes.
	required := mapset.NewSet[string]()
	required.Add("cooking")
	required.Add("english")
	required.Add("math")
	required.Add("biology")

	// Create a string-based set of science classes.
	sciences := mapset.NewSet[string]()
	sciences.Add("biology")
	sciences.Add("chemistry")

	// Create a string-based set of electives.
	electives := mapset.NewSet[string]()
	electives.Add("welding")
	electives.Add("music")
	electives.Add("automotive")

	// Create a string-based set of bonus programming classes.
	bonus := mapset.NewSet[string]()
	bonus.Add("beginner go")
	bonus.Add("python for dummies")

	all := required.
		Union(sciences).
		Union(electives).
		Union(bonus)
	fmt.Println(all)

	slice := all.ToSlice()
	fmt.Println(slice)
	set := mapset.NewSet(slice...)
	fmt.Println(set)
}
