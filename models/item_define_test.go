package models_test

import (
	"testing"

	"github.com/go_apps_101/models"
)

func TestItemFields(t *testing.T) {
	const id = "item-123"
	const name = "Burger"
	const price = 456.76
	const currency = "GBP"
	item := models.Item{
		ID:           "item-123",
		Name:         name,
		Price:        price,
		CurrencyCode: currency,
	}

	if item.ID != id {
		t.Errorf("item id incorrect, got: %v, want: %v", item.ID, id)
	}
	if item.Name != name {
		t.Errorf("item name incorrect, got: %v, want: %v", item.Name, name)
	}
	if item.Price != price {
		t.Errorf("item price incorrect, got: %v, want: %v", item.Price, price)
	}
	if item.CurrencyCode != currency {
		t.Errorf("item currency code incorrect, got: %v, want: %v", item.CurrencyCode, currency)
	}
}
