package models

import (
	"arena"
	"github.com/emreisler/go-arena-tracking/constants"
	"github.com/emreisler/go-arena-tracking/tracking"
)

// Order struct (private)
type Order struct {
	ID       int
	Quantity int
	Price    float64
	Items    []string
}

func NewOrder(ar *arena.Arena, id, quantity int, price float64, items []string) *Order {
	if constants.GcOff {
		return &Order{
			ID:       id,
			Quantity: quantity,
			Price:    price,
			Items:    items,
		}
	}

	o := arena.New[Order](ar)
	o.ID = id
	o.Quantity = quantity
	o.Price = price
	o.Items = items

	tracking.TrackHeapAlloc("Order", o)
	return o
}
