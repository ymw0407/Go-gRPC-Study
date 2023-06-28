package module

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LogIn struct {
	Id string
	Password string
}

func MongoUserLogInFind(user LogIn, coll *mongo.Collection) (User, error) {
	var result User
	err := coll.FindOne(context.TODO(), bson.D{{"id", user.Id}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		log.Printf("No document was found with the id %s\n", user.Id)
		return User{}, nil
	}
	if err != nil {
		panic(err)
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	if !Verify(result.Password, user.Password) {
		log.Printf("틀린비번... id %s\n", user.Id)
		return User{}, nil
	}
	fmt.Printf("%s\n", jsonData)
	return result, nil
	// var user_id, user_name, user_gender, user_email string
	// user_id, user_name, user_gender, user_email = jsonData.id, jsonData.name, jsonData.gender, jsonData.email

	// return User{jsonData.id, jsonData.name, jsonData.gender, jsonData.email, user.Password}, ""
}
