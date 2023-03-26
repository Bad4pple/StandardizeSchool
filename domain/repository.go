package domain

import "ordering/domain/models"

type OrderRepository interface {
	Create(order Order) (*models.CodeID, error)
	FormID(order_id models.CodeID) (*Order, error)
	Save(order Order) error
}
