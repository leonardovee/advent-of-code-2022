package day_one

import (
	"os"
	"testing"
)

func TestDayOneA(t *testing.T) {
	input, err := os.ReadFile("./input_test")
	if err != nil {
		t.Error("Error while reading the file")
	}
	got := DayOneA(string(input))
	if got != 24000 {
		t.Errorf("Wrong amount of calories, got %d, expected %d", got, 24000)
	}
}

func TestDayOneB(t *testing.T) {
	input, err := os.ReadFile("./input_test")
	if err != nil {
		t.Error("Error while reading the file")
	}
	got := DayOneB(string(input))
	if got != 45000 {
		t.Errorf("Wrong amount of calories, got %d, expected %d", got, 45000)
	}
}
