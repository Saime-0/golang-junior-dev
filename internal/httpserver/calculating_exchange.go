package httpserver

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"gitlab.com/llcmediatel/recruiting/golang-junior-dev/internal/usecase"
	"io"
	"net/http"
)

func (s *Server) calculatingExchange() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			logrus.Errorf("httpserver - Server - calculatingExchange - Body.Read: %v", err)
			return
		}
		var requestBody calculatingExchangeRequest
		err = json.Unmarshal(b, &requestBody)
		if err != nil {
			logrus.Errorf("httpserver - Server - calculatingExchange - requestBody - json.Unmarshal: %v", err)
			return
		}
		logrus.Infof("httpserver - Server - calculatingExchange: requestBody=%v", requestBody)
		output, err := s.uc.CalculatingExchange.Handle(usecase.CalculatingExchangeInput{
			Amount:    requestBody.Amount,
			Banknotes: requestBody.Banknotes,
		})
		if err != nil {
			logrus.Errorf("httpserver - Server - calculatingExchange - CalculatingExchange.Handle: %v", err)
			return
		}
		responseBody := calculatingExchangeResponse{
			Exchanges: output.ExchangeOptions,
		}
		b, err = json.Marshal(responseBody)
		if err != nil {
			logrus.Errorf("httpserver - Server - calculatingExchange - responseBody - json.Marshal: %v", err)
			return
		}
		_, err = w.Write(b)
		if err != nil {
			logrus.Errorf("httpserver - Server - calculatingExchange - w.Write: %v", err)
			return
		}
	}
}

type calculatingExchangeRequest struct {
	Amount    int   `json:"amount"`
	Banknotes []int `json:"banknotes"`
}
type calculatingExchangeResponse struct {
	Exchanges [][]int `json:"exchanges"`
}
