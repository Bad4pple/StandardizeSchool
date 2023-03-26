package services

import (
	"ordering/domain"
	"ordering/domain/models"
)

type OrderService struct {
	repo domain.OrderRepository
}

func InitializeOrderService(repo domain.OrderRepository) OrderService {
	return OrderService{
		repo: repo,
	}
}

func (u OrderService) CreateNewOrder(customer models.Customer, option []models.Option) (*models.CodeID, error) {
	order := domain.InitializeOrder()
	order_id, err := order.CreateNewOrderWithOptions(customer, option)
	if err != nil {
		return nil, err
	}
	_, err = u.repo.Create(order)
	if err != nil {
		return nil, err
	}
	return order_id, nil
}

func (u OrderService) SubmitOrder(order_id models.CodeID) error {
	order, err := u.repo.FormID(order_id)
	if err != nil {
		return err
	}

	err = order.Submit()
	if err != nil {
		return err
	}

	err = u.repo.Save(*order)
	if err != nil {
		return err
	}

	return nil
}

func (u OrderService) GetOrder(order_id models.CodeID) (*domain.Order, error) {
	order, err := u.repo.FormID(order_id)
	if err != nil {
		return nil, err
	}
	return order, nil
}
