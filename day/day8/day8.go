package main

import (
	"aoc2023/pkg/util"
	"fmt"
	"strings"
)

type mode int

const (
	TestMode mode = iota
	DevMode
)

type Day8 struct {
	Data []string
	mode mode
}

type Node struct {
	Lnode string
	Rnode string
}

func NewDay8(m mode) *Day8 {
	return &Day8{
		mode: m,
	}
}

func (d *Day8) Q1() int64 {
	var result int64 = 0
	if d.mode == TestMode {
		d.Data = util.ReadByLine("./data/day8_q1_test.txt")
	} else {
		d.Data = util.ReadByLine("./data/day8_input.txt")
	}

	instructions := strings.Split(d.Data[0], "")
	mapstr := d.Data[2:]
	nodeMap := make(map[string]*Node)
	for _, s := range mapstr {
		splits := strings.Split(s, " = ")
		nodeName := splits[0]
		nodeValues := strings.Split(splits[1][1:len(splits[1])-1], ", ")
		if _, ok := nodeMap[nodeName]; !ok {
			nodeMap[nodeName] = &Node{
				Lnode: nodeValues[0],
				Rnode: nodeValues[1],
			}
		}
	}

	startNode := "AAA"
	targetNode := "ZZZ"
	for !(startNode == targetNode) {
		for _, i := range instructions {
			if i == "L" {
				startNode = nodeMap[startNode].Lnode
			} else {
				startNode = nodeMap[startNode].Rnode
			}
			result++
			if startNode == targetNode {
				break
			}
		}
	}

	return result
}

func (d *Day8) Q2() int64 {
	var result int64 = 0
	if d.mode == TestMode {
		d.Data = util.ReadByLine("./data/day8_q2_test.txt")
	} else {
		d.Data = util.ReadByLine("./data/day8_input.txt")
	}

	instructions := strings.Split(d.Data[0], "")
	mapstr := d.Data[2:]
	nodeMap := make(map[string]*Node)
	for _, s := range mapstr {
		splits := strings.Split(s, " = ")
		nodeName := splits[0]
		nodeValues := strings.Split(splits[1][1:len(splits[1])-1], ", ")
		if _, ok := nodeMap[nodeName]; !ok {
			nodeMap[nodeName] = &Node{
				Lnode: nodeValues[0],
				Rnode: nodeValues[1],
			}
		}
	}

	startNodes := make([]string, 0)
	for k := range nodeMap {
		if k[len(k)-1] == 'A' {
			startNodes = append(startNodes, k)
		}
	}

	resArr := make([]int64, 0)
	for _, startNode := range startNodes {
		var res int64 = 0
		flag := false
		for !flag {
			for _, i := range instructions {
				res++
				if i == "L" {
					startNode = nodeMap[startNode].Lnode
				} else {
					startNode = nodeMap[startNode].Rnode
				}
				if startNode[len(startNode)-1] == 'Z' {
					flag = true
					break
				}
			}
		}
		resArr = append(resArr, res)
	}

	result = resArr[0]

	for i := 1; i < len(resArr); i++ {
		result = lcm(result, resArr[i])
	}

	return result
}

func gcd(a, b int64) int64 {
	if b > a {
		a, b = b, a
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int64) int64 {
	result := a * b / gcd(a, b)
	return result
}

func main() {
	day := NewDay8(DevMode)
	fmt.Printf("q1: %d\n", day.Q1())
	fmt.Printf("q2: %d\n", day.Q2())
}
