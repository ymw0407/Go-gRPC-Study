package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/ymw0407/GoLang-Studying/data"
	postpb "github.com/ymw0407/GoLang-Studying/protos/v1/post"
	userpb "github.com/ymw0407/GoLang-Studying/protos/v1/user"
	client "github.com/ymw0407/GoLang-Studying/simple"
)

const portNumber = "9001"

type postServer struct {
	postpb.PostServer
	userCli userpb.UserClient
}

func (s *postServer) ListPostsByUserId(ctx context.Context, req *postpb.ListPostsByUserIdRequest) (*postpb.ListPostsByUserIdResponse, error) {
	userID := req.UserId

	resp, err := s.userCli.GetUser(ctx, &userpb.GetUserRequest{UserId: userID})
	if err != nil {
		return nil, err
	}

	var postMessages []*postpb.PostMessage
	for _, up := range data.UserPosts {
		if up.UserID != userID {
			continue
		}

		for _, p := range up.Posts {
			p.Author = resp.UserMessage.Name
		}

		postMessages = up.Posts
		break
	}

	return &postpb.ListPostsByUserIdResponse{
		PostMessages: postMessages,
	}, nil
}

func (s *postServer) ListPosts(ctx context.Context, req *postpb.ListPostsRequest) (*postpb.ListPostsResponse, error) {
	var postMessages []*postpb.PostMessage
	for _, up := range data.UserPosts {
		resp, err := s.userCli.GetUser(ctx, &userpb.GetUserRequest{UserId: up.UserID})
		if err != nil {
			return nil, err
		}

		for _, p := range up.Posts {
			p.Author = resp.UserMessage.Name
		}

		postMessages = append(postMessages, up.Posts...)
	}

	return &postpb.ListPostsResponse{
		PostMessages: postMessages,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":"+portNumber)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	userCli := client.GetUserClient("localhost:9000")
	grpcServer := grpc.NewServer()
	postpb.RegisterPostServer(grpcServer, &postServer{
		userCli: userCli,
	})

	log.Printf("start gRPC server on %s port", portNumber)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
