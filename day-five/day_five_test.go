package day_five

import (
	"leonardovee.com/utility"
	"testing"
)

func TestDayFiveA(t *testing.T) {
	input, _ := utility.ReadInput("./input_test")
	got := DayFiveA(input)
	if got != "CMZ" {
		t.Errorf("Wrong count, got %s, expected %s", got, "CMZ")
	}
}

func TestDayFiveB(t *testing.T) {
	input, _ := utility.ReadInput("./input_test")
	got := DayFiveB(input)
	if got != "MCD" {
		t.Errorf("Wrong count, got %s, expected %s", got, "MCD")
	}
}
