package usecase

import "slices"

type CalculatingExchangeOutput struct {
	exchangeOptions []CalculatingExchangeOption
}

type CalculatingExchangeInput struct {
	Amount    int
	Banknotes []int
}
type CalculatingExchangeOption []int

type CalculatingExchange struct{}

func (c *CalculatingExchange) Handle(input CalculatingExchangeInput) (CalculatingExchangeOutput, error) {
	// 1. validate input if needed
	// 2. process
	out := CalculatingExchangeOutput{
		exchangeOptions: calculateExchangeOptions(input.Banknotes, input.Amount),
	}
	return out, nil
}

func calculateExchangeOptions(banknotes []int, amount int) []CalculatingExchangeOption {
	// sort
	sorted := make([]int, len(banknotes))
	copy(sorted, banknotes)
	slices.Sort(sorted)
	slices.Reverse(sorted)
	// process
	var exchangeOptions []CalculatingExchangeOption
	for i, banknote := range sorted {
		if banknote <= amount {
			remainingBanknotes := sorted[i:]
			option := calculateOptionForRemainingBanknotes(remainingBanknotes, amount)
			exchangeOptions = append(exchangeOptions, option)
		}
	}
	return exchangeOptions
}

func calculateOptionForRemainingBanknotes(remBanknotes []int, remAmount int) []int {
	var option []int
	i := 0
	for i < len(remBanknotes) {
		var b = remBanknotes[i]
		if remAmount < b { // if space not available then take next note
			i += 1
			continue
		} else { // else sub B from lostAmount and add B to option
			option = append(option, b)
			remAmount -= b
		}
	}
	return option
}
