package day_three

import (
	"bufio"
	"strings"
	"unicode"
)

func DayThreeA(input string) int64 {
	var sp int64

	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		i := make(map[rune]bool)
		for _, r := range s.Text()[:len(s.Text())/2] {
			i[r] = true
		}
		for _, r := range s.Text()[len(s.Text())/2:] {
			if i[r] {
				sp += int64(unicode.ToLower(r) - 96)
				if unicode.IsUpper(r) {
					sp += 26
				}
				break
			}
		}
	}
	return sp
}

func DayThreeB(input string) int64 {
	var sp int64

	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		i1 := make(map[rune]bool)
		for _, r := range s.Text() {
			i1[r] = true
		}
		s.Scan()
		i2 := make(map[rune]bool)
		for _, r := range s.Text() {
			i2[r] = true
		}
		s.Scan()
		i3 := make(map[rune]bool)
		for _, r := range s.Text() {
			i3[r] = true
		}
		for r := range i1 {
			if i2[r] && i3[r] {
				sp += int64(unicode.ToLower(r) - 96)
				if unicode.IsUpper(r) {
					sp += 26
				}
				break
			}
		}
	}
	return sp
}
