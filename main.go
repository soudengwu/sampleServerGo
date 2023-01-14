package main

import (
	"log"
	"net/http"

	"github.com/soudengwu/sampleServerGo/server"
)

func main() {
	server := &server.PlayerServer{server.NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":5000", server))
}
