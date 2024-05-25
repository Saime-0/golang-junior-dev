package usecase

import (
	"errors"
	"github.com/sirupsen/logrus"
)

type CalculatingExchangeOutput struct {
	ExchangeOptions [][]int
}

type CalculatingExchangeInput struct {
	Amount    int
	Banknotes []int
}

type CalculatingExchange struct{}

func (c *CalculatingExchange) Handle(input CalculatingExchangeInput) (CalculatingExchangeOutput, error) {
	// 1. validate:
	err := validateCalculatingExchangeInput(input)
	if err != nil {
		return CalculatingExchangeOutput{}, err
	}

	// 2. process:
	exchange, err := calculateExchange(input.Banknotes, input.Amount)
	if err != nil {
		return CalculatingExchangeOutput{}, err
	}
	return CalculatingExchangeOutput{
		ExchangeOptions: exchange,
	}, nil
}

var (
	ErrInvalidInputAmount    = errors.New("ErrInvalidInputAmount")
	ErrInvalidInputBanknotes = errors.New("ErrInvalidInputBanknotes")
)

func validateCalculatingExchangeInput(inp CalculatingExchangeInput) error {
	if inp.Amount < 0 {
		return ErrInvalidInputAmount
	}
	for _, banknote := range inp.Banknotes {
		if banknote <= 0 {
			return ErrInvalidInputBanknotes
		}
	}
	return nil
}

const (
	_ExchangeVariantsLimit   = 15000
	_ExchangeComplexityLimit = 50_000_000
)

var (
	ErrNotEnoughNecessaryBanknotes = errors.New("ErrNotEnoughNecessaryBanknotes")
	ErrReachedVariantsLimit        = errors.New("ErrReachedVariantsLimit")
	ErrReachedComplexityLimit      = errors.New("ErrReachedComplexityLimit")
)

// thx to ChatGPT (based on https://www.geeksforgeeks.org/coin-change-dp-7/) for help:

func calculateExchange(banknotes []int, amount int) ([][]int, error) {
	if len(banknotes) == 0 {
		if amount > 0 {
			return nil, ErrNotEnoughNecessaryBanknotes
		}
		return nil, nil
	} else if amount == 0 {
		return nil, nil
	}
	variants := 0
	complexity := 0
	var fn func(banknotes []int, n int, amount int) ([][]int, error)
	fn = func(banknotes []int, n, amount int) ([][]int, error) {
		complexity++
		if complexity > _ExchangeComplexityLimit {
			return nil, ErrReachedComplexityLimit
		}
		if amount == 0 {
			variants++
			if variants > _ExchangeVariantsLimit {
				return nil, ErrReachedVariantsLimit
			}
			return [][]int{{}}, nil // Базовый случай: пустой вариант
		}
		if amount < 0 || n <= 0 {
			return nil, nil // Невозможный случай: возвращаем пустой список вариантов
		}

		// Рассматриваем два случая:
		// 1. Включаем монету banknotes[n-1] в вариант
		variants1, err := fn(banknotes, n, amount-banknotes[n-1])
		if err != nil {
			return nil, err
		}
		for i := range variants1 {
			variants1[i] = append(variants1[i], banknotes[n-1]) // Добавляем монету в вариант
		}

		// 2. Не включаем монету banknotes[n-1] в вариант
		variants2, err := fn(banknotes, n-1, amount)
		if err != nil {
			return nil, err
		}

		// Объединяем варианты из обоих случаев
		return append(variants1, variants2...), nil
	}
	exchanges, err := fn(banknotes, len(banknotes), amount)
	logrus.Debugf("complexity: %v", complexity)
	logrus.Debugf("variants: %v", variants)
	if err != nil {
		return nil, err
	}
	if len(banknotes) > 0 && len(exchanges) == 0 {
		return nil, ErrNotEnoughNecessaryBanknotes
	}
	return exchanges, nil
}
