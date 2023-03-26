package main

import (
	"log"
	"ordering/adapter"
	"ordering/handler"
	"ordering/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	order_repo, err := adapter.InitializeMongoDBRepository(
		"mongodb://root:example@127.0.0.1:27017",
		"StandardizeWarehourse",
		"orders",
		"root",
		"example",
	)
	if err != nil {
		log.Fatal(err)
	}

	order_service := services.InitializeOrderService(order_repo)
	order_handler := handler.InitializeOrderHandler(order_service)

	app := fiber.New()
	api := app.Group("/api/orders")

	api.Post("/", order_handler.CreateNewOrderHandler)
	api.Get("/:order_id", order_handler.GetOrderByID)
	api.Put("/:order_id/options", order_handler.SubmitOrder)
	api.Put("/:order_id/submission", order_handler.SubmitOrder)

	app.Listen(":8000")
}
