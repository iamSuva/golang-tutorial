package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	

	"github.com/gorilla/mux"
	"github.com/iamSuva/mongoapi/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongourl = "mongodb+srv://suvadip:suvadip632@cluster0.f0dozwy.mongodb.net/"
const DBname = "netflix"
const collectionName = "watchlist"

// most important
var collection *mongo.Collection

// connection
func init() { // it will run for the first time
	//client option
	clientOption := options.Client().ApplyURI(mongourl)
	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	//check connection
	fmt.Println("connected to database")

	collection = client.Database(DBname).Collection(collectionName)
	fmt.Println("collection instance is ready")
}

// mongodb helpers -file

//insert 1 record

func insertOneMovie(movie models.Netfilx) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	//context.background():It’s an empty context with no deadline, no cancellation, and no values.
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db with id ", inserted.InsertedID)

}

func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count : ", result.ModifiedCount)
}
func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("movie got delete with delete count : ", deleteCount.DeletedCount)

}

//delete all records from mongodb

func deleteAllMovies() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleted count : ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount;
}

// get all records from mongodb
func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M
	// This indicates a slice (dynamic list) of BSON maps (primitive.M).
	// Each element in the slice is a BSON map,
	//  allowing you to store and manipulate multiple BSON documents in a structured way.
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
func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	movies := getAllMovies()
	json.NewEncoder(w).Encode(movies)
}
 func CreateMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Controll-Allow-Methods","POST")

	var movie models.Netfilx
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)

 }
 func  MarkedAsWatched(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","PUT")
	params :=mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
 }
 func DeleteOneMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")
     params :=mux.Vars(r)
	 
	 deleteOneMovie(params["id"])
	 json.NewEncoder(w).Encode(params["id"])

 }
 func DeleteAllMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")
    count:= deleteAllMovies();
	 json.NewEncoder(w).Encode(count)

 }
