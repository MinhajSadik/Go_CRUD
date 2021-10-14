package DBManager

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// var dbURL string = "mongodb://localhost:27017/?readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false"
var configErr = godotenv.Load()
var dbURL string = os.Getenv("DB_SOURCE_URL")
var SystemCollections DemoCollections

type DemoCollections struct {
	Student *mongo.Collection
}

func getMongoDbCollection() (*mongo.Client, error) {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(dbURL))

	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.Background(), readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

	return client, nil

}

func GetMongoDbCollection(DbName string, CollectionName string) (*mongo.Collection, error) {

	client, err := getMongoDbCollection()

	if err != nil {
		return nil, err
	}
	collection := client.Database(DbName).Collection(CollectionName)

	return collection, nil
}

func InitCollections() bool {
	if configErr != nil {
		return false
	}

	var err error
	SystemCollections.Student, err = GetMongoDbCollection("Demo", "student")

	return err == nil
}
