package main

import (
	"log"

	"github.com/edgar9694/social-network-with-go.git/bd"
	"github.com/edgar9694/social-network-with-go.git/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Println("Sin conexión a la BD")
		return
	}
	handlers.Handlers()
}
