package main

import (
	"context"
	"log"
	"time"

	pb "github.com/UCSDIOWA/login-api/protos"
	"google.golang.org/grpc"
)

func main() {
	// Connect to the server
	conn, err := grpc.Dial("127.0.0.1:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect to the server: %v", err)
	}
	defer conn.Close()
	c := pb.NewLoginAPIClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	signUpResp, err := c.SignUp(ctx, &pb.SignUpRequest{Email: "tea.noreply@gmail.com", Password: "password", FirstName: "Tea",
		LastName: "Test"})
	if err != nil {
		log.Println(err)
	}
	log.Println("Sign up response: ", signUpResp.Success)

	logInResp, err := c.LogIn(ctx, &pb.LogInRequest{Email: "tea.noreply@gmail.com", Password: "password"})
	if err != nil {
		log.Println(err)
	}
	log.Println("Login response with correct credentials: ", logInResp.Success)

	logInResp, err = c.LogIn(ctx, &pb.LogInRequest{Email: "tea.noreply@gmail.com", Password: "wrongpassword"})
	if err != nil {
		log.Println(err)
	}
	log.Println("Login response with wrong credentials: ", logInResp.Success)
}
