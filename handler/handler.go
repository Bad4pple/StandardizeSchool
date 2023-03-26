package handler

import (
	"ordering/domain/models"
	"ordering/services"

	"github.com/gofiber/fiber/v2"
)

type orderHandler struct {
	order_service services.OrderService
}

func InitializeOrderHandler(order_service services.OrderService) *orderHandler {
	return &orderHandler{order_service: order_service}
}

func (h *orderHandler) CreateNewOrderHandler(c *fiber.Ctx) error {
	var (
		RequestOrder  = CreateNewOrderRequestModel{}
		ResponseOrder = CreateNewOrderResponseModel{}
	)
	if err := c.BodyParser(&RequestOrder); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Sorry, it seems your request body is wrong")
	}
	order_id, err := h.order_service.CreateNewOrder(RequestOrder.Customer, RequestOrder.Options)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Sorry, your request is not success, please try again")
	}
	ResponseOrder.OrderID = *order_id
	return c.JSON(ResponseOrder)
}

func (h *orderHandler) GetOrderByID(c *fiber.Ctx) error {
	order_id := c.Params("order_id")
	order, err := h.order_service.GetOrder(models.CodeID(order_id))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Sorry, it seems that order id is not found ")
	}

	return c.Status(fiber.StatusOK).JSON(order)
}

func (h *orderHandler) UpdateOptionsToOrder(c *fiber.Ctx) error {
	order_id := c.Params("order_id")
	return c.JSON(order_id)
}

func (h *orderHandler) SubmitOrder(c *fiber.Ctx) error {
	orderID := c.Params("order_id")
	err := h.order_service.SubmitOrder(models.CodeID(orderID))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Sorry, We can't not submission your order, please check your order that is submitted, isn't")
	}
	return c.Status(fiber.StatusNoContent).SendString("Order submitted")
}
