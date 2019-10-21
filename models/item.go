package models

import "fmt"

// EXERCISE 1: Define Item struct
// ----------------------------
type Item struct {
	ID           string
	Name         string
	Price        float64
	CurrencyCode string
	// define ID as string,
	// Name as string
	// Price as float64,
	// CurrencyCode as string
}

// EXERCISE 3: Print the price of a menu Item
// ----------------------------------------
func (i Item) PrintWithPrice() string {
	//present the item in the format Name: PriceCurrencyCode
	//NOTE: fmt.Sprintf allows us to format strings just like with fmt.Println
	//https://golang.org/pkg/fmt/#Sprintf
	// Use the format string "%s: %v%s"
	return fmt.Sprintf("%s: %v%s", i.Name, i.Price, i.CurrencyCode)
}

// EXERCISE 4: Change price of a menu item
// -------------------------------------
func (i *Item) ChangePrice(price float64, code string) {
	//Should this method take a pointer receiver?
	//assign new price & currency code here
	i.Price = price
	i.CurrencyCode = code
}
