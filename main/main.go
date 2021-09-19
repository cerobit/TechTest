package main

import (
	"LealTechTest/infrastructure/drivenadapters"
	"LealTechTest/infrastructure/handlers"
	"net/http"
)


func init() {
	// DrivenAdapter for Postgres Relational Database, im
	bGateway := birdsqladapter.Dba
	//Api Routes registering
	mainHandler := handlers.MainHandler{}
	// Initialization
	mainHandler.Start(bGateway)
}

func main() {
	http.ListenAndServe(":8080", nil)
}


