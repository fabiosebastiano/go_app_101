package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go_apps_101/db"
	"github.com/go_apps_101/handlers"
	"github.com/go_apps_101/models"
	"github.com/go_apps_101/repo"
)

func main() {

	burger := models.Item{
		ID:           "1",
		Name:         "Burger",
		Price:        float64(21.78),
		CurrencyCode: "EUR",
	}

	pizza := models.Item{
		ID:           "2",
		Name:         "Pizza",
		Price:        float64(21.78),
		CurrencyCode: "EUR",
	}

	orderRepo := &repo.OrderRepo{
		DB: db.InitDB(),
	}
	handler := &handlers.Handler{
		Repo: orderRepo,
	}

	// EXERCISE 8: Create & save an order
	// --------------------------------

	order := models.Order{
		Items:  []*models.Item{&pizza, &burger},
		Status: models.ACCEPTED,
	}
	fmt.Printf("%v", order)

	orderRepo.Upsert(&order)

	// Server config
	router := handlers.ConfigureServer(handler)
	fmt.Println("Listening on localhost:3000...")
	log.Fatal(http.ListenAndServe(":3000", router))

}
