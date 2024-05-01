package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"math"
	"net/http"
	"sync"

	"github.com/julienschmidt/httprouter"
)

type ABStruct struct {
	A int `json:"a"`
	B int `json:"b"`
}

type ResponseError struct {
	Error string `json:"error"`
}

// factorial function
func factorial(n int, wg *sync.WaitGroup, result *int) {
	defer wg.Done()
	fact := 1

	for i := 2; i <= n; i++ {
		if fact > math.MaxInt32/i {
			*result = 0
			return
		}
		fact *= i
	}

	*result = fact
}

// middleware function, checks for positive numbers
func PositiveNumbersMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var bodyMap map[string]interface{}
		var req ABStruct
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
			errorMessage := ResponseError{"Incorrect input"}
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

// handler function for POST /calcuate endpoint, receives request body in json format
func CalculateHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//gets request body from context
	body := r.Context().Value("body").([]byte)

	var req ABStruct

	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := ABStruct{}
	//f means factorial
	var af, bf int

	var wg sync.WaitGroup

	wg.Add(2)
	go factorial(req.A, &wg, &af)
	go factorial(req.B, &wg, &bf)
	wg.Wait()
	//checks if factorial is too large
	if af == 0 || bf == 0 {
		errorMessage := ResponseError{"Factorial is too large"}
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

func main() {
	port := 8989
	router := httprouter.New()
	router.POST("/calculate", PositiveNumbersMiddleware(CalculateHandler))
	log.Println("Server started on port:", port)
	log.Fatal(http.ListenAndServe(":8989", router))
}
