package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/souravdev-eng/mogoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://sourav:Iamdeveloper1$@cluster0.h30vf4u.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

// !Most Important
var collection *mongo.Collection

func init() {
	// client option
	clientOption := options.Client().ApplyURI(connectionString)
	// connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB connected successfully...")
	collection = client.Database(dbName).Collection(colName)
	//
	fmt.Println("Collection instance is ready...")
}

// MONGODB helpers -file

// insert one record
func insertOneMovie(movie model.Netflix) {
	data, err := collection.InsertOne(context.Background(), movie)
	ThrowError(err)
	fmt.Println("Inserted 1 movie", data.InsertedID)
}

// Update one
func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

// Delete one
func deleteOneMovie(movieId string) int64 {
	id, err := primitive.ObjectIDFromHex(movieId)
	ThrowError(err)
	filter := bson.M{"_id": id}
	data, err := collection.DeleteOne(context.Background(), filter)
	ThrowError(err)
	fmt.Println("Data deleted successfully", data.DeletedCount)
	return data.DeletedCount
}

// get all movie
func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	ThrowError(err)
	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		ThrowError(err)
		movies = append(movies, movie)
	}

	defer cursor.Close(context.Background())
	return movies
}

// Delete all
func deleteAllMovie() int64 {
	data, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	ThrowError(err)

	fmt.Println("All data deleted successfully", data.DeletedCount)
	return data.DeletedCount
}

func ThrowError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//? Actual controller

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-ww-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-ww-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix

	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-ww-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-ww-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}
