package handler

import "ordering/domain/models"

type CreateNewOrderRequestModel struct {
	Customer models.Customer `json:"customer"`
	Options  []models.Option `json:"options"`
}

type CreateNewOrderResponseModel struct {
	OrderID models.CodeID `json:"order_id"`
}
