package httpserver

import "net/http"

func (s *Server) calculatingExchange() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

	}
}
