package usecase

import (
	"errors"
	"slices"
)

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

var ErrNotEnoughNecessaryBanknotes = errors.New("there are not enough necessary banknotes")

// banknotes sorted descending order
func calculate(banknotes []int, amount int) ([]CalculatingExchangeOption, error) {
	// sort
	sorted := make([]int, len(banknotes))
	copy(sorted, banknotes)
	slices.Sort(sorted)
	banknotes = sorted

	// process
	if len(banknotes) == 0 {
		return nil, nil
	}
	if amount%banknotes[0] != 0 {
		return nil, ErrNotEnoughNecessaryBanknotes
	}
	var initial []int // contains only min banknotes
	for i := 0; i < amount/banknotes[0]; i++ {
		initial = append(initial, banknotes[0])
	}
	var result []CalculatingExchangeOption
	result = append(result, initial)
	if len(banknotes) == 1 {
		return result, nil
	}
	curOpt := make([]int, len(initial))
	copy(curOpt, initial)
	for _, b := range banknotes[1:] {
		for {
			if len(curOpt) == 1 {
				break
			}
			curOpt = append([]int{b}, curOpt[:len(curOpt)-2]...)
			var curSum int
			for _, v := range curOpt {
				curSum += v
			}
			if curSum == amount {
				result = append(result, curOpt)
			} else if curSum > amount {
				break
			}
		}
	}
	return result, nil
}
