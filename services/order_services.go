package services

import (
	"fmt"
	"ordering_v2/domain"
	"ordering_v2/domain/models"
)

type orderServices struct {
	repo domain.OrderRepository
}

func NewOrderServices(repo domain.OrderRepository) OrderServices {
	return orderServices{repo: repo}
}

func (order_service orderServices) GetOrder(order_id string) (*domain.Order, error) {
	order, err := order_service.repo.FormID(order_id)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (order_service orderServices) CreateNewOrder(customer models.Customer, options []models.Option) (string, error) {
	order_id := order_service.repo.NextIdentity()
	order := domain.NewOrder()
	err := order.CreateOrderWithOptionAndCustomer(order_id, customer, options)
	if err != nil {
		return "", err
	}
	order_service.repo.Save(order)
	return order_id, nil
}

func (order_service orderServices) UpdateOrderOptions(order_id string, options []models.Option) error {
	order, err := order_service.repo.FormID(order_id)
	if err != nil {
		return err
	}
	err = order.UpdateOptionsToOrder(options)
	if err != nil {
		return err
	}

	return nil
}

func (order_service orderServices) SubmitOrder(order_id string) error {
	order, err := order_service.repo.FormID(order_id)
	fmt.Println(order)
	if err != nil {
		return err
	}
	err = order.Submit()
	if err != nil {
		return err
	}
	order_service.repo.Save(*order)
	return nil
}
