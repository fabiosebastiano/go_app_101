package models_test

import (
	"fmt"
	"testing"

	"github.com/go_apps_101/models"
)

func TestOrder_GetTotal(t *testing.T) {
	burger := models.Item{
		Price: 14.0,
	}
	pizza := models.Item{
		Price: 4.0,
	}
	or := models.Order{
		Items: []*models.Item{&burger, &pizza},
	}
	expectedTotal := burger.Price + pizza.Price

	total := or.Total()

	fmt.Println(total, " Vs ", expectedTotal)

	if total != expectedTotal {
		t.Fatalf("order total incorrect, got: %v, want: %v", total, expectedTotal)
	}
}
