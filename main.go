package main

import (
	"personalidades-api/config"
	"personalidades-api/database"
	"personalidades-api/routes"
)

func main() {
	config.LoadEnvs()
	database.ConnectDb()
	routes.HandleRequests()
}
