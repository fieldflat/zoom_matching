package main

import (
	"os"

	"github.com/fieldflat/zoom_matching/app/infrastructure"
)

func main() {
	db := infrastructure.NewDB()
	r := infrastructure.NewRouting(db)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}
	r.Run(":" + port)
}
