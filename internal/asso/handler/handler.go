package handler

import (
	"asso-api/internal/config"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Search struct {
	Q string `form:"q" json:"q"`
}

func GetAssociations(w http.ResponseWriter, r *http.Request) {
	coll := config.Client.Database("asso").Collection("asso")

	vars := mux.Vars(r)

	search := "\"" + strings.Join(strings.Split(vars["q"], " "), "\"") + "\""
	filter := bson.D{{"$text", bson.D{{"$search", search}}}}
	sort := bson.D{{"score", bson.D{{"$meta", "textScore"}}}}
	projection := bson.D{{"name", 1}, {"idRna", 1}, {"score", bson.D{{"$meta", "textScore"}}}, {"zipCode", 1}, {"city", 1}, {"_id", 0}}
	opts := options.Find().SetLimit(50).SetSort(sort).SetProjection(projection)
	cursor, err := coll.Find(context.TODO(), filter, opts)
	if err != nil {
		panic(err)
	}
	var results []primitive.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	for cursor.Next(context.Background()) {
		var result bson.M
		e := cursor.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		results = append(results, result)

	}

	json.NewEncoder(w).Encode(results)
}
