package main

import (
	"EFB-User-Backend/routes"
	"EFB-User-Backend/database"
	"log"
)

func main() { 
	database.ConnectDatabase()
	r := routes.SetupRouter()
	log.Println("🚀 Server running at http://localhost:8080")
	r.Run(":8080")
}
