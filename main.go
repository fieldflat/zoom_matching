package main

import (
	"./app/infrastructure"
)

func main() {
	db := infrastructure.NewDB()
	r := infrastructure.NewRouting(db)
	r.Run(":8080")
}
