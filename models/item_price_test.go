package models_test

import (
	"fmt"
	"testing"

	"github.com/go_apps_101/models"
)

func TestItemList(t *testing.T) {
	const id = "item-123"
	const name = "Burger"
	const price = float64(456.76)
	const currency = "GBP"
	item := models.Item{
		ID:           id,
		Name:         name,
		Price:        price,
		CurrencyCode: currency,
	}

	expectedListing := fmt.Sprintf("%s: %v%s", item.Name, item.Price, item.CurrencyCode)
	listing := item.PrintWithPrice()

	fmt.Println("*******", listing)

	if listing != expectedListing {
		t.Errorf("item listing incorrect, got: %v, want: %v", listing, expectedListing)
	}
}
