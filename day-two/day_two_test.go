package day_two

import (
	"leonardovee.com/utility"
	"testing"
)

func TestDayTwoA(t *testing.T) {
	input, _ := utility.ReadInput("./input_test")
	got := DayTwoA(input)
	if got != 15 {
		t.Errorf("Wrong score, got %d, expected %d", got, 15)
	}
}

func TestDayTwoB(t *testing.T) {
	input, _ := utility.ReadInput("./input_test")
	got := DayTwoB(input)
	if got != 12 {
		t.Errorf("Wrong score, got %d, expected %d", got, 12)
	}
}
