package day_four

import (
	"bufio"
	"fmt"
	"strings"
)

func DayFourA(input string) int64 {
	var co int64

	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		var a, b, c, d int
		_, err := fmt.Sscanf(s.Text(), "%d-%d,%d-%d", &a, &b, &c, &d)
		if err != nil {
			return 0
		}
		if c >= a && d <= b || a >= c && b <= d {
			co++
		}
	}

	return co
}

func DayFourB(input string) int64 {
	var co int64

	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		var a, b, c, d int
		_, err := fmt.Sscanf(s.Text(), "%d-%d,%d-%d", &a, &b, &c, &d)
		if err != nil {
			return 0
		}
		if c <= b && d >= a || a <= d && b >= c {
			co++
		}
	}

	return co
}
