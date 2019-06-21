package mongo

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func GetMongoDb() *mongo.Database {
	if client == nil {
		client = connect()
		if client == nil {
			return GetMongoDb()
		}
	}
	return client.Database(os.Getenv("MONGO_CONNECTION_DBNAME"))
}

func connect() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_CONNECTION_URI")+os.Getenv("MONGO_CONNECTION_DBNAME")))
	// if err != nil {
	// 	fmt.Println(err)
	// }
	return client
}
