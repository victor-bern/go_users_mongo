package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func GetDatabase() (*mongo.Database, *mongo.Client) {
	godotenv.Load()
	uri := fmt.Sprintf("mongodb+srv://%s:%s@teste.sy7ap.mongodb.net/Teste?retryWrites=true&w=majority",
		os.Getenv("MONGO_USER"), os.Getenv("MONGO_PASSWORD"))
	fmt.Print(uri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	return client.Database("teste"), client
}
