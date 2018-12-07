package main

import (
	"context"
	"crypto/tls"
	"flag"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/smtp"
	"os"
	"strings"
	"time"

	pb "github.com/UCSDIOWA/login-api/protos"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc"
)

type server struct{}

type mongo struct {
	Operation *mgo.Collection
}

// Mail struct holds info for sending reset emails
type Mail struct {
	From, Host, Port string
}

type getLoginResponseStruct struct {
	Password     string `json:"password" bson:"password"`
	FirstName    string `json:"firstname" bson:"firstname"`
	LastName     string `json:"lastname" bson:"lastname"`
	ProfileImage string `json:"profileimage" bson:"profileimage"`
}

// Possible characters for code generator
const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

var (
	// DB is a pointer to mongo struct
	DB           *mongo
	echoEndpoint = flag.String("echo_endpoint", "localhost:50052", "endpoint of login-api")
)

func main() {
	errors := make(chan error)

	go func() {
		errors <- startGRPC()
	}()

	go func() {
		flag.Parse()
		defer glog.Flush()

		errors <- startHTTP()
	}()

	for err := range errors {
		log.Fatal(err)
		return
	}
}

func startGRPC() error {
	// Host mongo server
	m, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		log.Fatalf("Could not connect to the MongoDB server: %v", err)
	}
	defer m.Close()
	log.Println("Connected to MongoDB server")

	// Accessing user collection in tea database
	DB = &mongo{m.DB("tea").C("users")}
	// Using email to find users in the database. Won't allow duplicated emails.
	DB.Operation.EnsureIndex(mgo.Index{
		Key:        []string{"email"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true})

	// Host grpc server
	listen, err := net.Listen("tcp", "127.0.0.1:50052")
	if err != nil {
		log.Fatalf("Could not listen on port: %v", err)
	}

	log.Println("Hosting server on", listen.Addr().String())

	s := grpc.NewServer()
	pb.RegisterLoginAPIServer(s, &server{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	return err
}

func startHTTP() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterLoginAPIHandlerFromEndpoint(ctx, gwmux, *echoEndpoint, opts)
	if err != nil {
		return err
	}
	log.Println("Listening on port 8080")

	mux := http.NewServeMux()
	mux.HandleFunc("/.*", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
	})
	mux.Handle("/", gwmux)
	handler := cors.Default().Handler(mux)

	herokuPort := os.Getenv("PORT")
	if herokuPort == "" {
		herokuPort = "8080"
	}

	return http.ListenAndServe(":"+herokuPort, handler)
}

// SignUp creates a new user with the input data
func (s *server) SignUp(ctx context.Context, signUpReq *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	err := DB.Operation.Insert(signUpReq)
	if err != nil {
		return &pb.SignUpResponse{Success: false}, nil
	}

	return &pb.SignUpResponse{Success: true}, nil
}

// Login verifies if the user from the login request is in the database
func (s *server) LogIn(ctx context.Context, logInReq *pb.LogInRequest) (*pb.LogInResponse, error) {
	// Fetching user from database
	var response pb.LogInResponse
	var responseStruct getLoginResponseStruct
	err := DB.Operation.Find(bson.M{"email": logInReq.Email}).One(&responseStruct)
	if err != nil {
		return &pb.LogInResponse{Success: false}, nil
	}

	// Validate user password
	if strings.Compare(responseStruct.Password, logInReq.Password) != 0 {
		return &pb.LogInResponse{Success: false}, nil
	}
	response.Success = true
	response.Firstname = responseStruct.FirstName
	response.Lastname = responseStruct.LastName
	response.Profileimage = responseStruct.ProfileImage

	return &response, nil
}

func (s *server) ForgotPassword(ctx context.Context, forgotPassReq *pb.ForgotPasswordRequest) (*pb.ForgotPasswordResponse, error) {
	// Find user with email
	user := &pb.SignUpRequest{}
	err := DB.Operation.Find(bson.M{"email": forgotPassReq.Email}).One(user)

	// Check if user exists
	if err != nil {
		return &pb.ForgotPasswordResponse{Success: false}, nil
	}

	forgotMsg := "Subject: Your Tea Account Password\n\nHello!\n\nYour tea account password is: " + user.Password + "\n\nSincerely,\nTeam Tea\n"

	sendCode(forgotPassReq.Email, forgotMsg)

	return &pb.ForgotPasswordResponse{Success: true}, nil
}

// randCode generates a random code of input length using the characters in chars var
func randCode(length int) string {
	rand.Seed(time.Now().UnixNano())
	code := make([]byte, length)
	for i := range code {
		code[i] = chars[rand.Intn(len(chars))]
	}
	return string(code)
}

// sendCode sends an email containing the generated code
func sendCode(email string, message string) error {
	info := Mail{
		"tea.noreply@gmail.com",
		"smtp.gmail.com",
		"465",
	}

	auth := smtp.PlainAuth(
		"",
		info.From,
		"cse110IOWA",
		info.Host)

	conn, err := tls.Dial(
		"tcp",
		info.Host+":"+info.Port,
		&tls.Config{
			InsecureSkipVerify: true,
			ServerName:         info.Host})

	if err != nil {
		log.Panic(err)
		return err
	}

	client, err := smtp.NewClient(conn, info.Host)
	if err != nil {
		log.Panic(err)
		return err
	}

	err = client.Auth(auth)
	if err != nil {
		log.Panic(err)
	}

	err = client.Mail(info.From)
	if err != nil {
		log.Panic(err)
		return err
	}

	err = client.Rcpt(email)
	if err != nil {
		log.Panic(err)
		return err
	}

	w, err := client.Data()
	if err != nil {
		log.Panic(err)
		return err
	}
	msg := "From: " + info.From + "\n" +
		"To: " + email + "\n" +
		message
	_, err = w.Write([]byte(msg))
	if err != nil {
		log.Panic(err)
		return err
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	client.Quit()
	return nil
}
