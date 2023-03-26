package main

import (
	"log"
	"ordering_v2/adapter"
	"ordering_v2/domain"
	"ordering_v2/handler"
	"ordering_v2/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	order_repository := initMongoConfig()
	order_service := services.NewOrderServices(order_repository)
	order_handler := handler.NewOrderHanlder(order_service)
	app := fiber.New()
	api := app.Group("/api/orders")

	api.Post("/", order_handler.CreateNewOrderHanlder)
	api.Get("/:order_id", order_handler.GetOrderByID)
	api.Put("/:order_id", order_handler.UpdateOptionsToOrder)
	api.Put("/:order_id/submission", order_handler.SubmitOrder)

	app.Listen(":8000")
}

func initMongoConfig() domain.OrderRepository {

	order_repository, err := adapter.NewOrderRepositoryMongoDB(
		"mongodb://root:example@127.0.0.1:27017",
		"StandardizeWarehourse",
		"orders",
		"root",
		"example",
	)
	if err != nil {
		log.Fatal(err)
	}

	return order_repository
}
