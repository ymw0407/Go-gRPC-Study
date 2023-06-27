package db_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"

	db "grpc-mongo/db"
)

func Test_MongoConnection(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		t.Error(".env file not found")
	}
	MONGODB_URI := os.Getenv("MONGODB_URI")
	db.MongoConnection(MONGODB_URI)
}

func Test_MongoDisconnection(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		t.Error(".env file not found")
	}
	MONGODB_URI := os.Getenv("MONGODB_URI")

	client := db.MongoConnection(MONGODB_URI)
	defer db.MongoDisconnection(client)
}

func Test_MongoUserInsert(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		t.Error(".env file not found")
	}
	MONGODB_URI := os.Getenv("MONGODB_URI")

	client := db.MongoConnection(MONGODB_URI)
	defer db.MongoDisconnection(client)

	user := db.User{"민우", "yunminwo1211@gmail.com"}
	db.MongoUserInsert(user, client.Database("grpc").Collection("users"))
}