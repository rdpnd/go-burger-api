package burger

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const BurgerColName = "burgers"

type Burger struct {
	ID           primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name         string             `json:"name"`
	Ingredients  []string           `json:"ingredients"`
	Notes        []string           `json:"notes"`
	Source       string             `json:"source"`
	Instructions []string           `json:"instructions"`
	B64Image     string             `json:"encodedImage"`
}
