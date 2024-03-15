package service

import (
	"context"
	"fmt"
	"log"

	"github.com/shublakhan-kaur/Golang_learning/mongoapi/helper"
	"github.com/shublakhan-kaur/Golang_learning/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertOneMovie(movie model.Netflix) {
	result, err := helper.GetDBConnection().InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in the DB", result.InsertedID)
}

func UpdateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := helper.GetDBConnection().UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated records are: ", result.ModifiedCount)
}

func DeleteOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	deletedCount, err := helper.GetDBConnection().DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("No of movies deleted are: ", deletedCount)
}

func DeleteAllMovies() {
	deleteResult, err := helper.GetDBConnection().DeleteMany(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("No of movies deleted from the database are: ", deleteResult.DeletedCount)
}

func GetAllMovies() []primitive.M {
	cur, err := helper.GetDBConnection().Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	var movies []primitive.M
	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)

	}
	return movies
}
