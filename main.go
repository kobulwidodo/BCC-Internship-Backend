package main

import (
	"bengkel/config"
	"bengkel/routes"
	"log"
	"os"

	"github.com/subosito/gotenv"
)


func main() {
	gotenv.Load()
	config.InitDB()
	log.Println(os.Getenv("DB_HOST"))
	r := routes.AddRoutes()

	r.Run(":"+os.Getenv("PORT"));
}
