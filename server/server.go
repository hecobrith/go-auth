package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	//"github.com/golang/protobuf/proto"
	pb "github.com/hecobrith/go-auth/protofiles"
)

const (
	port = ":9000"
)

type server struct{}

func (s *server) Login(ctx context.Context, in *pb.UserData) (*pb.SessionResponse, error) {
	log.Printf("Login attempt....")
	log.Printf("Amount: %s, From A/c:%s", in.Username, in.Password)
	// Do database logic here....
	return &pb.SessionResponse{Token: "some token",Message:"done wright"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLoginServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
