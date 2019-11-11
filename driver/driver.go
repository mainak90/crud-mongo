package driver

import (
	"context"
	"crud-mongo/utils"
	"fmt"
	"os"

	mongo "go.mongodb.org/mongo-driver/mongo"
	options "go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() *mongo.Database {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URL"))

	client, err := mongo.Connect(context.TODO(), clientOptions)

	utils.LogFatal(err)

	fmt.Println(clientOptions)

	err = client.Ping(context.TODO(), nil)

	utils.LogFatal(err)

	return client.Database(os.Getenv("MONGO_DB"))
}
