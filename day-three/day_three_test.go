package day_three

import (
	"leonardovee.com/utility"
	"testing"
)

func TestDayThreeA(t *testing.T) {
	input, _ := utility.ReadInput("./input_test")
	got := DayThreeA(input)
	if got != 157 {
		t.Errorf("Wrong score, got %d, expected %d", got, 157)
	}
}

func TestDayThreeB(t *testing.T) {
	input, _ := utility.ReadInput("./input_test")
	got := DayThreeB(input)
	if got != 70 {
		t.Errorf("Wrong score, got %d, expected %d", got, 70)
	}
}
