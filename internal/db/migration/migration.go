package migration

import (
	"burger-api/internal/config"
	"burger-api/internal/db"
	"burger-api/internal/model/burger"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"time"
)

func CreateCollections() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err := db.DB.CreateCollection(ctx, "burgers", &options.CreateCollectionOptions{
		Validator: bson.M{
			"$jsonSchema": bson.M{
				"bsonType": "object",
				"required": []string{"name", "ingredients"},
				"properties": bson.M{
					"name": bson.M{
						"bsonType":    "string",
						"description": "must be a string and is required",
					},
					"ingredients": bson.M{
						"bsonType": "array",
					},
				},
			}}})
	if err != nil {
		config.Logger.Printf("Burger collection was not created %v: %v\n", "burgers", err)
		return
	} else {
		insertFixtures()
	}
}

func insertFixtures() {
	file, err := ioutil.ReadFile("./dev/burgers_test_data.json")
	if err != nil {
		config.Logger.Printf("No burger fixture data was found %v", err)
		return
	}
	var data []burger.Burger

	err = json.Unmarshal(file, &data)
	if err != nil {
		config.Logger.Printf("Error while reading burger fixture data %v", err)
		return
	}

	for _, b := range data {
		_, err := burger.Repo.Save(&b)
		if err != nil {
			config.Logger.Printf("Error while inserting fixture data %v", err)
		}
	}

}
