# login-api

## Overview ##
This repository contains the necessary files to host restful API's using Protocol Buffers (a.k.a protobuf) under golang to run a database. Information on protocol buffers
can be found on [protobufs Google Developers site](https://developers.google.com/protocol-buffers/docs/proto3).
All of the endpoints are hosted using [Heroku](https://www.heroku.com). The database was implemented using [MongoDB](https://mongodb.com)
with the help of the public MongoDB driver [mgo](https://github.com/globalsign/mgo) and is being hosted using [mLab](https://mlab.com).
This repository handles requests from the login page of our website.

## Program Execution ##
Make sure [mgo](https://github.com/globalsign/mgo), [glog](https://github.com/golang/glog), [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway), 
[cors](https://github.com/rs/cors), and [grpc](https://godoc.org/google.golang.org/grpc) are installed in your golang environemnt. To execute the program 
run the server.go file as follows,

	go run server.go

This will execute the server file.

## Endpoints ##
Each enpoint expect to receive specific filds to process a request. The following are the expectations for each endpoint and the resopnse

| Endpoint | Request | Response |
|:--------:|---------|----------|
| SignUp   | string email = 1;<br>string password = 2;<br>string firstname = 3;<br>string lastname = 4;<br>string profileimage = 5; | bool success; |
| LogIn    | string email = 1;<br>string password = 2;| bool success = 1;<br>string firstname =2;<br>string lastname = 3;<br>string profileimage = 4; |
| ForgotPassword | string email; | bool success; |
