package usecase

import (
	"errors"
	"slices"
)

type CalculatingExchangeOutput struct {
	exchangeOptions []CalcExchangeOption
}

type CalculatingExchangeInput struct {
	Amount    int
	Banknotes []int
}
type CalcExchangeOption []int

type CalculatingExchange struct{}

func (c *CalculatingExchange) Handle(input CalculatingExchangeInput) (CalculatingExchangeOutput, error) {
	// 1. validate input if needed
	// 2. process
	out := CalculatingExchangeOutput{
		exchangeOptions: calculateExchangeOptions(input.Banknotes, input.Amount),
	}
	return out, nil
}

func calculateExchangeOptions(banknotes []int, amount int) []CalcExchangeOption {
	// sort
	sorted := make([]int, len(banknotes))
	copy(sorted, banknotes)
	slices.Sort(sorted)
	slices.Reverse(sorted)
	// process
	var exchangeOptions []CalcExchangeOption
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
func calculate(banknotes []int, amount int) ([]CalcExchangeOption, error) {
	// sort:

	sorted := make([]int, len(banknotes))
	copy(sorted, banknotes)
	slices.Sort(sorted)
	banknotes = sorted

	// process:

	if len(banknotes) == 0 {
		return nil, nil
	}
	if amount%banknotes[0] != 0 {
		return nil, ErrNotEnoughNecessaryBanknotes
	}

	var result []CalcExchangeOption
	for _, maxb := range banknotes[1:] {
		//if amount-banknotes[i] == 0 {
		//	return append(result, ), nil
		//} else if amount-banknotes[i] < 0 {
		//
		//}
		var initial []int // contains only min banknotes
		for i := 0; i < amount/maxb; i++ {
			initial = append(initial, maxb)
		}
		result = append(result, initial)
		if len(banknotes) == 1 {
			return result, nil
		}
		for _, b := range banknotes[1:] {
			curOpt := make([]int, len(initial))
			copy(curOpt, initial)
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
					continue
				}
			}
		}
	}
	return result, nil
}

// banknotes sorted descending order
func calculateV2(banknotes []int, amount int) ([]CalcExchangeOption, error) {
	// sort:

	sorted := make([]int, len(banknotes))
	copy(sorted, banknotes)
	slices.Sort(sorted)
	slices.Reverse(sorted)
	banknotes = sorted

	// upper bound

	for i, b := range banknotes {
		if b <= amount {
			banknotes = banknotes[i:]
			break
		}
	}

	// ex cases:

	if len(banknotes) == 0 {
		return nil, nil
	}

	if amount%banknotes[len(banknotes)-1] != 0 {
		println(banknotes[len(banknotes)-1])
		return nil, ErrNotEnoughNecessaryBanknotes
	}

	// process:

	var result []CalcExchangeOption
	for iii := range banknotes {
		for ii := range banknotes[iii:] {
			remAmount := amount
			var curOpt []int
			for _, b := range banknotes[ii:] { // брать начиная с initialB
				for remAmount >= b { // сколько раз поместить сумму
					curOpt = append(curOpt, b)
					remAmount -= b
				}
			}
			if remAmount == 0 {
				result = append(result, curOpt)
				tmpOpts := curOpt
				curOpt = nil
				remAmount = 0
				for i := 0; i < len(tmpOpts); i++ {
					t
				}
			}
		}

	}
	return result, nil
}

// banknotes sorted descending order
func calculateV3(banknotes []int, amount int) ([]CalcExchangeOption, error) {
	// sort:
	sorted := make([]int, len(banknotes))
	copy(sorted, banknotes)
	slices.Sort(sorted)
	banknotes = sorted

	// upper bound:
	slices.Reverse(banknotes)
	for i, b := range banknotes {
		if b <= amount {
			banknotes = banknotes[i:]
			break
		}
	}
	slices.Reverse(banknotes)

	// ex cases:
	if len(banknotes) == 0 {
		return nil, nil
	}
	if amount%banknotes[0] != 0 {
		return nil, ErrNotEnoughNecessaryBanknotes
	}

	// process:

	return result, nil
}

func processV3(banknotes []int, amount int, index int, currentOpt CalcExchangeOption) (CalcExchangeOption, bool) {

	for i := index; i < len(banknotes); i++ {
		b := banknotes[i]
		nextAmount := currentOpt.amount() + b
		if nextAmount < amount {
			currentOpt = append(currentOpt, b)
			return processV3(banknotes, amount, i, currentOpt)
		} else if nextAmount > amount {
			if i != len(banknotes)-1 {
				return processV3(banknotes, amount, i+1, currentOpt)
			} else {
				return nil, false
			}
		} else {
			return currentOpt, true
		}
	}

}

func (o CalcExchangeOption) amount() int {
	currentAmount := 0
	for _, b := range o {
		currentAmount += b
	}
	return currentAmount
}

func (o CalcExchangeOption) canAdd(amount int, banknote int) bool {
	return o.amount()+banknote <= amount
}

//canAdd := currentOpt.canAdd(amount, banknotes[index])
//if canAdd {
//	currentOpt = append(currentOpt, banknotes[index])
//}
//if currentOpt.amount() == amount {
//	return currentOpt, true
//}

//var nextIndex int
//if isLastIndex {
//	nextIndex = index
//} else {
//	nextIndex = index + 1
//}
