package usecase

import (
	"errors"
	"reflect"
	"strconv"
	"testing"
)

func TestCalculateExchangeOptions(t *testing.T) {
	var tests = []struct {
		banknotes    []int
		amount       int
		wantExchange []CalculatingExchangeOption
		wantErr      error
	}{
		{
			banknotes:    []int{50},
			amount:       401,
			wantExchange: nil,
			wantErr:      ErrNotEnoughNecessaryBanknotes,
		},
		{
			banknotes: []int{5000, 2000, 1000, 500, 200, 100, 50},
			amount:    400,
			wantErr:   nil,
			wantExchange: []CalculatingExchangeOption{
				{50, 50, 50, 50, 50, 50, 50, 50},
				{100, 50, 50, 50, 50, 50, 50},
				{100, 100, 50, 50, 50, 50},
				{100, 100, 100, 50, 50},
				{100, 100, 100, 100},
				{200, 50, 50, 50, 50},
				{200, 100, 50, 50},
				{200, 100, 100},
				{200, 200},
			},
		},
	}
	// The execution loop
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ans, err := calculate(tt.banknotes, tt.amount)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Got error:\n %v\nWant error:\n %v", ans, tt.wantExchange)
			} else if !reflect.DeepEqual(ans, tt.wantExchange) {
				t.Errorf("Got:\n %v\nWant:\n %v", ans, tt.wantExchange)
			}
		})
	}
}
