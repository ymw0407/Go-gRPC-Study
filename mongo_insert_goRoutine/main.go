package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/ymw0407/Go-gRPC-Study/mongo_insert_goRoutine/mongodb"
	"github.com/ymw0407/Go-gRPC-Study/mongo_insert_goRoutine/mongodb/find"
	insertfriend "github.com/ymw0407/Go-gRPC-Study/mongo_insert_goRoutine/mongodb/insertFriendCollTest"
)

const ENV_FILE = ".env"

func main() {
	err := godotenv.Load(ENV_FILE)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	initialize()
	codes := find.FindSchool()
	// insert.SharedCode = codes
	// insert.InsertTest(codes)
	// insertLogin.InsertTest()
	insertfriend.InsertTest(codes)
}

func initialize() (err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered from panic during initialization:", r)
			err = fmt.Errorf("panic: %v", r)
		}
	}()

	mongodb.Initialize()
	return nil
}
