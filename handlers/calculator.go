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
        errResp := OperationResponse{Result: 0, Error: "Invalid numbers"}
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(errResp)
        return
    }

    result, errMsg := PerformOperation(a, b, req.Operator)

    resp := OperationResponse{Result: result}
    if errMsg != "" {
        resp.Error = errMsg
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}

// PerformOperation is now testable
func PerformOperation(a, b float64, operator string) (float64, string) {
    switch operator {
    case "+":
        return a + b, ""
    case "-":
        return a - b, ""
    case "*":
        return a * b, ""
    case "/":
        if b == 0 {
            return 0, "Division by zero"
        }
        return a / b, ""
    default:
        return 0, "Unsupported operator"
    }
}
