package main

import (
	"bengkel/routes"
	"fmt"
	"log"
	"os"

	"github.com/subosito/gotenv"
)


func main() {
	if err := gotenv.Load(); err != nil {
		fmt.Println(err)
		panic("Failed load gotenv")
	} // handler error
	log.Println(os.Getenv("DB_HOST"))
	r := routes.AddRoutes()

	r.Run(":"+os.Getenv("PORT"));
}
