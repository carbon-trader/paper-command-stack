package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/carbon-trader/paper-core/repository"
	"github.com/gorilla/mux"
)

var service = repository.PaperService{}

//Const
const (
	PORT = ":3000"
)

func init() {

}

func serverUP(router *mux.Router) {
	fmt.Printf("Server Runing in port: %s", PORT)
	http.ListenAndServe(PORT, router)
}

func newRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	addRouter(router)
	return router
}

func addRouter(router *mux.Router) {
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})

	}).Methods("GET")
}

func main() {
	serverUP(newRouter())
}
