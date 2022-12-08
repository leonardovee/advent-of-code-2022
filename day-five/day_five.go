package day_five

import (
	"bufio"
	"fmt"
	"strings"
)

type stack struct {
	elem []rune
}

func (s *stack) fill(r rune) {
	s.elem = append([]rune{r}, s.elem...)
}

func (s *stack) push(r rune) {
	s.elem = append(s.elem, r)
}

func (s *stack) pushMany(r []rune) {
	s.elem = append(s.elem, r...)
}

func (s *stack) pop() rune {
	r := s.elem[len(s.elem)-1]
	s.elem = s.elem[:len(s.elem)-1]
	return r
}

func (s *stack) popMany(n int) []rune {
	r := s.elem[len(s.elem)-n : len(s.elem)]
	s.elem = s.elem[:len(s.elem)-n]
	return r
}

func DayFiveA(input string) string {
	//sl := make([]stack, 9)
	sl := make([]stack, 3)

	s := bufio.NewScanner(strings.NewReader(input))
	s.Scan()
	for s.Text() != " 1   2   3" {
		//for s.Text() != " 1   2   3   4   5   6   7   8   9 " {
		for i, v := range s.Text() {
			if v != ' ' && v != '[' && v != ']' {
				sl[i/4].fill(v)
			}
		}
		s.Scan()
	}
	s.Scan()

	for s.Scan() {
		var box, from, to int
		_, err := fmt.Sscanf(s.Text(), "move %d from %d to %d", &box, &from, &to)
		if err != nil {
			return ""
		}
		for m := 0; m < box; m++ {
			sl[to-1].push(sl[from-1].pop())
		}
	}

	var r string
	for _, s := range sl {
		r = fmt.Sprintf("%s%s", r, string(s.pop()))
	}
	return r
}

func DayFiveB(input string) string {
	//sl := make([]stack, 9)
	sl := make([]stack, 3)

	s := bufio.NewScanner(strings.NewReader(input))
	s.Scan()
	for s.Text() != " 1   2   3" {
		//for s.Text() != " 1   2   3   4   5   6   7   8   9" {
		for i, v := range s.Text() {
			if v != ' ' && v != '[' && v != ']' {
				sl[i/4].fill(v)
			}
		}
		s.Scan()
	}
	s.Scan()

	for s.Scan() {
		var box, from, to int
		_, err := fmt.Sscanf(s.Text(), "move %d from %d to %d", &box, &from, &to)
		if err != nil {
			return ""
		}
		sl[to-1].pushMany(sl[from-1].popMany(box))
	}

	var r string
	for _, s := range sl {
		r = fmt.Sprintf("%s%s", r, string(s.popMany(1)))
	}
	return r
}
