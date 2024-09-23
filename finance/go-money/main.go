package main

import (
	"fmt"
	"github.com/Rhymond/go-money"
)

func main() {
	pound := money.New(100, money.CNY)
	twoPounds, _ := pound.Add(pound)
	result, _ := pound.Add(twoPounds)
	greater, _ := twoPounds.GreaterThan(pound)
	fmt.Println("greater:", greater)
	fmt.Println("amount:", twoPounds.Amount())
	fmt.Println("isZero:", pound.IsZero())
	fmt.Println("isPositive:", pound.IsPositive())
	fmt.Println("isNegative:", pound.IsNegative())
	fmt.Println("result:", result.Amount())
	parties, _ := twoPounds.Split(3)
	fmt.Println(parties[0].Display(),
		parties[1].Display(),
		parties[2].Display())

	allocate, _ := twoPounds.Allocate(33, 33, 33)
	fmt.Println(allocate[0].Display())
}
