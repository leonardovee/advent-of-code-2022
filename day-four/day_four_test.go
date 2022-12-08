package day_four

import (
	"leonardovee.com/utility"
	"testing"
)

func TestDayFourA(t *testing.T) {
	input, _ := utility.ReadInput("./input_test")
	got := DayFourA(input)
	if got != 2 {
		t.Errorf("Wrong count, got %d, expected %d", got, 2)
	}
}

func TestDayFourB(t *testing.T) {
	input, _ := utility.ReadInput("./input_test")
	got := DayFourB(input)
	if got != 4 {
		t.Errorf("Wrong count, got %d, expected %d", got, 4)
	}
}
