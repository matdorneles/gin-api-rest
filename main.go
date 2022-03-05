package main

import (
	"github.com/matdorneles/gin-api-rest/database"
	"github.com/matdorneles/gin-api-rest/routes"
)

//pkmg
func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
