# eddy
unfinished chat app using Golang, React, Docker

## Notes
Used to gain familiarity with Golang 

https://youtu.be/YS4e4q9oBaU

https://golang.org/doc/effective_go.html

## Issues 
Error: Firefox can’t establish a connection to the server at ws://localhost:8080/ws

Having trouble establishing socket connection despite using code straight out the documentation:

https://godoc.org/github.com/gorilla/websocket


## Improvements
- Websocket connections
- Rendered chat.messages to client 
- Deploy via Docker

# How to use locally

#### Cient 
- cd server/client && yarn install && yarn start

#### Server 
- in /server/ $ go run main.go 
