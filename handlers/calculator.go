package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
)

type OperationRequest struct {
    A        string `json:"a"`
    B        string `json:"b"`
    Operator string `json:"operator"` // "+", "-", "*", "/"
}

type OperationResponse struct {
    Result float64 `json:"result"`
    Error  string  `json:"error,omitempty"`
}

func Calculate(w http.ResponseWriter, r *http.Request) {
    var req OperationRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    a, errA := strconv.ParseFloat(req.A, 64)
    b, errB := strconv.ParseFloat(req.B, 64)
    if errA != nil || errB != nil {
        http.Error(w, "Invalid numbers", http.StatusBadRequest)
        return
    }

    var result float64
    var errMsg string

    switch req.Operator {
    case "+":
        result = a + b
    case "-":
        result = a - b
    case "*":
        result = a * b
    case "/":
        if b == 0 {
            errMsg = "Division by zero"
        } else {
            result = a / b
        }
    default:
        errMsg = "Unsupported operator"
    }

    resp := OperationResponse{Result: result}
    if errMsg != "" {
        resp.Error = errMsg
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}
