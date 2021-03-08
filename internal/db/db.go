package db

import (
	"burger-api/internal/config"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var Client *mongo.Client
var DB *mongo.Database

const dbName = "burger"

func InitDb(config config.Config) {
	clientOptions := options.Client()
	clientOptions.SetAuth(options.Credential{Username: config.Database.User, Password: config.Database.Password})
	clientOptions.ApplyURI("mongodb://" + config.Database.Host + ":" + config.Database.Port + "/")
	var err error
	Client, err = mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = Client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	DB = Client.Database(dbName)
}
