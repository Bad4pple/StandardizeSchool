package adapter

import (
	"context"
	"fmt"
	"ordering/domain"
	"ordering/domain/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBRepositoryAdapter struct {
	collection *mongo.Collection
}

func InitializeMongoDBRepository(connection_string, database, collection, username, password string) (domain.OrderRepository, error) {
	clientOption := options.Client().ApplyURI(connection_string)

	if username != "" && password != "" {
		credentials := options.Credential{
			Username: username,
			Password: password,
		}
		clientOption.Auth = &credentials
	}

	client, err := mongo.NewClient(clientOption)
	if err != nil {
		return nil, err
	}

	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return &MongoDBRepositoryAdapter{
		collection: client.Database(database).Collection(collection),
	}, nil
}

func (repo *MongoDBRepositoryAdapter) Create(order domain.Order) (*models.CodeID, error) {
	ctx := context.Background()
	fmt.Println(order)
	_, err := repo.collection.InsertOne(ctx, order)
	if err != nil {
		return nil, err
	}
	return &order.OrderID, nil
}
func (repo *MongoDBRepositoryAdapter) FormID(order_id models.CodeID) (*domain.Order, error) {
	ctx := context.Background()
	var order = domain.Order{}

	err := repo.collection.FindOne(ctx, bson.M{"order_id": order_id}).Decode(&order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}
func (repo *MongoDBRepositoryAdapter) Save(order domain.Order) error {
	ctx := context.Background()
	filter := bson.M{"order_id": order.OrderID}
	update := bson.M{"$set": order}
	_, err := repo.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}
