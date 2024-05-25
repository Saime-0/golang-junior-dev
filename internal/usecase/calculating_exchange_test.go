package usecase

import (
	"errors"
	"reflect"
	"sort"
	"strconv"
	"testing"
)

func TestCountV2Exchange(t *testing.T) {
	var tests = []struct {
		banknotes    []int
		amount       int
		wantExchange [][]int
		wantError    error
	}{
		// what to happen if you give negative values?
		{
			banknotes: []int{1, 2},
			amount:    4200000,
			wantError: ErrReachedComplexityLimit,
		},
		{
			banknotes: nil,
			amount:    1,
			wantError: ErrNotEnoughNecessaryBanknotes,
		},
		{
			banknotes: []int{2},
			amount:    3,
			wantError: ErrNotEnoughNecessaryBanknotes,
		},
		{
			banknotes: []int{5000, 2000, 1000, 500, 200, 100, 50},
			amount:    400,
			wantExchange: [][]int{
				{50, 50, 50, 50, 50, 50, 50, 50},
				{100, 50, 50, 50, 50, 50, 50},
				{100, 100, 50, 50, 50, 50},
				{100, 100, 100, 50, 50},
				{200, 50, 50, 50, 50},
				{100, 100, 100, 100},
				{200, 100, 50, 50},
				{200, 100, 100},
				{200, 200},
			},
		},
		{
			banknotes: []int{1, 2, 3, 6},
			amount:    4,
			wantExchange: [][]int{
				{1, 3},
				{2, 2},
				{1, 1, 2},
				{1, 1, 1, 1},
			},
		},
		{
			banknotes: []int{1, 2, 3, 6},
			amount:    5,
			wantExchange: [][]int{
				{2, 3},
				{1, 1, 3},
				{1, 2, 2},
				{1, 1, 1, 2},
				{1, 1, 1, 1, 1},
			},
		},
	}
	// The execution loop
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := calculateExchange(tt.banknotes, tt.amount)
			if !errors.Is(err, tt.wantError) {
				t.Errorf("\nbanknotes: %#v\namount: %#v\ngotError: %#v\nwantError: %#v\n", tt.banknotes, tt.amount, err, tt.wantError)
			}
			if !exchangesEquals(got, tt.wantExchange) {
				t.Errorf("\nbanknotes: %#v\namount: %#v\ngot: %#v\nwantExchange: %#v\n", tt.banknotes, tt.amount, got, tt.wantExchange)
			}
		})
	}
}

// https://stackoverflow.com/a/42629936

func sort2DSlice(arr [][]int) {
	if len(arr) == 0 {
		return
	}
	for i := 0; i < len(arr); i++ {
		sort.Ints(arr[i])
	}
	sort.Slice(arr[:], func(i, j int) bool {
		for x := range arr[i] {
			if arr[i][x] == arr[j][x] {
				continue
			}
			return arr[i][x] < arr[j][x]
		}
		return false
	})
}

func TestSort2DSlice(t *testing.T) {
	var tests = []struct {
		unprocessed [][]int
		wantSorted  [][]int
	}{
		{
			unprocessed: [][]int{{3}, {2}, {1}},
			wantSorted:  [][]int{{1}, {2}, {3}},
		},
		{
			unprocessed: [][]int{
				{1, 3, 2},
				{6, 5, 4},
			},
			wantSorted: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
		},
		{
			unprocessed: [][]int{
				{5},
				{3, 2},
				{0},
			},
			wantSorted: [][]int{
				{0},
				{2, 3},
				{5},
			},
		},
		{
			unprocessed: [][]int{},
			wantSorted:  [][]int{},
		},
		{
			unprocessed: nil,
			wantSorted:  [][]int{},
		},
	}
	// The execution loop
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := make([][]int, len(tt.unprocessed))
			for i, el := range tt.unprocessed {
				got[i] = make([]int, len(el))
				copy(got[i], tt.unprocessed[i])
			}
			sort2DSlice(got)
			if !reflect.DeepEqual(tt.wantSorted, got) {
				t.Errorf("\nunprocessed: %#v\ngot: %#v\nwantSorted: %#v\n", tt.unprocessed, got, tt.wantSorted)
			}
		})
	}
}

func exchangesEquals(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 {
		a = nil
	}
	if len(b) == 0 {
		b = nil
	}
	if a == nil && nil == b {
		return true
	}
	sort2DSlice(a)
	sort2DSlice(b)
	return reflect.DeepEqual(a, b)
}

func TestExchangesEquals(t *testing.T) {
	var tests = []struct {
		a          [][]int
		b          [][]int
		wantEquals bool
	}{
		{a: [][]int{{1}}, b: nil, wantEquals: false},
		{a: [][]int{{1, 2, 3}}, b: [][]int{{3, 2, 1}}, wantEquals: true},
		{a: [][]int{{3}, {2}}, b: [][]int{{2}, {3}}, wantEquals: true},
	}
	// The execution loop
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := exchangesEquals(tt.a, tt.b)
			if got != tt.wantEquals {
				t.Errorf("\na: %#v\nb: %#v\nGot: %#v\nWant: %#v\n", tt.a, tt.b, got, tt.wantEquals)
			}
		})
	}
}
