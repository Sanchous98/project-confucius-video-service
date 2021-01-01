package main

import (
	confucius "github.com/Sanchous98/project-confucius-base"
	"github.com/gorilla/mux"
	"log"
)

func main() {
	server := confucius.NewConfuciusServer(map[string]confucius.Service{})
	router := mux.NewRouter()

	if err := server.Container.Init(); err != nil {
		log.Fatal(err)
	}

	if err := server.Container.Serve(router); err != nil {
		log.Fatal(err)
	}
}
