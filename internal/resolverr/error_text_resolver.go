package resolverr

import (
	"errors"
	"gitlab.com/llcmediatel/recruiting/golang-junior-dev/internal/usecase"
)

// Text
// tip: This may be a struct implements that interface:
//
//	type ErrorTextResolver interface {
//		Text(error) string
//	}
//	...but for simplicity, it is now a static function
func Text(err error) string {
	switch {
	case errors.Is(err, usecase.ErrNotEnoughNecessaryBanknotes):
		return "There are no necessary banknotes to calculate the exchange"
	}
	return ""
}
