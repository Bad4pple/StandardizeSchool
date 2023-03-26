package domain

import (
	"fmt"
	"ordering_v2/domain/models"
	"reflect"
)

var (
	error_order_had_submitted             = models.NewOrderHadSubmittedError()
	error_customer_field_is_empty         = models.NewOrderCustomerFieldIsEmptyError()
	error_options_less_than_equal_to_zero = models.NewOrderOptionsFieldLessThanEqualToZeroError()
)

type Order struct {
	OrderID     string          `json:"order_id" bson:"order_id"`
	Options     []models.Option `json:"options" bson:"options"`
	Customer    models.Customer `json:"customer" bson:"customer"`
	IsSubmitted bool            `bson:"is_submitted"`
}

func NewOrder() Order {
	return Order{}
}

func (ordering *Order) CreateOrderWithOptionAndCustomer(order_id string, customer models.Customer, options []models.Option) error {

	customer_is_empty := reflect.DeepEqual(customer, models.Customer{})

	if customer_is_empty {
		return error_customer_field_is_empty
	}
	if len(options) <= 0 {
		return error_options_less_than_equal_to_zero
	}

	ordering.OrderID = order_id
	ordering.Customer = customer
	ordering.Options = options

	// append event
	return nil
}

func (ordering *Order) UpdateOptionsToOrder(options []models.Option) error {

	if ordering.IsSubmitted {
		return error_order_had_submitted
	}

	ordering.Options = options

	// append event
	return nil
}

func (ordering *Order) Submit() error {
	fmt.Println(ordering.IsSubmitted)
	if ordering.IsSubmitted {
		return error_order_had_submitted
	}
	ordering.IsSubmitted = true

	// append event
	return nil
}
