package main

import (
	"log"

	"github.com/d782/tuut/bd"
	"github.com/d782/tuut/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Not connection")
		return
	}
	handlers.Handlers()
}
