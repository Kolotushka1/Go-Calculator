package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalculateHandler(t *testing.T) {
	tests := []struct {
		name         string
		method       string
		payload      CalculateRequest
		expectedCode int
		expectedBody interface{}
	}{
		{
			name:   "Успешный расчет",
			method: http.MethodPost,
			payload: CalculateRequest{
				Expression: "2+2*2",
			},
			expectedCode: http.StatusOK,
			expectedBody: CalculateResponse{Result: 6},
		},
		{
			name:   "Некорректное выражение",
			method: http.MethodPost,
			payload: CalculateRequest{
				Expression: "2+a",
			},
			expectedCode: http.StatusUnprocessableEntity,
			expectedBody: ErrorResponse{Error: "Expression is not valid"},
		},
		{
			name:   "Деление на ноль",
			method: http.MethodPost,
			payload: CalculateRequest{
				Expression: "10/0",
			},
			expectedCode: http.StatusUnprocessableEntity,
			expectedBody: ErrorResponse{Error: "Expression is not valid"},
		},
		{
			name:         "Неправильный метод",
			method:       http.MethodGet,
			payload:      CalculateRequest{},
			expectedCode: http.StatusMethodNotAllowed,
			expectedBody: `{"error": "Method not allowed"}`,
		},
		{
			name:         "Неправильный JSON",
			method:       http.MethodPost,
			payload:      CalculateRequest{},
			expectedCode: http.StatusBadRequest,
			expectedBody: ErrorResponse{Error: "Invalid JSON"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var reqBody []byte
			var err error

			if tt.name != "Неправильный JSON" {
				reqBody, err = json.Marshal(tt.payload)
				if err != nil {
					t.Fatalf("Не удалось маршалить JSON: %v", err)
				}
			} else {
				reqBody = []byte(`invalid json`)
			}

			req, err := http.NewRequest(tt.method, "/api/v1/calculate", bytes.NewBuffer(reqBody))
			if err != nil {
				t.Fatalf("Не удалось создать запрос: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(CalculateHandler)
			handler.ServeHTTP(rr, req)

			if rr.Code != tt.expectedCode {
				t.Errorf("Ожидался код %d, получен %d", tt.expectedCode, rr.Code)
			}

			if tt.expectedCode == http.StatusOK {
				var resp CalculateResponse
				err := json.Unmarshal(rr.Body.Bytes(), &resp)
				if err != nil {
					t.Errorf("Не удалось декодировать ответ: %v", err)
				}
				expected := tt.expectedBody.(CalculateResponse)
				if resp.Result != expected.Result {
					t.Errorf("Ожидался результат %v, получен %v", expected.Result, resp.Result)
				}
			} else if tt.expectedCode == http.StatusUnprocessableEntity || tt.expectedCode == http.StatusBadRequest {
				var resp ErrorResponse
				err := json.Unmarshal(rr.Body.Bytes(), &resp)
				if err != nil {
					t.Errorf("Не удалось декодировать ответ: %v", err)
				}
				expected := tt.expectedBody.(ErrorResponse)
				if resp.Error != expected.Error {
					t.Errorf("Ожидалась ошибка %q, получена %q", expected.Error, resp.Error)
				}
			}
		})
	}
}
