package usecase

import "testing"

func TestCalculateOptionForRemainingBanknotes(t *testing.T) {
	var tests = []struct {
		name         string
		remBanknotes []int
		remAmount    int
		want         []int
	}{
		// the table itself
		{"9 should be Foo", 9, "Foo"},
		{"3 should be Foo", 3, "Foo"},
		{"1 is not Foo", 1, "1"},
		{"0 should be Foo", 0, "Foo"},
	}

	calculateOptionForRemainingBanknotes()
	result := Fooer(3)
	if result != "Foo" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "Foo")
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := calculateOptionForRemainingBanknotes(tt.input)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
