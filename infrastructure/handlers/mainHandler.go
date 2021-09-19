package handlers

import (
	"LealTechTest/domain/entities"
	"LealTechTest/domain/usecases"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type MainHandler struct{
	// Case use programmed to the interface ( Gateway )
	bUseCases usecases.CrudUseCases
}


func (mh MainHandler ) Start(bGateway entities.BirdGateway)  {

	mh.bUseCases =  usecases.CrudUseCases{bGateway}

	http.HandleFunc("/api/v1/list", mh.listBirds)
	http.HandleFunc("/api/v1/byid", mh.getBird)
	http.HandleFunc("/api/v1/add", mh.add)
	http.HandleFunc("/api/v1/update", mh.updateBird)
	http.HandleFunc("/api/v1/delete/byid",mh.deleteBird)

}

func (h MainHandler) getBird(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		ids := r.URL.Query()["id"]
		if len(ids[0]) < 1 {
			http.Error(w, "Invalid Request Info", http.StatusBadRequest)
		}
		i, err := strconv.Atoi(ids[0])
		if err != nil {
			http.Error(w, "Invalid Request Info", http.StatusBadRequest)
		}

		bird, err2 := h.bUseCases.BGateway.FindById(i)
		if err2 != nil {
			http.Error(w, "Invalid Request Info", http.StatusNotFound)
			json.NewEncoder(w).Encode("Bird Information, Not Found")
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(bird)
		}
	} else {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
	}
}

func (h MainHandler) updateBird(w http.ResponseWriter, r *http.Request) {
	var bird entities.Bird
	if r.Method == "PATCH" {
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "Request error", http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(body, &bird)
		if err != nil {
			http.Error(w, "Request error on input data", http.StatusBadRequest)
			return
		}

		bird, err = h.bUseCases.Update(bird)
		if err != nil {
			w.Header().Set("Contet-Type", "appication/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(bird)
			return
		}

	} else {
		http.Error(w, "Request error on input data", http.StatusBadRequest)
	}

}

func (h MainHandler) deleteBird(w http.ResponseWriter, r *http.Request) {

	if r.Method == "DELETE" {

		ids := r.URL.Query()["id"]

		if len(ids[0]) < 1 {
			http.Error(w, "Invalid Request Info", http.StatusBadRequest)
			json.NewEncoder(w).Encode(string("Invalid information, to proccess request"))
		}

		i, err := strconv.Atoi(ids[0])
		if err != nil {
			http.Error(w, "Invalid Request Info", http.StatusBadRequest)
			json.NewEncoder(w).Encode(string("Invalid information, to proccess request"))
		}

		// Check if exists
		bird, err :=  h. bUseCases.FindById(i)
		if err != nil {

		}
		err =  h.bUseCases.Delete(bird)

		if err != nil {
			http.Error(w, "Invalid Request Info", http.StatusInternalServerError)
			json.NewEncoder(w).Encode(string("Error deleting bird"))
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

	} else {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
	}
}

func (h MainHandler) listBirds(w http.ResponseWriter, r *http.Request) {

	h.bUseCases.List()
	birds, err := h.bUseCases.List()

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error getting list of birds")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(birds)

}

func (h MainHandler) add(w http.ResponseWriter, r *http.Request) {

	var bird entities.Bird

	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Bad Request Info", http.StatusBadRequest)
			json.NewEncoder(w).Encode(string("Bad request"))
		}

		err = json.Unmarshal(body, &bird)
		if err != nil {
			http.Error(w, "Bad Request Info", http.StatusBadRequest)
			json.NewEncoder(w).Encode(string("Bad request"))
		}

		bird, err = h.bUseCases.Add(bird)
		if err != nil {
			http.Error(w, "Internal server Error ", http.StatusInternalServerError)
		} else {

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(bird)
		}

	} else {
		http.Error(w, "Invalid Request Method", http.StatusMethodNotAllowed)
	}
}