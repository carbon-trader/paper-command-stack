package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/carbon-trader/paper-core/model"
	"github.com/carbon-trader/paper-core/repository"
	"gopkg.in/mgo.v2/bson"
)

//FastPaper struct
type FastPaper struct {
	ID bson.ObjectId `json:"id"`
}

var service = repository.PaperService{}

func respondWithERROR(w http.ResponseWriter, code int, msg string) {
	responWithJSON(w, code, msg)
}

func responWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func Save(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var paper model.PaperModel

	if err := json.NewDecoder(r.Body).Decode(&paper); err != nil {
		respondWithERROR(w, http.StatusBadRequest, "Invalid request payload.")
		return
	}

	idR, err := service.Save(paper)

	if err != nil {
		respondWithERROR(w, http.StatusInternalServerError, err.Error())
		return
	}

	_fast := FastPaper{
		ID: idR,
	}

	responWithJSON(w, http.StatusOK, _fast)
}
