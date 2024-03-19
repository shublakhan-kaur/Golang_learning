package helper

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://golang:golang%40123@cluster0.yhpcvw1.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbName = "netflix"
const collectionName = "watchlist"

var collection *mongo.Collection

func GetDBConnection() *mongo.Collection {

	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection is successful")
	collection = client.Database(dbName).Collection(collectionName)
	//log.Println(collection)
	fmt.Println("Collection instance is ready")
	return collection

}
