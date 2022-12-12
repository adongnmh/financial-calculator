package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"quoter/internal/mortgage"
	"quoter/util"
)

var router *mux.Router

func statusHandler(w http.ResponseWriter, r *http.Request) {
	util.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func initHandlers() {
	statusRouter := router.Methods(http.MethodGet).Subrouter()
	statusRouter.HandleFunc("/api/status", statusHandler)

	mortgageRouter := router.Methods(http.MethodPost).Subrouter()
	mortgageRouter.HandleFunc("/api/calculate-mortgage-payments", mortgage.CalculateMortgagePayments).Methods("POST")
	mortgageRouter.Use(mortgage.ValidatePayload)
}

func Start() {
	router = mux.NewRouter()
	initHandlers()
	fmt.Printf("router initialized and listening on 3000\n")
	log.Fatal(http.ListenAndServe(":3000", router))
}
