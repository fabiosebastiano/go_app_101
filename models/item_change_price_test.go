package models_test

import (
	"fmt"
	"testing"

	"github.com/go_apps_101/models"
)

func TestItemChangePrice(t *testing.T) {
	const id = "item-123"
	const name = "Burger"
	const price = float64(456.76)
	const newPrice = float64(789.76)
	const currency = "GBP"
	const newCurrency = "EUR"

	item := &models.Item{
		ID:           id,
		Name:         name,
		Price:        price,
		CurrencyCode: currency,
	}

	fmt.Println("vecchio prezzo", item.Price)
	fmt.Println("vecchia concurrency", item.CurrencyCode)
	item.ChangePrice(newPrice, newCurrency)

	fmt.Println("Nuovo prezzo", item.Price)
	fmt.Println("Nuova concurrency", item.CurrencyCode)

	if item.ID != id {
		t.Errorf("item id incorrect, got: %v, want: %v", item.ID, id)
	}
	if item.Name != name {
		t.Errorf("item name incorrect, got: %v, want: %v", item.Name, name)
	}
	if item.Price != newPrice {
		t.Errorf("item price incorrect, got: %v, want: %v", item.Price, newPrice)
	}
	if item.CurrencyCode != newCurrency {
		t.Errorf("item currency code incorrect, got: %v, want: %v", item.CurrencyCode, newCurrency)
	}
}
