package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"main.go/model"
)

const connectionString = "mongodb+srv://rishabh:rishabh27@cluster0.5enlb12.mongodb.net/"
const dbname = "mydb"
const collectionName = "watchlist"

// most imp part
var collection *mongo.Collection

func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connection
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")

	collection = client.Database(dbname).Collection(collectionName)

	fmt.Println("Collection Created")
}

//helper functions

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne((context.TODO()), movie)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted One Movie with ID", inserted.InsertedID)

}

func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	upadte := bson.M{"$set": bson.M{"watched": true}}
	_, err := collection.UpdateOne(context.Background(), filter, upadte)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated One Movie with ID", movieId)

}

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	movieCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie got delete with delete count", movieCount)
}

func deleteAllMovies() {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All movies got delete with delete count", deleteResult.DeletedCount)
}

func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M
	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cur.Close(context.Background())
	return movies
}

//Actual controllers

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreatreMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow_Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow_Methods", "PUT")

	params := mux.Vars(r)
	movieId := params["id"]
	updateOneMovie(movieId)
	json.NewEncoder(w).Encode(movieId)

}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("AllowControlControl-Allow_Methods", "DELETE")
	params := mux.Vars(r)
	movieId := params["id"]
	deleteOneMovie(movieId)
	json.NewEncoder(w).Encode(movieId)
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow_Methods", "DELETE")
	deleteAllMovies()
	json.NewEncoder(w).Encode("All Movies got delete")
}
