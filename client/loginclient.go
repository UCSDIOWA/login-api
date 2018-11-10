package main

import (
	"context"
	"log"
	"time"

	pb "login-api/protos"
	"google.golang.org/grpc"
)

func main() {
	// Connect to the server
	conn, err := grpc.Dial("127.0.0.1:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect to the server: %v", err)
	}
	defer conn.Close()
	c := pb.NewLogInClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	response, err := c.LogIn(ctx, &pb.LogInRequest{Email: "test@tea.com", Password: "" })

	log.Println(response)
}
