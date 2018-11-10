package main

import (
	"context"
	"log"
	"time"
  "strconv"

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
	c := pb.NewForgotPasswordClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	response, err := c.ForgotPassword(ctx, &pb.ForgotPasswordRequest{Email: "mag030@ucsd.edu"})

	log.Println("Success: "+strconv.FormatBool(response.Success))
}
