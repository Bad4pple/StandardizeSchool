package adapter

import (
	"context"
	"fmt"
	"ordering_v2/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OrderRepositoryMongoDB struct {
	collection *mongo.Collection
}

func NewOrderRepositoryMongoDB(connection_string, database, collection, username, password string) (domain.OrderRepository, error) {
	client_option := options.Client().ApplyURI(connection_string)
	if username != "" && password != "" {
		credentials := options.Credential{
			Username: username,
			Password: password,
		}
		client_option.Auth = &credentials
	}

	client, err := mongo.NewClient(client_option)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return OrderRepositoryMongoDB{collection: client.Database(database).Collection(collection)}, nil
}

func (repo OrderRepositoryMongoDB) NextIdentity() string {
	_id := primitive.NewObjectID()
	order_id := fmt.Sprintf("OR-%s", _id.Hex())
	return order_id
}
func (repo OrderRepositoryMongoDB) FormID(order_id string) (*domain.Order, error) {
	ctx := context.Background()
	var order = domain.NewOrder()

	err := repo.collection.FindOne(ctx, bson.M{
		"order_id": order_id,
	}).Decode(&order)

	if err != nil {
		return nil, err
	}

	return &order, nil
}
func (repo OrderRepositoryMongoDB) Save(entity domain.Order) error {
	ctx := context.Background()
	filter := bson.M{"order_id": entity.OrderID}

	// Check if entity exists in database
	count, err := repo.collection.CountDocuments(ctx, filter, nil)
	if err != nil {
		return err
	}

	if count > 0 {
		// Entity exists in database, update it
		update := bson.M{
			"$set": entity,
		}
		_, err = repo.collection.UpdateOne(ctx, filter, update)
		if err != nil {
			return err
		}
	} else {
		// Entity does not exist in database, insert new document
		_, err := repo.collection.InsertOne(ctx, entity)
		if err != nil {
			return err
		}
	}

	return nil
}
