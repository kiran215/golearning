package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mongo/model"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionstring = "mongodb+srv://kiransadaye:nxkIjERN6bgq7DwO@cluster0.d6ggfmv.mongodb.net/?retryWrites=true&w=majority"
const dbname = "netflix"
const colname = "watchlist"

var collection *mongo.Collection

//Connect with mongo db

// init method run very first time only one time
func init() {
	clientOption := options.Client().ApplyURI(connectionstring)

	//connet to mongo db
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongo db connection success")

	collection = client.Database(dbname).Collection(colname)

	fmt.Println("collection instance is ready")

}

// mongo helper
func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted one movie with id: ", inserted.InsertedID)

}

func updateOneMovie(movieId string) int64 {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"iswatched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count :", result.ModifiedCount)
	return result.ModifiedCount
}

func deleteOneMovie(movieId string) int64 {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("delet record count :", result.DeletedCount)
	return result.DeletedCount
}

func deleteAllMovie() {

	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("delet all record count :", deleteResult.DeletedCount)
}

// Get All moview from DB MONGO
func getAllMovies() []primitive.M {

	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cursor.Close(context.Background())

	return movies
}

// Actual Controller
func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)

}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}
func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applcation/json")

	params := mux.Vars(r)

	updateCount := updateOneMovie(params["id"])

	json.NewEncoder(w).Encode(updateCount)
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applcation/json")

	params := mux.Vars(r)

	deleteCount := deleteOneMovie(params["id"])

	json.NewEncoder(w).Encode(deleteCount)
}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applcation/json")
	// deleteCount :=
	deleteAllMovie()

	// json.NewEncoder(w).Encode(deleteCount)
}
