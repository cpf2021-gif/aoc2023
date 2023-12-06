package main

import (
	"aoc2023/pkg/util"
	"fmt"
	"regexp"
)

var scoreMap = map[string]int{
	"one":   1,
	"1":     1,
	"two":   2,
	"2":     2,
	"three": 3,
	"3":     3,
	"four":  4,
	"4":     4,
	"five":  5,
	"5":     5,
	"six":   6,
	"6":     6,
	"seven": 7,
	"7":     7,
	"eight": 8,
	"8":     8,
	"nine":  9,
	"9":     9,
}

type mode int

const (
	TestMode mode = iota
	DevMode
)

type Day1 struct {
	Data []string
	mode mode
}

func findFirstMatch(s string, pattern string) (string, bool) {
	r := regexp.MustCompile(pattern)
	match := r.FindStringSubmatch(s)
	if len(match) > 0 {
		return match[0], true
	}
	return "", false
}

func findLastMatch(s string, pattern string) (string, bool) {
	r := regexp.MustCompile(pattern)
	matches := r.FindAllStringSubmatch(s, -1)

	if len(matches) > 0 {
		lastMatch := matches[len(matches)-1]
		return lastMatch[0], true
	}
	return "", false
}

func NewDay1(m mode) *Day1 {
	return &Day1{
		mode: m,
	}
}

func (d *Day1) Q1() int64 {
	if d.mode == TestMode {
		d.Data = util.ReadByLine("./data/day1_q1_test.txt")
	} else {
		d.Data = util.ReadByLine("./data/day1_input.txt")
	}

	var res int64 = 0

	for _, v := range d.Data {
		var lnum int = 0
		var rnum int = 0

		str, _ := findFirstMatch(v, "1|2|3|4|5|6|7|8|9")
		lnum = scoreMap[str]

		str, _ = findLastMatch(v, "1|2|3|4|5|6|7|8|9")
		rnum = scoreMap[str]

		res += int64(lnum*10 + rnum)
	}

	return res
}

func (d *Day1) Q2() int64 {
	if d.mode == TestMode {
		d.Data = util.ReadByLine("./data/day1_q2_test.txt")
	} else {
		d.Data = util.ReadByLine("./data/day1_input.txt")
	}

	var res int64 = 0

	for _, v := range d.Data {
		var lnum int = 0
		var rnum int = 0

		str, _ := findFirstMatch(v, "1|2|3|4|5|6|7|8|9|one|two|three|four|five|six|seven|eight|nine")
		lnum = scoreMap[str]

		str, _ = findFirstMatch(util.ReverseSrting(v), "1|2|3|4|5|6|7|8|9|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin")
		rnum = scoreMap[util.ReverseSrting(str)]

		res += int64(lnum*10 + rnum)
	}

	return res
}

func main() {
	day := NewDay1(DevMode)
	fmt.Printf("q1: %d\n", day.Q1())
	fmt.Printf("q2: %d\n", day.Q2())
}
