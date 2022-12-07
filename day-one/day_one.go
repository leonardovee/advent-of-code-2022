package day_one

import (
	"bufio"
	"sort"
	"strconv"
	"strings"
)

func DayOneA(input string) int64 {
	var m []int64
	var c int64

	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		if s.Text() == "" {
			m = append(m, c)
			c = 0
			continue
		}
		v, _ := strconv.ParseInt(s.Text(), 0, 64)
		c += v
	}
	m = append(m, c)

	var b int64
	for _, v := range m {
		if v > b {
			b = v
		}
	}

	return b
}

func DayOneB(input string) int64 {
	var m []int64
	var c int64

	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		if s.Text() == "" {
			m = append(m, c)
			c = 0
			continue
		}
		v, _ := strconv.ParseInt(s.Text(), 0, 64)
		c += v
	}
	m = append(m, c)

	sort.Slice(m, func(i, j int) bool {
		return m[i] > m[j]
	})

	return m[0] + m[1] + m[2]
}
