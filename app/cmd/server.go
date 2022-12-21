package server

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"quoter/internal/mortgage"
	"quoter/util"
)

var chiRouter chi.Router

func statusHandler(w http.ResponseWriter, r *http.Request) {
	util.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func initHandlers() {
	chiRouter.Get("/api/status", statusHandler)
	chiRouter.Route("/api/calculate-mortgage-payments", func(r chi.Router) {
		r.Use(mortgage.ValidatePayload)
		r.Post("/", mortgage.CalculateMortgagePayments)
	})
}

func Start() {
	chiRouter = chi.NewRouter()
	initHandlers()
	fmt.Printf("router initialized and listening on 3000\n")
	log.Fatal(http.ListenAndServe(":3000", chiRouter))
}
