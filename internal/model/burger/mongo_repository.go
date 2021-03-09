package burger

import (
	"burger-api/internal/model"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	collection *mongo.Collection
}

var ctxBG = context.Background()

func NewMongoRepository(mdb *mongo.Database) MongoRepository {
	return MongoRepository{collection: mdb.Collection(BurgerColName)}
}

func (r MongoRepository) FindAll(page model.Page, nameEq string) ([]Burger, error) {
	filter := bson.M{}
	if len(nameEq) != 0 {
		filter = bson.M{"name": nameEq}
	}
	cursor, err := r.collection.Find(ctxBG, filter, &options.FindOptions{Limit: &page.PerPage, Skip: &page.Page})
	if err != nil {
		return nil, err
	}
	var burgers []Burger
	if err = cursor.All(ctxBG, &burgers); err != nil {
		return nil, err
	}
	return burgers, nil
}

func (r MongoRepository) FindOne(id string) (*Burger, error) {
	var result Burger

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = r.collection.FindOne(ctxBG, bson.M{"_id": objId}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r MongoRepository) FindRandom() (*Burger, error) {
	cursor, err := r.collection.Aggregate(ctxBG, bson.A{bson.M{"$sample": bson.M{"size": 1}}})
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

func (r MongoRepository) Save(burger *Burger) (*Burger, error) {
	err := validateBurger(burger)
	if err != nil {
		return nil, err
	}
	burger.ID = primitive.NewObjectID()
	insrtResult, err := r.collection.InsertOne(ctxBG, &burger)
	if err != nil {
		return nil, err
	}
	oid := insrtResult.InsertedID.(primitive.ObjectID).Hex()

	return r.FindOne(oid)
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
