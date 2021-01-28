package main

import (
	"log"

	pb "github.com/hecobrith/go-auth/protofiles"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:9000"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewLoginClient(conn)

	// Prepare data. Get this from clients like Frontend or App
	username := "hecobrith"
	password := "somepass"

	// Contact the server and print out its response.
	r, err := c.Login(context.Background(), &pb.UserData{Username:username,
	Password:password})
	if err != nil {
		log.Fatalf("Could not transact: %v", err)
	}
	log.Printf("Transaction confirmed: %t", r.Token)
}