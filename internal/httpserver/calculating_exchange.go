package httpserver

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"gitlab.com/llcmediatel/recruiting/golang-junior-dev/internal/usecase"
	"io"
	"net/http"
)

func (s *Server) calculatingExchange() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			err := fmt.Errorf("httpserver - Server - calculatingExchange - Body.Read: %w", err)
			return _StatusError{error: err, HTTPStatus: http.StatusInternalServerError}
		}
		var requestBody calculatingExchangeRequest
		err = json.Unmarshal(b, &requestBody)
		if err != nil {
			err := fmt.Errorf("httpserver - Server - calculatingExchange - requestBody - json.Unmarshal: %w", err)
			return _StatusError{error: err, HTTPStatus: http.StatusBadRequest}
		}
		logrus.Debugf("httpserver - Server - calculatingExchange: requestBody=%v", requestBody)
		output, err := s.uc.CalculatingExchange.Handle(usecase.CalculatingExchangeInput{
			Amount:    requestBody.Amount,
			Banknotes: requestBody.Banknotes,
		})
		if err != nil {
			err := fmt.Errorf("httpserver - Server - calculatingExchange - CalculatingExchange.Handle: %w", err)
			return _StatusError{error: err, HTTPStatus: http.StatusUnprocessableEntity}
		}
		responseBody := calculatingExchangeResponse{
			Exchanges: output.ExchangeOptions,
		}
		b, err = json.Marshal(responseBody)
		if err != nil {
			err := fmt.Errorf("httpserver - Server - calculatingExchange - responseBody - json.Marshal: %w", err)
			return _StatusError{error: err, HTTPStatus: http.StatusUnprocessableEntity}
		}
		_, err = w.Write(b)
		if err != nil {
			err := fmt.Errorf("httpserver - Server - calculatingExchange - w.Write: %w", err)
			return _StatusError{error: err, HTTPStatus: http.StatusInternalServerError}
		}
		return nil
	}
}

type calculatingExchangeRequest struct {
	Amount    int   `json:"amount"`
	Banknotes []int `json:"banknotes"`
}
type calculatingExchangeResponse struct {
	Exchanges [][]int `json:"exchanges"`
}
