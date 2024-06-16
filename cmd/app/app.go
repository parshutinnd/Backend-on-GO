package main

import (
	"HTTP-server/internal/user"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	log.Println("router created")
	router := httprouter.New()

	log.Println("register user handler")
	hanler := user.NewHandler()
	hanler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	listener, err := net.Listen("tcp", "0.0.0.0:1234")

	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.Serve(listener))
}
