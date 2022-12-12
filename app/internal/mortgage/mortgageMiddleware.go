package mortgage

import (
	"context"
	"fmt"
	"net/http"
	"quoter/util"
)

type KeyRequest struct{}

func ValidatePayload(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := &PayloadMortgage{}
		err := payload.decodePayload(r.Body)
		if err != nil {
			util.RespondWithError(w, http.StatusBadRequest, "Bad Payload Mortgage")
			return
		}

		err = payload.validateFields(payload)
		if err != nil {
			util.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error validating payload: %s", err))
			return
		}

		ctx := context.WithValue(r.Context(), KeyRequest{}, payload)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
