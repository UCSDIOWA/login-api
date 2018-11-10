package main

import (
	"context"
	"log"
	"net"
  "math/rand"
  "strings"
  "time"

	pb "login-api/protos"
	"github.com/globalsign/mgo"
  "github.com/globalsign/mgo/bson"
	"google.golang.org/grpc"
)

type server struct{}

//Possible characters for code generator
const chars =  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

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

  //Accesing user collection in tea database
	DB = &mongo{m.DB("tea").C("users")}
  //Using email to find users in the database. Won't allow duplicated emails.
  DB.Operation.EnsureIndex( mgo.Index{
    Key:        []string{"email"},
    Unique:     true,
    DropDups:   true,
    Background: true,
    Sparse:     true,})

	// Host grpc server
	listen, err := net.Listen("tcp", "127.0.0.1:50052")
	if err != nil {
		log.Fatalf("Could not listen on port: %v", err)
	}

	log.Println("Hosting server on", listen.Addr().String())

	s := grpc.NewServer()
  pb.RegisterSignUpServer(s, &server{})
  pb.RegisterLogInServer(s, &server{})
  pb.RegisterForgotPasswordServer(s, &server{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}


/* Function Name: Signup
 * Description: Creates a new user with the input data. If the email already
 *              exists, it won't be added to the database and return false.
 */
func (s *server) SignUp(ctx context.Context, signUpReq *pb.SignUpRequest) (*pb.SignUpResponse, error) {
  //Set pasword to blank password not expected in signUp
  signUpReq.Password = ""
  //Generate code for email confirmation
  signUpReq.Secret.Value = RandCode(12)
  signUpReq.Secret.Sent = true
  //Inserting
  err := DB.Operation.Insert(signUpReq)
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
  //Fetching user from database
  user := &pb.SignUpRequest{}
  err := DB.Operation.Find(bson.M{"email":logInReq.Email}).One(user)
  if err != nil {
    return &pb.LogInResponse{}, err
  }

  //Validate user password
  if strings.Compare(user.Password, logInReq.Password) != 0 {
    return &pb.LogInResponse{}, nil
  }

  //Update user secret code
  user.Secret.Sent = false
  err = DB.Operation.Update(bson.M{"email":logInReq.Email}, user )
  //Error occurs if email is not found
  if err != nil {
    return &pb.LogInResponse{}, err
  }
  //Send blank password
  user.Password = ""
  return &pb.LogInResponse{ User: user}, err
}


/* Function Name: ForgotPassword
 * Description
 */
func (s *server) ForgotPassword(ctx context.Context, forgotPassReq *pb.ForgotPasswordRequest) (*pb.ForgotPasswordResponse, error) {
  //Find user with email
  user := &pb.SignUpRequest{}
  err := DB.Operation.Find(bson.M{"email":forgotPassReq.Email}).One(user)
  //Check if user exists
  if( err != nil || strings.Compare(user.Email, "") == 0 ) {
    return &pb.ForgotPasswordResponse{ Success : false }, nil
  }
  code := RandCode(12)
  user.Secret.Value = code
  user.Secret.Sent = true
  err = DB.Operation.Update(bson.M{"email":forgotPassReq.Email}, user )
  if err != nil {
    return &pb.ForgotPasswordResponse{ Success: false }, nil
  }
  return &pb.ForgotPasswordResponse{Success: true}, nil
}


/* Function Name: RandCode
 * Description: Generates a random code of input length using the characters in chars var.
 */
func RandCode( length int ) string {
  rand.Seed(time.Now().UnixNano())
  code := make([]byte, length)
  for i := range code {
    code[i] = chars[rand.Intn(len(chars))]
  }
  return string(code)
}
