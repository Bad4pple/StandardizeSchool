package handler

import "ordering_v2/domain/models"

type CreateNewOrderScheme struct {
	Customer models.Customer `json:"customer"`
	Options  []models.Option `json:"options"`
}

type OrderUpdateOptions struct {
	Options []models.Option `json:"options"`
}
