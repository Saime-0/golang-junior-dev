package usecase

import (
	"errors"
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
	// 1. validate unprocessed if needed
	// 2. process
	ex, err := calculateExchange(input.Banknotes, input.Amount)
	if err != nil {
		return CalculatingExchangeOutput{}, err
	}
	out := CalculatingExchangeOutput{
		ExchangeOptions: ex,
	}
	return out, nil
}

var ErrNotEnoughNecessaryBanknotes = errors.New("ErrNotEnoughNecessaryBanknotes")

// thx to ChatGPT (based on https://www.geeksforgeeks.org/coin-change-dp-7/) for help:

func calculateExchange(banknotes []int, amount int) ([][]int, error) {
	var fn func(banknotes []int, n int, amount int) ([][]int, error)
	fn = func(banknotes []int, n, amount int) ([][]int, error) {
		if amount == 0 {
			return [][]int{{}}, nil // Базовый случай: пустой вариант
		}
		if amount < 0 {
			return nil, ErrNotEnoughNecessaryBanknotes // Невозможный случай: возвращаем пустой список вариантов
		}
		if n <= 0 {
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
	return fn(banknotes, len(banknotes), amount)
}
