package model

import (
	"burger-api/internal/db"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const BurgerColName = "burgers"

var ctxBG = context.Background()

func burgerCol() *mongo.Collection {
	return db.Client.Database("burger").Collection(BurgerColName)
}

type Burger struct {
	ID           primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name         string             `json:"name"`
	Ingredients  []string           `json:"ingredients"`
	Notes        []string           `json:"notes"`
	Source       string             `json:"source"`
	Instructions []string           `json:"instructions"`
	B64Image     string             `json:"encodedImage"`
}

func FindAll(page Page, nameEq string) ([]Burger, error) {
	filter := bson.M{}
	if len(nameEq) != 0 {
		filter = bson.M{ "name" : nameEq }
	}
	cursor, err := burgerCol().Find(ctxBG, filter, &options.FindOptions{Limit: &page.PerPage, Skip: &page.Page})
	if err != nil {
		return nil, err
	}
	var burgers []Burger
	if err = cursor.All(ctxBG, &burgers); err != nil {
		return nil, err
	}
	return burgers, nil
}

func FindOne(id string) (*Burger, error) {
	var result Burger

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = burgerCol().FindOne(ctxBG, bson.M{"_id": objId}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func FindRandom() (*Burger, error) {
	cursor, err := burgerCol().Aggregate(ctxBG, bson.A{bson.M{"$sample": bson.M{"size": 1}}})
	if err != nil {
		return nil, err
	}
	var burgers []Burger
	if err = cursor.All(ctxBG, &burgers); err != nil {
		return nil, err
	}
	if burgers == nil || len(burgers) == 0 {
		return nil, errors.New("no burgers available")
	}
	return &burgers[0], nil
}

func InsertOne(burger *Burger) (*mongo.InsertOneResult, error) {
	err := validateBurger(burger)
	if err != nil {
		return nil, err
	}
	burger.ID = primitive.NewObjectID()
	insrBurger, err := burgerCol().InsertOne(ctxBG, &burger)
	return insrBurger, err
}

func validateBurger(burger *Burger) error {
	if len(burger.Name) == 0 {
		return errors.New("burger name is required")
	}
	if len(burger.Ingredients) == 0 {
		return errors.New("burger ingredients are required")
	}
	return nil
}
