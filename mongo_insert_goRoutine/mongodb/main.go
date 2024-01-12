package mongodb

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	url  string
	name string

	Client    *mongo.Client
	ClientErr error

	UserColl   = &mongo.Collection{}
	SchoolColl = &mongo.Collection{}
)

func Initialize() {
	url = os.Getenv("MONGODB_URL")
	name = os.Getenv("MONGODB_NAME")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(url).SetServerAPIOptions(serverAPI)

	Client, ClientErr = mongo.Connect(context.TODO(), opts)

	if ClientErr != nil {
		panic(ClientErr)
	} else {
		log.Printf("MongoDB connected successfully!")
	}

	defineCollection()
}

func defineCollection() {
	UserColl = Client.Database(name).Collection("userTest")
	SchoolColl = Client.Database(name).Collection("schoolTest2")
}
