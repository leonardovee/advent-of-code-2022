package day_one

import (
	"leonardovee.com/utility"
	"testing"
)

func TestDayOneA(t *testing.T) {
	input, _ := utility.ReadInput("./input_test")
	got := DayOneA(input)
	if got != 24000 {
		t.Errorf("Wrong amount of calories, got %d, expected %d", got, 24000)
	}
}

func TestDayOneB(t *testing.T) {
	input, _ := utility.ReadInput("./input_test")
	got := DayOneB(input)
	if got != 45000 {
		t.Errorf("Wrong amount of calories, got %d, expected %d", got, 45000)
	}
}
