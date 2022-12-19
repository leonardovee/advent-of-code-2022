package day_seven

import (
	"leonardovee.com/utility"
	"testing"
)

func TestDaySevenA(t *testing.T) {
	input, _ := utility.ReadInput("./input_test")
	got := DaySevenA(input)
	if got != 95437 {
		t.Errorf("got %d, want %d", got, 95437)
	}
}

func TestDaySevenB(t *testing.T) {
	input, _ := utility.ReadInput("./input_test")
	got := DaySevenB(input)
	if got != 24933642 {
		t.Errorf("got %d, want %d", got, 24933642)
	}
}
