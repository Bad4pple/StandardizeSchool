package services

import (
	"ordering_v2/domain"
	"ordering_v2/domain/models"
)

type OrderServices interface {
	GetOrder(order_id string) (*domain.Order, error)
	CreateNewOrder(customer models.Customer, options []models.Option) (string, error)
	UpdateOrderOptions(order_id string, options []models.Option) error
	SubmitOrder(order_id string) error
}
