package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/carbon-trader/paper-core/model"
	"github.com/carbon-trader/paper-core/repository"
	"github.com/gorilla/mux"
)

var service = repository.PaperService{}

func respondWithERROR(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, msg)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

//Save information in database
func Save(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	//
	var paper model.Paper

	//
	if err := json.NewDecoder(r.Body).Decode(&paper); err != nil {
		respondWithERROR(w, http.StatusBadRequest, "Invalid request payload.")
		return
	}

	//
	idR, err := service.Save(paper)

	//
	if err != nil {
		respondWithERROR(w, http.StatusInternalServerError, err.Error())
		return
	}

	//
	respondWithJSON(w, http.StatusOK, map[string]string{"id": idR.Hex()})
}

//Delete information in database
func Delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	//Get the params in request
	params := mux.Vars(r)

	//
	if err := service.Delete(params["id"]); err != nil {
		respondWithERROR(w, http.StatusInternalServerError, err.Error())
		return
	}

	//
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

//Update information in database
func Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	//initilize a paper as PaperModel
	var paper model.Paper

	// returns an error if the body was different of model.Papermodel
	if err := json.NewDecoder(r.Body).Decode(&paper); err != nil {
		respondWithERROR(w, http.StatusBadRequest, "Invalid resquest payload.")
		return
	}

	// returns an error if receive a bad signal of database
	if err := service.Update(paper.ID.Hex(), paper); err != nil {
		respondWithERROR(w, http.StatusInternalServerError, err.Error())
		return
	}

	//
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
