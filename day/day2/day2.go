package main

import (
	"aoc2023/pkg/util"
	"fmt"
	"strconv"
	"strings"
)

type mode int

const (
	TestMode mode = iota
	DevMode
)

type color int

const (
	RED color = iota
	GREEN
	BLUE
)

type Day2 struct {
	Data []string
	mode mode
}

func NewDay2(mode mode) *Day2 {
	return &Day2{
		mode: mode,
	}
}

func (d *Day2) Q1() int64 {
	var score int64 = 0

	const (
		RedNum   = 12
		GreenNum = 13
		BlueNum  = 14
	)

	if d.mode == TestMode {
		d.Data = util.ReadByLine("./data/day2_q1_test.txt")
	} else {
		d.Data = util.ReadByLine("./data/day2_input.txt")
	}

	for i, line := range d.Data {
		done := true
		// Game 1: 2 green, 12 blue; 6 red, 6 blue; 8 blue, 5 green, 5 red; 5 green, 13 blue; 3 green, 7 red, 10 blue; 13 blue, 8 red
		// 1. get data
		goulp := strings.Split(line, ": ")[1]
		// 2. split by semicolon
		datas := strings.Split(goulp, "; ")
		for _, data := range datas {
			red := 0
			green := 0
			blue := 0
			colors := strings.Split(data, ", ")
			for _, c := range colors {
				num := getcolorNum(c)
				switch getColorType(c) {
				case RED:
					red += num
				case GREEN:
					green += num
				case BLUE:
					blue += num
				}
			}

			if red > RedNum || green > GreenNum || blue > BlueNum {
				done = false
			}

			if !done {
				break
			}
		}

		if done {
			score += int64(i + 1)
		}
	}

	return score
}

func (d *Day2) Q2() int64 {
	var score int64 = 0

	if d.mode == TestMode {
		d.Data = util.ReadByLine("./data/day2_q1_test.txt")
	} else {
		d.Data = util.ReadByLine("./data/day2_input.txt")
	}

	for _, line := range d.Data {
		// Game 1: 2 green, 12 blue; 6 red, 6 blue; 8 blue, 5 green, 5 red; 5 green, 13 blue; 3 green, 7 red, 10 blue; 13 blue, 8 red
		// 1. get data
		goulp := strings.Split(line, ": ")[1]
		// 2. split by semicolon
		datas := strings.Split(goulp, "; ")
		maxred := 0
		maxgreen := 0
		maxblue := 0
		for _, data := range datas {
			red := 0
			green := 0
			blue := 0
			colors := strings.Split(data, ", ")
			for _, c := range colors {
				num := getcolorNum(c)
				switch getColorType(c) {
				case RED:
					red += num
				case GREEN:
					green += num
				case BLUE:
					blue += num
				}
			}

			maxred = max(maxred, red)
			maxgreen = max(maxgreen, green)
			maxblue = max(maxblue, blue)
		}

		score += int64(maxred * maxgreen * maxblue)
	}

	return score
}

func getColorType(s string) color {
	if strings.Contains(s, "red") {
		return RED
	} else if strings.Contains(s, "green") {
		return GREEN
	} else {
		return BLUE
	}
}

func getcolorNum(s string) int {
	num, _ := strconv.Atoi(strings.Split(s, " ")[0])
	return num
}

func main() {
	day := NewDay2(DevMode)
	fmt.Printf("q1: %d\n", day.Q1())
	fmt.Printf("q2: %d\n", day.Q2())
}
