package main

import (
	"context"
	"log"
	"net"

	pb "github.com/UCSDIOWA/login-api/protos"
	"github.com/globalsign/mgo"
	"google.golang.org/grpc"
)

type server struct{}

type mongo struct {
	Operation *mgo.Collection
}

// DB is a pointer to mongo struct
var DB *mongo

func main() {
	// Host mongo server
	m, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		log.Fatalf("Could not connect to the MongoDB server: %v", err)
	}
	defer m.Close()

	DB = &mongo{m.DB("tea").C("login-api")}

	// Host grpc server
	listen, err := net.Listen("tcp", "127.0.0.1:50052")
	if err != nil {
		log.Fatalf("Could not listen on port: %v", err)
	}

	log.Println("Hosting server on", listen.Addr().String())

	s := grpc.NewServer()
	pb.RegisterSignUpServer(s, &server{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) SignUp(ctx context.Context, signUpReq *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	err := DB.Operation.Insert(signUpReq)
	if err != nil {
		return &pb.SignUpResponse{Success: false}, err
	}

	return &pb.SignUpResponse{Success: true}, nil
}
