package handler

import (
	"ordering_v2/services"

	"github.com/gofiber/fiber/v2"
)

type orderHandler struct {
	order_service services.OrderServices
}

func NewOrderHanlder(order_service services.OrderServices) *orderHandler {
	return &orderHandler{order_service: order_service}
}

func (h *orderHandler) CreateNewOrderHanlder(c *fiber.Ctx) error {
	var order = CreateNewOrderScheme{}

	if err := c.BodyParser(&order); err != nil {
		return err
	}

	order_id, err := h.order_service.CreateNewOrder(order.Customer, order.Options)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(order_id)
}

func (h *orderHandler) GetOrderByID(c *fiber.Ctx) error {
	order_id := c.Params("order_id")
	order, err := h.order_service.GetOrder(order_id)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(order)
}

func (h *orderHandler) UpdateOptionsToOrder(c *fiber.Ctx) error {
	var options OrderUpdateOptions
	order_id := c.Params("order_id")
	err := c.BodyParser(&options)
	if err != nil {
		return err
	}
	if err = h.order_service.UpdateOrderOptions(order_id, options.Options); err != nil {
		return err
	}
	return c.JSON(order_id)
}

func (h *orderHandler) SubmitOrder(c *fiber.Ctx) error {
	order_id := c.Params("order_id")
	err := h.order_service.SubmitOrder(order_id)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusNoContent).SendString("Order submitted")
}
