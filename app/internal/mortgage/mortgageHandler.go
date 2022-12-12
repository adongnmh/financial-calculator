package mortgage

import (
	"net/http"
	"quoter/util"
)

// CalculateMortgagePayments will calculate the mortgage payments with CMHC insurance
func CalculateMortgagePayments(w http.ResponseWriter, r *http.Request) {
	payload := r.Context().Value(KeyRequest{}).(*PayloadMortgage)
	p := CalculateMortgage(payload)
	util.RespondWithJSON(w, http.StatusOK, map[string]float64{"mortgage-payments": p})
}
