package main

import (
	"fmt"
	"net/http"

	"eddy-web-chat/server/pkg/websocket"
)

func serveWebServer(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Reached...")
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}
	pool.Register <- client
	client.Read()
}

func setUpRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWebServer(pool, w, r)
	})
}

func main() {
	fmt.Println("Eddy Chat Test")
	setUpRoutes()
	http.ListenAndServe(":8000", nil)
}
