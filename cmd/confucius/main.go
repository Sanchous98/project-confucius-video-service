package main

import (
	confucius "github.com/Sanchous98/project-confucius-base"
	"github.com/Sanchous98/project-confucius-video-service/services/socket"
	"github.com/gorilla/mux"
	"log"
)

var services = map[string]confucius.Service{
	"websocket": socket.NewService(&socket.Config{}),
}

func main() {
	server := confucius.NewConfuciusServer(services)
	router := mux.NewRouter()

	if err := server.Container.Init(); err != nil {
		log.Fatal(err)
	}

	if err := server.Container.Serve(router); err != nil {
		log.Fatal(err)
	}
}
