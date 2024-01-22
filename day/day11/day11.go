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

type Day11 struct {
	Data []string
	mode mode
}

type Galaxy struct {
	X int
	Y int
}

func NewDay11(m mode) *Day11 {
	return &Day11{
		mode: m,
	}
}

func (d *Day11) Q1(delta int64) int64 {
	var res int64 = 0

	if d.mode == TestMode {
		d.Data = util.ReadByLine("./data/day11_q1_test.txt")
	} else {
		d.Data = util.ReadByLine("./data/day11_input.txt")
	}

	emptyRow := make([]int, len(d.Data))
	emptyCol := make([]int, len(d.Data[0]))

	galaxies := make([]Galaxy, 0)

	for i, l := range d.Data {
		for j, c := range []rune(l) {
			if c == '#' {
				emptyRow[i] = 1
				emptyCol[j] = 1
				galaxies = append(galaxies, Galaxy{X: i, Y: j})
			}
		}
	}

	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			minX, maxX := galaxies[i].X, galaxies[j].X
			minY, maxY := galaxies[i].Y, galaxies[j].Y

			if galaxies[i].X > galaxies[j].X {
				minX, maxX = galaxies[j].X, galaxies[i].X
			}

			if galaxies[i].Y > galaxies[j].Y {
				minY, maxY = galaxies[j].Y, galaxies[i].Y
			}

			res += int64((maxX - minX) + (maxY - minY))

			for i := minX + 1; i < maxX; i++ {
				if emptyRow[i] == 0 {
					res += delta
				}
			}

			for j := minY + 1; j < maxY; j++ {
				if emptyCol[j] == 0 {
					res += delta
				}
			}
		}
	}

	return res
}

func (d *Day11) Q2() int64 {
	return d.Q1(1000000 - 1)
}

func main() {
	day := NewDay11(DevMode)
	fmt.Printf("q1: %d\n", day.Q1(1))
	fmt.Printf("q2: %d\n", day.Q2())
}
