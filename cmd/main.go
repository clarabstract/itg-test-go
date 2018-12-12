package main

import (
	"github.com/rubyruy/itg-test/internal/app"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	app.ItgTestServer(port)
}
