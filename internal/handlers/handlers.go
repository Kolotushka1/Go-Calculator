package handlers

import (
	"Go-Calculator/internal/calculator"
	"encoding/json"
	"net/http"
	"strings"
)

type CalculateRequest struct {
	Expression string `json:"expression"`
}

type CalculateResponse struct {
	Result float64 `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error": "Method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var req CalculateRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(ErrorResponse{Error: "Invalid JSON"})
		if err != nil {
			return
		}
		return
	}

	if !isValidExpression(req.Expression) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		err := json.NewEncoder(w).Encode(ErrorResponse{Error: "Expression is not valid"})
		if err != nil {
			return
		}
		return
	}

	result, err := calculator.Calc(req.Expression)
	if err != nil {
		if err.Error() == "деление на ноль" || err.Error() == "некорректное выражение" || len(err.Error()) > 0 {
			w.WriteHeader(http.StatusUnprocessableEntity)
			err := json.NewEncoder(w).Encode(ErrorResponse{Error: "Expression is not valid"})
			if err != nil {
				return
			}
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(ErrorResponse{Error: "Internal server error"})
		if err != nil {
			return
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(CalculateResponse{Result: result})
}

func isValidExpression(expr string) bool {
	allowed := "0123456789+-*/(). "
	for _, ch := range expr {
		if !strings.ContainsRune(allowed, ch) {
			return false
		}
	}
	return true
}
