package day_seven

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type node struct {
	name   string
	size   int
	childs map[string]*node
	father *node
	file   bool
}

func DaySevenA(input string) int {
	s := bufio.NewScanner(strings.NewReader(input))
	var cd *node
	var dirs []*node
	for s.Scan() {
		cl := strings.Fields(s.Text())
		if cl[1] == "cd" {
			switch cl[2] {
			case "/":
				cd = &node{"/", 0, make(map[string]*node), nil, false}
			case "..":
				cd = cd.father
			default:
				cd = cd.childs[cl[2]]
			}
		} else if cl[0] == "dir" {
			cd.childs[cl[1]] = &node{cl[1], 0, make(map[string]*node), cd, false}
			dirs = append(dirs, cd.childs[cl[1]])
		} else if cl[0] != "$" {
			s, _ := strconv.ParseInt(cl[0], 10, 0)
			cd.childs[cl[1]] = &node{cl[1], int(s), nil, cd, true}
		}
	}
	var ts int
	for _, d := range dirs {
		s := getSize(*d)
		fmt.Printf("dir: %s size: %d\n", d.name, s)
		if s <= 100000 {
			ts += s
		}
	}
	return ts
}

func DaySevenB(input string) int {
	s := bufio.NewScanner(strings.NewReader(input))
	var cd *node
	var dirs []*node
	for s.Scan() {
		cl := strings.Fields(s.Text())
		if cl[1] == "cd" {
			switch cl[2] {
			case "/":
				cd = &node{"/", 0, make(map[string]*node), nil, false}
				dirs = append(dirs, cd)
			case "..":
				cd = cd.father
			default:
				cd = cd.childs[cl[2]]
			}
		} else if cl[0] == "dir" {
			cd.childs[cl[1]] = &node{cl[1], 0, make(map[string]*node), cd, false}
			dirs = append(dirs, cd.childs[cl[1]])
		} else if cl[0] != "$" {
			s, _ := strconv.ParseInt(cl[0], 10, 0)
			cd.childs[cl[1]] = &node{cl[1], int(s), nil, cd, true}
		}
	}
	se := getSize(*dirs[0])
	tf := 30000000 - (70000000 - se)
	for _, d := range dirs {
		si := getSize(*d)
		if si > tf && si-tf < se-tf {
			se = si
		}
	}
	return se
}

func getSize(r node) int {
	if r.file {
		return r.size
	}
	s := 0
	for _, c := range r.childs {
		s += getSize(*c)
	}
	return s
}
