package main

import (
	"context"
	"crypto/tls"
	"log"
	"math/rand"
	"net"
	"net/smtp"
	"strings"
	"time"

	pb "github.com/UCSDIOWA/login-api/protos"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"google.golang.org/grpc"
)

type server struct{}

// Mail struct holds info for sending reset emails
type Mail struct {
	From, Host, Port string
}

//Possible characters for code generator
const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

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
}

/* Function Name: Signup
 * Description: Creates a new user with the input data. If the email already
 *              exists, it won't be added to the database and return false.
 */
func (s *server) SignUp(ctx context.Context, signUpReq *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	// Set password to blank password not expected in signUp
	signUpReq.Password = ""

	// Generate code for email confirmation
	code := randCode(12)
	signUpReq.Secret.Value = code
	signUpReq.Secret.Sent = true

	// Inserting
	err := DB.Operation.Insert(signUpReq)
	if err != nil {
		return &pb.SignUpResponse{Success: false}, nil
	}

	confirmMsg := "Subject: Email Confirmation Code \n\nPleasa enter the code " +
		code + " to complete your registration.\n"
	err = sendCode(signUpReq.Email, confirmMsg)
	if err != nil {
		return &pb.SignUpResponse{Success: false}, nil
	}

	return &pb.SignUpResponse{Success: true}, nil
}

/* Function Name: Login
 * Description: Verefies if the user from the login request is in the database.
 * Returns: Blank User struct if not found. Otherwise user info except for password.
 */
func (s *server) LogIn(ctx context.Context, logInReq *pb.LogInRequest) (*pb.LogInResponse, error) {
	// Fetching user from database
	user := &pb.SignUpRequest{}
	err := DB.Operation.Find(bson.M{"email": logInReq.Email}).One(user)
	if err != nil {
		return &pb.LogInResponse{}, err
	}

	// Validate user password
	if strings.Compare(user.Password, logInReq.Password) != 0 {
		return &pb.LogInResponse{}, nil
	}

	// Update user secret code
	user.Secret.Sent = false
	err = DB.Operation.Update(bson.M{"email": logInReq.Email}, user)
	// Error occurs if email is not found
	if err != nil {
		return &pb.LogInResponse{}, err
	}
	// Send blank password
	user.Password = ""
	return &pb.LogInResponse{User: user}, err
}

/* Function Name: ForgotPassword
 * Description
 */
func (s *server) ForgotPassword(ctx context.Context, forgotPassReq *pb.ForgotPasswordRequest) (*pb.ForgotPasswordResponse, error) {
	// Find user with email
	user := &pb.SignUpRequest{}
	err := DB.Operation.Find(bson.M{"email": forgotPassReq.Email}).One(user)

	// Check if user exists
	if err != nil || strings.Compare(user.Email, "") == 0 {
		return &pb.ForgotPasswordResponse{Success: false}, err
	}

	code := randCode(12)
	user.Secret.Value = code
	user.Secret.Sent = true
	err = DB.Operation.Update(bson.M{"email": forgotPassReq.Email}, user)
	if err != nil {
		return &pb.ForgotPasswordResponse{Success: false}, err
	}

	forgotMsg := "Subject: Pasword Reset Code \n\nPlease use the code " +
		code + " to reset your account password.\n"

	sendCode(forgotPassReq.Email, forgotMsg)

	return &pb.ForgotPasswordResponse{Success: true}, nil
}

/* Function Name: randCode
 * Description: Generates a random code of input length using the characters in chars var.
 */
func randCode(length int) string {
	rand.Seed(time.Now().UnixNano())
	code := make([]byte, length)
	for i := range code {
		code[i] = chars[rand.Intn(len(chars))]
	}
	return string(code)
}

/* Function Name: sendCode
 * Description: Sends an email containing the generated code.
 */
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
	log.Println("Mail sent successfully")
	return nil
}
