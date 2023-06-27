package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	db "grpc-mongo/db"
	userpb "grpc-mongo/protos/v1/user"
)

const portNumber = "9000"

type signServer struct {
	userpb.UserServiceServer
}

func (s *signServer) SignUp(ctx context.Context, req *userpb.SignUpRequest) (*userpb.SignUpResponse, error) {
	user_id, user_name, user_gender, user_email, user_password := req.User.Id, req.User.Name, req.User.Gender, req.User.Email, req.Password
	newUser := db.User{user_id, user_name, user_gender, user_email, user_password}

	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found")
		return &userpb.SignUpResponse{
			Success: false,
			Message: ".env file not gound",
		}, nil
	}
	MONGODB_URI := os.Getenv("MONGODB_URI")

	db.MongoConnection(MONGODB_URI)

	client := db.MongoConnection(MONGODB_URI)
	defer db.MongoDisconnection(client)

	db.MongoUserInsert(newUser, client.Database("grpc").Collection("users"))

	return &userpb.SignUpResponse{
		Success: true,
		Message: "User Seccessfully Inserted",
	}, nil

}

func main() {
	lis, err := net.Listen("tcp", ":"+portNumber)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, &signServer{})

	log.Printf("start gRPC server on %s port", portNumber)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
