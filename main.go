package main

import (
	"github.com/fieldflat/zoom_matching/app/infrastructure"
)

func main() {
	db := infrastructure.NewDB()
	r := infrastructure.NewRouting(db)
	r.Run(":8080")
}
