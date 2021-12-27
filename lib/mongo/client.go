package mongo

import (
	"context"
	"with_coffee/lib/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoCollection(ctx context.Context, collection string) *mongo.Collection {
	cnf, _ := config.LoadConfig()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cnf.Mongo.Uri))

	if err != nil {
		panic(err)
	}

	result := client.Database(cnf.Mongo.Database).Collection(collection)

	return result
}
