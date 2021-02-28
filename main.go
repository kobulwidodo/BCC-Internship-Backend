package main

import (
	"bengkel/config"
	"bengkel/routes"
)


func main() {
	config.InitDB();
	r := routes.AddRoutes();

	r.Run();
}
