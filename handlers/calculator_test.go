package handlers

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestPerformOperation(t *testing.T) {
    tests := []struct {
        a, b     float64
        operator string
        expected float64
        errMsg   string
    }{
        {10, 5, "+", 15, ""},
        {10, 5, "-", 5, ""},
        {10, 5, "*", 50, ""},
        {10, 5, "/", 2, ""},
        {10, 0, "/", 0, "Division by zero"},
        {10, 5, "%", 0, "Unsupported operator"},
    }

    for _, tt := range tests {
        result, err := PerformOperation(tt.a, tt.b, tt.operator)
        if result != tt.expected || err != tt.errMsg {
            t.Errorf("PerformOperation(%v, %v, %q) = (%v, %q); want (%v, %q)",
                tt.a, tt.b, tt.operator, result, err, tt.expected, tt.errMsg)
        }
    }
}

func TestCalculateHandler(t *testing.T) {
    tests := []struct {
        name       string
        request    OperationRequest
        wantResult float64
        wantError  string
    }{
        {"Addition", OperationRequest{"10", "5", "+"}, 15, ""},
        {"Subtraction", OperationRequest{"10", "5", "-"}, 5, ""},
        {"Multiplication", OperationRequest{"10", "5", "*"}, 50, ""},
        {"Division", OperationRequest{"10", "5", "/"}, 2, ""},
        {"Divide by zero", OperationRequest{"10", "0", "/"}, 0, "Division by zero"},
        {"Unsupported operator", OperationRequest{"10", "5", "%"}, 0, "Unsupported operator"},
        {"Invalid number", OperationRequest{"ten", "5", "+"}, 0, "Invalid numbers"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            body, _ := json.Marshal(tt.request)
            req := httptest.NewRequest(http.MethodPost, "/calculate", bytes.NewBuffer(body))
            req.Header.Set("Content-Type", "application/json")

            rr := httptest.NewRecorder()
            Calculate(rr, req)

            var resp OperationResponse
            if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
                t.Fatalf("Failed to decode response: %v", err)
            }

            if resp.Result != tt.wantResult || resp.Error != tt.wantError {
                t.Errorf("Got result=%v error=%q; want result=%v error=%q",
                    resp.Result, resp.Error, tt.wantResult, tt.wantError)
            }
        })
    }
}
