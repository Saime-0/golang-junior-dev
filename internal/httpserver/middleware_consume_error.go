package httpserver

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"gitlab.com/llcmediatel/recruiting/golang-junior-dev/internal/resolverr"
	"net/http"
)

// thx for idea https://habr.com/ru/articles/811361/
func (s *Server) consumeError(next HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		err := next(w, r)
		if err == nil {
			return nil
		}
		// pull status from specific error
		var status int
		var statusErr _StatusError
		if errors.As(err, &statusErr) {
			//if status != 0 && errors.As(err, &statusErr) {
			status = statusErr.HTTPStatus
		}
		// print error
		if status >= 500 && status <= 599 {
			logrus.Errorf("consumeError: %v", err)
		} else {
			logrus.Infof("consumeError: %v", err)
		}
		// pull error text from specific error
		text := resolverr.Text(err)
		if text == "" {
			text = err.Error() // or set from err
		}
		http.Error(w, fmt.Sprintf("{\"error\":\"%v\"}", text), status)
		return err
	}
}

type _StatusError struct {
	error
	HTTPStatus int
}

func (e _StatusError) Unwrap() error { return e.error }
