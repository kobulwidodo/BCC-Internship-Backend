package main

import (
	"bengkel/config"
	"bengkel/routes"

	"github.com/subosito/gotenv"
)


func main() {
	config.InitDB()
	gotenv.Load()
	r := routes.AddRoutes()

	r.Run();
}
