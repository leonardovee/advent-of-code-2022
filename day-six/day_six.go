package day_six

import (
	"bufio"
	"strings"
)

func DaySixA(input string) int {
	s := bufio.NewScanner(strings.NewReader(input))
	var res int
	for s.Scan() {
		var last4 []int32
		for pos, char := range s.Text() {
			if len(last4) == 4 {
				if isUnique(last4) {
					res = pos
					break
				}
				last4 = last4[1:]
			}
			last4 = append(last4, char)
		}
	}
	return res
}

func DaySixB(input string) int {
	s := bufio.NewScanner(strings.NewReader(input))
	var res int
	for s.Scan() {
		var last14 []int32
		for pos, char := range s.Text() {
			if len(last14) == 14 {
				if isUnique(last14) {
					res = pos
					break
				}
				last14 = last14[1:]
			}
			last14 = append(last14, char)
		}
	}
	return res
}

func isUnique(input []int32) bool {
	buffer := make(map[int32]bool)
	var result []int32
	for _, val := range input {
		if _, ok := buffer[val]; !ok {
			buffer[val] = true
			result = append(result, val)
		}
	}
	return len(result) == len(input)
}
