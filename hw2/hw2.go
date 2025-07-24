package main

// Homework 2: Object Oriented Programming
// Due February 7, 2017 at 11:59pm

import "fmt"

func main() {
	// Feel free to use the main function for testing your functions
	// world := struct {
	// 	English string
	// 	Spanish string
	// 	French  string
	// }{
	// 	"world",
	// 	"mundo",
	// 	"monde",
	//}
	//fmt.Println("Hello, %s/%s/%s!", world.English, world.Spanish, world.French)

}

// Price is the cost of something in US cents.

type Price int64

// String is the string representation of a Price
// These should be represented in US Dollars
// Example: 2595 cents => $25.95
func (p Price) String() string {
	// TODO
	var number []Price
	count := 0
	if p == Price(0) {
		return "$0"
	}
	for p > Price(0) {
		number = append(number, p%(Price(10)))
		count += 1
		p /= Price(10)
	}
	var price_str string
	if count > 2 {
		price_str += "$"
		for i := count - 1; i >= 0; i-- {
			if i == 1 {
				price_str += "."
			}
			price_str += string(number[i] + 48)
		}
	} else {
		price_str += "$0."
		for i := count - 1; i >= 0; i-- {
			price_str += string(number[i] + 48)
		}
	}
	return price_str
}

// Prices is a map from an item to its price.
var Prices = map[string]Price{
	"eggs":          219,
	"bread":         199,
	"milk":          295,
	"peanut butter": 445,
	"chocolate":     150,
}

// RegisterItem adds the new item in the prices map.
// If the item is already in the prices map, a warning should be displayed to the user,
// but the value should be overwritten.
// Bonus (1pt) - Use the "log" package to print the error to the user
func RegisterItem(prices map[string]Price, item string, price Price) {
	// TODO
	ok := false
	_, ok = prices[item]
	if ok {
		fmt.Println("this a warrning")
	}
	prices[item] = price
}

// Cart is a struct representing a shopping cart of items.
type Cart struct {
	Items      []string
	TotalPrice Price
}

// hasMilk returns whether the shopping cart has "milk".
func (c *Cart) hasMilk() bool {
	// TODO
	for _, r := range c.Items {
		if r == "milk" {
			return true
		}
	}
	return false
}

// HasItem returns whether the shopping cart has the provided item name.
func (c *Cart) HasItem(item string) bool {
	// TODO
	for _, r := range c.Items {
		if r == item {
			return true
		}
	}
	return false
}

// AddItem adds the provided item to the cart and update the cart balance.
// If item is not found in the prices map, then do not add it and print an error.
// Bonus (1pt) - Use the "log" package to print the error to the user
func (c *Cart) AddItem(item string) {
	// TODO
	if c.HasItem(item) {
		c.TotalPrice += Prices[item]
	} else {
		fmt.Println("error")
	}

}

// Checkout displays the final cart balance and clears the cart completely.
func (c *Cart) Checkout() {
	// TODO
	fmt.Println(c.TotalPrice)
	c.TotalPrice = 0
	var new_array []string
	c.Items = new_array
}
