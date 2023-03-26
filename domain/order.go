package domain

import (
	"errors"
	"ordering/domain/models"
	"reflect"
)

const (
	prefix = "order"
)

var (
	ErrCustomerOrOptionsMustNotBeEmpty = errors.New("customer or options must not be empty")
	ErrOrderHasAlreadyBeenComfirmed    = errors.New("the order has already been confirmed")
)

type Order struct {
	OrderID     models.CodeID   `json:"order_id" bson:"order_id"`
	Customer    models.Customer `json:"customer" bson:"customer"`
	Product     models.CodeID   `json:"product_id" bson:"product_id"`
	Options     []models.Option `json:"options" bson:"options"`
	TotalPrice  float64         `json:"total_price" bson:"total_price"`
	IsSubmitted bool            `json:"is_submitted" bson:"is_submitted"`
}

func InitializeOrder() Order {
	return Order{}
}
func (ordering *Order) CreateNewOrderWithOptions(customer models.Customer, options []models.Option) (*models.CodeID, error) {

	if reflect.DeepEqual(customer, models.Customer{}) || options == nil {
		return nil, ErrCustomerOrOptionsMustNotBeEmpty
	}

	var price float64
	for _, o := range options {
		price += o.Price
	}

	order_id := models.NewCodeID(prefix)
	ordering.OrderID = order_id
	ordering.Customer = customer
	ordering.Options = options
	ordering.TotalPrice = price
	return &order_id, nil
}

func (ordering *Order) Submit() error {
	if ordering.IsSubmitted {
		return ErrOrderHasAlreadyBeenComfirmed
	}
	ordering.IsSubmitted = true
	return nil
}
