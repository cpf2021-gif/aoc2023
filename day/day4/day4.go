package main

import (
	"aoc2023/pkg/util"
	"fmt"
	"strings"
	"unicode"
)

type mode int

const (
	TestMode mode = iota
	DevMode
)

type Day4 struct {
	Data []string
	mode mode
	Card []int
}

func NewDay4(m mode) *Day4 {
	return &Day4{
		mode: m,
	}
}

func (d *Day4) Q1() int64 {
	var score int64 = 0

	if d.mode == TestMode {
		d.Data = util.ReadByLine("./data/day4_q1_test.txt")
	} else {
		d.Data = util.ReadByLine("./data/day4_input.txt")
	}

	for _, s := range d.Data {
		if d.getPoint(s) > 0 {
			score += 1 << (d.getPoint(s) - 1)
		}
	}

	return score
}

func (d *Day4) Q2() int64 {
	var score int64 = 0

	if d.mode == TestMode {
		d.Data = util.ReadByLine("./data/day4_q1_test.txt")
	} else {
		d.Data = util.ReadByLine("./data/day4_input.txt")
	}

	d.Card = make([]int, len(d.Data))
	for i := range d.Card {
		d.Card[i] = 1
	}

	for i, s := range d.Data {
		point := d.getPoint(s)
		for j := 1; j <= point; j++ {
			d.Card[i+j] += d.Card[i]
		}
	}

	for _, c := range d.Card {
		score += int64(c)
	}

	return score
}

func (d *Day4) getPoint(s string) int {
	point := 0
	numSet := make(map[string]struct{})
	numstr := strings.Split(strings.Split(s, ": ")[1], " | ")
	f := func(c rune) bool {
		return !unicode.IsDigit(c)
	}
	targetNums := strings.FieldsFunc(numstr[0], f)
	getNums := strings.FieldsFunc(numstr[1], f)

	for _, num := range targetNums {
		numSet[num] = struct{}{}
	}

	for _, num := range getNums {
		if _, ok := numSet[num]; ok {
			point++
		}
	}

	return point
}

func main() {
	day := NewDay4(DevMode)
	fmt.Printf("q1: %d\n", day.Q1())
	fmt.Printf("q2: %d\n", day.Q2())
}
