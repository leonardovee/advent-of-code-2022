package day_six

import (
	"fmt"
	"testing"
)

func TestDaySixA(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}
	for _, test := range cases {
		t.Run(fmt.Sprintf("%s first marker after character %d", test.input, test.expected), func(t *testing.T) {
			got := DaySixA(test.input)
			if got != test.expected {
				t.Errorf("got %d, want %d", got, test.expected)
			}
		})
	}
}

func TestDaySixB(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
	}
	for _, test := range cases {
		t.Run(fmt.Sprintf("%s first marker after character %d", test.input, test.expected), func(t *testing.T) {
			got := DaySixB(test.input)
			if got != test.expected {
				t.Errorf("got %d, want %d", got, test.expected)
			}
		})
	}
}
