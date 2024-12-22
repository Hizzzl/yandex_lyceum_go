package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"calc_service/internal/calculator"
)

type RequestBody struct {
	Expression string `json:"expression"`
}

type ResponseBody struct {
	Result string `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	var reqBody RequestBody
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Invalid request body"})
		return
	}

	validExpression := regexp.MustCompile(`^[0-9+\-*/()\s]+$`)
	if !validExpression.MatchString(reqBody.Expression) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Expression is not valid"})
		return
	}

	result, err := calculator.Calc(reqBody.Expression)
	if err != nil {
		if err.Error() == "division by zero" || err.Error() == "invalid operator" || err.Error() == "invalid expression" {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(ErrorResponse{Error: "Expression is not valid"})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Internal server error"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResponseBody{Result: formatResult(result)})
}

func formatResult(result float64) string {
	if result == float64(int64(result)) {
		return fmt.Sprintf("%d", int64(result))
	}
	return fmt.Sprintf("%f", result)
}

func main() {
	http.HandleFunc("/api/v1/calculate", calculateHandler)
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
