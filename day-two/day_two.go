package day_two

import (
	"bufio"
	"strings"
)

const (
	Win        = 6
	Loss       = 0
	Draw       = 3
	ShouldLose = "X"
	ShouldDraw = "Y"
	ShouldWin  = "Z"
)

func DayTwoB(input string) int64 {
	var sc int64

	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		f := strings.Fields(s.Text())
		o := f[0]
		m := f[1]

		switch o {
		case "A": // rock
			switch m {
			case ShouldLose:
				// scissors
				sc += 3
				sc += 0
			case ShouldDraw:
				// rock
				sc += 1
				sc += 3
			case ShouldWin:
				// paper
				sc += 2
				sc += 6
			}
		case "B": // paper
			switch m {
			case ShouldLose:
				// rock
				sc += 1
				sc += 0
			case ShouldDraw:
				// paper
				sc += 2
				sc += 3
			case ShouldWin:
				// scissors
				sc += 3
				sc += 6
			}
		case "C": // scissors
			switch m {
			case ShouldLose:
				// paper
				sc += 2
				sc += 0
			case ShouldDraw:
				// scissors
				sc += 3
				sc += 3
			case ShouldWin:
				// rock
				sc += 1
				sc += 6
			}
		}
	}

	return sc
}

func DayTwoA(input string) int64 {
	var sc int64

	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		f := strings.Fields(s.Text())
		o := f[0]
		m := f[1]

		switch m {
		case "X":
			sc += 1
			switch o {
			case "A":
				sc += Draw
			case "B":
				sc += Loss
			case "C":
				sc += Win
			}
		case "Y":
			sc += 2
			switch o {
			case "A":
				sc += Win
			case "B":
				sc += Draw
			case "C":
				sc += Loss
			}
		case "Z":
			sc += 3
			switch o {
			case "A":
				sc += Loss
			case "B":
				sc += Win
			case "C":
				sc += Draw
			}
		}
	}

	return sc
}
