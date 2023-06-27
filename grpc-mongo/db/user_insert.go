package db

import (
	"context"
	// "hash"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "golang.org/x/crypto/bcrypt"
)

type User struct {
	Id string
	Name string
	Gender string
	Email string
	Password string
}

func MongoConnection(MONGODB_URI string) (client *mongo.Client) {
	if MONGODB_URI == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MONGODB_URI))
	if err != nil {
		panic(err)
	}
	return client
}

func MongoDisconnection(client *mongo.Client) {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

func MongoUserInsert(user User, coll *mongo.Collection) {
	coll.InsertOne(context.TODO(), user)
}
