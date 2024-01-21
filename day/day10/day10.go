package main

import (
	"aoc2023/pkg/util"
	"fmt"
	"slices"
)

type mode int

type Idx struct {
	row int
	col int
}

type Node struct {
	north bool
	south bool
	east  bool
	west  bool
}

const (
	TestMode mode = iota
	DevMode
)

type Day10 struct {
	Data []string
	mode mode
}

func NewDay10(m mode) *Day10 {
	return &Day10{
		mode: m,
	}
}

func (d *Day10) Q1() int64 {
	if d.mode == TestMode {
		d.Data = util.ReadByLine("./data/day10_q1_test.txt")
	} else {
		d.Data = util.ReadByLine("./data/day10_input.txt")
	}

	var grid [][]Node
	var start Idx
	var rowIdx int

	for _, line := range d.Data {
		lineNodes := make([]Node, len(line))
		l := []rune(line)

		for i, r := range l {
			switch r {
			case '|':
				lineNodes[i].north = true
				lineNodes[i].south = true
			case '-':
				lineNodes[i].east = true
				lineNodes[i].west = true
			case 'L':
				lineNodes[i].north = true
				lineNodes[i].east = true
			case 'J':
				lineNodes[i].north = true
				lineNodes[i].west = true
			case '7':
				lineNodes[i].south = true
				lineNodes[i].west = true
			case 'F':
				lineNodes[i].south = true
				lineNodes[i].east = true
			case 'S':
				start = Idx{row: rowIdx, col: i}
			}
		}

		grid = append(grid, lineNodes)
		rowIdx++
	}

	startShape := findStartShape(grid, start)
	grid[start.row][start.col] = startShape

	var visited []Idx
	steps, _ := traverse(grid, start, visited, 0)

	return int64((steps + 1) / 2)
}

func (d *Day10) Q2() int64 {
	if d.mode == TestMode {
		d.Data = util.ReadByLine("./data/day10_q2_test.txt")
	} else {
		d.Data = util.ReadByLine("./data/day10_input.txt")
	}

	var grid [][]Node
	var start Idx
	var rowIdx int

	for _, line := range d.Data {
		lineNodes := make([]Node, len(line))
		l := []rune(line)

		for i, r := range l {
			switch r {
			case '|':
				lineNodes[i].north = true
				lineNodes[i].south = true
			case '-':
				lineNodes[i].east = true
				lineNodes[i].west = true
			case 'L':
				lineNodes[i].north = true
				lineNodes[i].east = true
			case 'J':
				lineNodes[i].north = true
				lineNodes[i].west = true
			case '7':
				lineNodes[i].south = true
				lineNodes[i].west = true
			case 'F':
				lineNodes[i].south = true
				lineNodes[i].east = true
			case 'S':
				start = Idx{row: rowIdx, col: i}
			}
		}

		grid = append(grid, lineNodes)
		rowIdx++
	}

	startShape := findStartShape(grid, start)
	grid[start.row][start.col] = startShape

	var visited []Idx
	var insides []Idx
	_, loopPoints := traverse(grid, start, visited, 0)

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			point := Idx{row: row, col: col}
			if !slices.Contains(loopPoints, point) && isInside(grid, loopPoints, point) {
				insides = append(insides, point)
			}
		}
	}

	return int64(len(insides))
}

func main() {
	day := NewDay10(DevMode)
	fmt.Printf("q1: %d\n", day.Q1())
	fmt.Printf("q2: %d\n", day.Q2())
}

func findStartShape(pipeMap [][]Node, start Idx) Node {
	var startNode Node

	neighbourRow := start.row - 1
	if neighbourRow >= 0 && pipeMap[neighbourRow][start.col].south {
		startNode.north = true
	}

	neighbourRow = start.row + 1
	if neighbourRow < len(pipeMap) && pipeMap[neighbourRow][start.col].north {
		startNode.south = true
	}

	neighbourCol := start.col - 1
	if neighbourCol >= 0 && pipeMap[start.row][neighbourCol].east {
		startNode.west = true
	}

	neighbourCol = start.col + 1
	if neighbourCol < len(pipeMap[0]) && pipeMap[start.row][neighbourCol].west {
		startNode.east = true
	}

	return startNode
}

func traverse(pipeMap [][]Node, node Idx, visited []Idx, steps int) (int, []Idx) {
	c := Idx{row: node.row, col: node.col + 1}
	if pipeMap[node.row][node.col].east && !slices.Contains(visited, c) {
		visited = append(visited, node)
		return traverse(pipeMap, c, visited, steps+1)
	}

	c = Idx{row: node.row, col: node.col - 1}
	if pipeMap[node.row][node.col].west && !slices.Contains(visited, c) {
		visited = append(visited, node)
		return traverse(pipeMap, c, visited, steps+1)
	}

	c = Idx{row: node.row + 1, col: node.col}
	if pipeMap[node.row][node.col].south && !slices.Contains(visited, c) {
		visited = append(visited, node)
		return traverse(pipeMap, c, visited, steps+1)
	}

	c = Idx{row: node.row - 1, col: node.col}
	if pipeMap[node.row][node.col].north && !slices.Contains(visited, c) {
		visited = append(visited, node)
		return traverse(pipeMap, c, visited, steps+1)
	}

	return steps, append(visited, node)
}

func isInside(pipeMap [][]Node, loopPoints []Idx, point Idx) bool {
	var nCrossings int

	// 使用射线法判断点是否在多边形内
	for _, loopPoint := range loopPoints {
		if loopPoint.row == point.row && loopPoint.col > point.col &&
			(pipeMap[loopPoint.row][loopPoint.col].print() == '|' ||
				// 射线与多边形的边界重合
				// 1. '7' 和 'L' 一起出现时，只能算一次
				// . -------- └-┐
				pipeMap[loopPoint.row][loopPoint.col].print() == '7' ||

				//	2. 'J' 和 'F'  一起出现时，只能算一次
				// . -------- ┌-┘

				pipeMap[loopPoint.row][loopPoint.col].print() == 'F') {
			// 但 '7' 和 'F' 一起出现时，需要算两次
			// . -------- ┌-┐
			nCrossings += 1

		}
	}
	return nCrossings%2 == 1
}

func (m Node) print() rune {
	if m.north && m.south {
		return '|'
	} else if m.north && m.west {
		return 'J'
	} else if m.north && m.east {
		return 'L'
	} else if m.east && m.west {
		return '-'
	} else if m.south && m.west {
		return '7'
	} else if m.south && m.east {
		return 'F'
	}
	return 'X'
}
