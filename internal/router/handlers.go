// handler function for POST /calcuate endpoint, receives request body in json format
package router

import (
	"encoding/json"
	"net/http"
	"sync"
	"test8989/internal/types"
	"test8989/pkg/factorial"

	"github.com/julienschmidt/httprouter"
)

func CalculateHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//gets request body from context
	body := r.Context().Value("body").([]byte)

	var req types.ABStruct

	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := types.ABStruct{}
	//f means factorial
	var af, bf int

	var wg sync.WaitGroup

	wg.Add(2)
	go factorial.Factorial(req.A, &wg, &af)
	go factorial.Factorial(req.B, &wg, &bf)
	wg.Wait()

	//checks if factorial is too large
	if af == 0 || bf == 0 {
		errorMessage := types.ResponseError{Error: "Factorial is too large"}
		jsonError, _ := json.Marshal(errorMessage)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonError)
		return
	}

	response.A = af
	response.B = bf

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)

}
