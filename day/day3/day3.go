package main

import (
	"aoc2023/pkg/util"
	"fmt"
)

type mode int

const (
	TestMode mode = iota
	DevMode
)

type Day3 struct {
	Data    []string
	mode    mode
	starMap map[int][]int
	h       int
	w       int
}

func NewDay3(m mode) *Day3 {
	return &Day3{
		mode:    m,
		starMap: make(map[int][]int),
	}
}

func (d *Day3) Q1() int64 {
	var score int64 = 0

	if d.mode == TestMode {
		d.Data = util.ReadByLine("./data/day3_q1_test.txt")
	} else {
		d.Data = util.ReadByLine("./data/day3_input.txt")
	}

	d.h = len(d.Data)
	d.w = len(d.Data[0])

	for i, line := range d.Data {
		for j := 0; j < len(line); j++ {
			if isdigit(line[j]) {
				num := int(line[j] - '0')
				r := j + 1
				for r < len(line) && isdigit(line[r]) {
					num = num*10 + int(line[r]-'0')
					r++
				}
				if d.scan(j, r-1, i) {
					score += int64(num)
				}
				j = r
			}
		}
	}

	return score
}

func (d *Day3) Q2() int64 {
	var score int64 = 0

	if d.mode == TestMode {
		d.Data = util.ReadByLine("./data/day3_q1_test.txt")
	} else {
		d.Data = util.ReadByLine("./data/day3_input.txt")
	}

	d.h = len(d.Data)
	d.w = len(d.Data[0])

	for i, line := range d.Data {
		for j := 0; j < len(line); j++ {
			if isdigit(line[j]) {
				num := int(line[j] - '0')
				r := j + 1
				for r < len(line) && isdigit(line[r]) {
					num = num*10 + int(line[r]-'0')
					r++
				}
				d.scanStar(j, r-1, i, num)
				j = r
			}
		}
	}

	for _, v := range d.starMap {
		if len(v) == 2 {
			score += int64(v[0] * v[1])
		}
	}

	return score
}

func (d *Day3) scan(l, r, h int) bool {
	//	   xxxxxxxxxx
	// h   xl......rx
	//     xxxx....xx

	// (l-1, h-1) - (r+1, h-1)
	for i := l - 1; i <= r+1; i++ {
		if !d.outOfRange(i, h-1) {
			if d.getTargetChar(d.Data[h-1][i]) {
				return true
			}
		}
	}

	// (l-1, h), (r+1, h)
	if !d.outOfRange(l-1, h) && d.getTargetChar(d.Data[h][l-1]) {
		return true
	}
	if !d.outOfRange(r+1, h) && d.getTargetChar(d.Data[h][r+1]) {
		return true
	}

	// (l-1, h+1) - (r+1, h+1)
	for i := l - 1; i <= r+1; i++ {
		if !d.outOfRange(i, h+1) {
			if d.getTargetChar(d.Data[h+1][i]) {
				return true
			}
		}
	}

	return false
}

func (d *Day3) scanStar(l, r, h, num int) {
	//	   xxxxxxxxxx
	// h   xl......rx
	//     xxxx....xx

	// (l-1, h-1) - (r+1, h-1)
	for i := l - 1; i <= r+1; i++ {
		if !d.outOfRange(i, h-1) && d.getStarChar(d.Data[h-1][i]) {
			key := (h-1)*d.w + i
			if _, ok := d.starMap[key]; !ok {
				d.starMap[key] = []int{num}
			} else {
				d.starMap[key] = append(d.starMap[key], num)
			}
		}
	}

	// (l-1, h), (r+1, h)
	if !d.outOfRange(l-1, h) && d.getStarChar(d.Data[h][l-1]) {
		key := h*d.w + l - 1
		if _, ok := d.starMap[key]; !ok {
			d.starMap[key] = []int{num}
		} else {
			d.starMap[key] = append(d.starMap[key], num)
		}
	}
	if !d.outOfRange(r+1, h) && d.getStarChar(d.Data[h][r+1]) {
		key := h*d.w + r + 1
		if _, ok := d.starMap[key]; !ok {
			d.starMap[key] = []int{num}
		} else {
			d.starMap[key] = append(d.starMap[key], num)
		}
	}

	// (l-1, h+1) - (r+1, h+1)
	for i := l - 1; i <= r+1; i++ {
		if !d.outOfRange(i, h+1) && d.getStarChar(d.Data[h+1][i]) {
			key := (h+1)*d.w + i
			if _, ok := d.starMap[key]; !ok {
				d.starMap[key] = []int{num}
			} else {
				d.starMap[key] = append(d.starMap[key], num)
			}
		}
	}
}

func (d *Day3) getTargetChar(c byte) bool {
	return !isdigit(c) && c != '.'
}

func (d *Day3) getStarChar(c byte) bool {
	return c == '*'
}

func (d *Day3) outOfRange(x, y int) bool {
	return x < 0 || x >= d.w || y < 0 || y >= d.h
}

func isdigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func main() {
	day := NewDay3(DevMode)
	fmt.Printf("q1: %d\n", day.Q1())
	fmt.Printf("q2: %d\n", day.Q2())
}
