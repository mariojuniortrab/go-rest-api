package main

import (
	"fmt"

	"github.com/mariojuniortrab/go-rest-api/database"
	"github.com/mariojuniortrab/go-rest-api/routes"
)

func main() {

	database.DatabaseConnect()
	fmt.Println("init rest api")
	routes.HandleRequest()
}
