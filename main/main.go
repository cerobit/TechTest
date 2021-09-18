package main

import (
	"LealTechTest/infrastructure/drivenadapters"
	"encoding/json"
	"net/http"
)


func main() {
	http.HandleFunc("/api/v1/list", listBirds)
	http.ListenAndServe(":8080", nil)

}


func  listBirds(w http.ResponseWriter, r *http.Request) {
	birds := drivenadapters.List()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(birds)

}