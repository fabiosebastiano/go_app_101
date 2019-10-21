package models

// EXERCISE 5: Define the Order model
// ---------------------------------
type Order struct {
	ID     string
	Status OrderStatus
	Items  []*Item
}

// EXERCISE 6: Implement Order total
// -------------------------------
func (or Order) Total() float64 {
	var total float64
	for _, item := range or.Items {
		total += item.Price
	}

	return total
}
