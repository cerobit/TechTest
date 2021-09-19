package main

import (
	"LealTechTest/domain/entities"
	"LealTechTest/domain/usecases"
	"LealTechTest/infrastructure/drivenadapters"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

var bGateway entities.BirdInterfce
var bUseCases usecases.CrudUseCases

func init(){
	// Specific Db Implementation
	bGateway = birdsqladapter.Dba
	// Case use programmed to the interface ( Gateway )
	bUseCases = usecases.CrudUseCases{ bGateway }
}

func main() {
	http.HandleFunc("/api/v1/list", listBirds)
	http.HandleFunc("/api/v1/byid", getBird)
	http.HandleFunc("/api/v1/add", add)
	http.HandleFunc("/api/v1/update", updateBird)
	http.HandleFunc("/api/v1/delete/byid", deleteBird)
	http.ListenAndServe(":8080", nil)

}

func getBird(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		ids := r.URL.Query()["id"]

		if len(ids[0]) < 1 {
			http.Error(w, "Invalid Request Info", http.StatusBadRequest)
		}
		i, err := strconv.Atoi(ids[0])
		if err != nil {
			http.Error(w, "Invalid Request Info", http.StatusBadRequest)
		}

		bird := bGateway.FindById(i)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(bird)

	} else {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
	}
}


func updateBird(w http.ResponseWriter, r *http.Request) {

	var bird entities.Bird

	if r.Method == "PATCH" {

	body, err :=  ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Request error", http.StatusBadRequest)
		return
	}
	os.Stdout.Write(body)
	err = json.Unmarshal(body, &bird)
	if err != nil {
		http.Error(w, "Request error on input data", http.StatusBadRequest)
		return
	}

	 err = bGateway.Update(bird)
		if  err !=nil {
			w.Header().Set("Contet-Type", "appication/json")
			w.WriteHeader(http.StatusOK)
			return
		}

	}else {
		http.Error(w, "Request error on input data", http.StatusBadRequest)
	}

}

func deleteBird(w http.ResponseWriter, r *http.Request) {

	if r.Method == "DELETE" {

		ids := r.URL.Query()["id"]

		if len(ids[0]) < 1 {
			http.Error(w, "Invalid Request Info", http.StatusBadRequest)
		}

		i, err := strconv.Atoi(ids[0])
		if err != nil {
			http.Error(w, "Invalid Request Info", http.StatusBadRequest)
		}

		err = bGateway.DeleteById(i)
		if err != nil {
			http.Error(w, "Invalid Request Info", http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

	} else {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
	}
}

func listBirds(w http.ResponseWriter, r *http.Request) {

	bUseCases.List()
	birds, err := bUseCases.List()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error getting list of birds")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(birds)

}

func add(w http.ResponseWriter, r *http.Request) {

	var bird entities.Bird

	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Invalid Request Method", http.StatusBadRequest)
		}

		err = json.Unmarshal(body, &bird)
		if err != nil {
			http.Error(w, "Invalid Request Method", http.StatusBadRequest)
		}

		err = bGateway.Add(bird)
		if err != nil {
			http.Error(w, "Invalid Request Method", http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(bird)

	} else {
		http.Error(w, "Invalid Request Method", http.StatusMethodNotAllowed)
	}

}
