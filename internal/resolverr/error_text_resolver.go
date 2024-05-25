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
	case errors.Is(err, usecase.ErrInvalidInputAmount):
		return "Attempt to process an invalid amount value"
	case errors.Is(err, usecase.ErrInvalidInputBanknotes):
		return "Attempt to process an invalid banknotes value"
	case errors.Is(err, usecase.ErrReachedVariantsLimit):
		return "The data in the request suggests too many options"
	case errors.Is(err, usecase.ErrReachedComplexityLimit):
		return "The data in the query requires too many calculations"
	}
	return ""
}
