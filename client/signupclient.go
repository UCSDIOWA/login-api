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
	c := pb.NewSignUpClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	response, err := c.SignUp(ctx, &pb.SignUpRequest{Email: "mag030@ucsd.edu", FirstName: "John",
                   LastName: "Smith", Secret: &pb.SignUpRequest_SecretCode{ Value: "12345", Sent: true } })

	log.Println(response.Success)
}
