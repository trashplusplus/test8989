package router

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"test8989/internal/types"
)

// middleware function, checks for positive numbers
func PositiveNumbersMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var bodyMap map[string]interface{}
		var req types.ABStruct
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		defer r.Body.Close()

		if err := json.Unmarshal(body, &req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := json.Unmarshal(body, &bodyMap); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//checks if input is negative or doesn't exist
		if req.A < 0 || req.B < 0 || bodyMap["b"] == nil || bodyMap["a"] == nil {
			errorMessage := types.ResponseError{Error: "Incorrect input"}
			jsonError, _ := json.Marshal(errorMessage)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(jsonError)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), "body", body))

		next(w, r, ps)
	}
}
