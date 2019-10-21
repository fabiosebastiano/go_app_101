package main

import (
	"fmt"

	"github.com/go_apps_101/models"
)

func main() {
	// EXERCISE 2: Create & print some items
	// ------------------------------------

	var itemA models.Item
	fmt.Printf("Item solo dichiarato %v \n", itemA)

	itemB := models.Item{}
	fmt.Printf("Item valorizzato con struct vuoto %+v \n", itemB)

	itemC := models.Item{
		ID:           "item Uno",
		Name:         "primo item",
		Price:        10,
		CurrencyCode: "12",
	}

	fmt.Printf("Item inizializzato completamente %+v \n", itemC)

}
