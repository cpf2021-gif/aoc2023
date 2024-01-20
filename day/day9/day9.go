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

type Day9 struct {
	Data []string
	mode mode
}

func NewDay9(m mode) *Day9 {
	return &Day9{
		mode: m,
	}
}

func (d *Day9) Q1() int64 {
	var res int64 = 0
	if d.mode == TestMode {
		d.Data = util.ReadByLine("./data/day9_q1_test.txt")
	} else {
		d.Data = util.ReadByLine("./data/day9_input.txt")
	}

	for _, line := range d.Data {
		numStrs := strings.Split(line, " ")
		nums := make([]int64, len(numStrs))
		for i, numStr := range numStrs {
			num, _ := strconv.Atoi(numStr)
			nums[i] = int64(num)
		}

		res += nums[len(nums)-1]
		nlen := len(nums) - 1

		for nlen > 0 {
			for i := 0; i < nlen; i++ {
				nums[i] = nums[i+1] - nums[i]
			}
			res += nums[nlen-1]
			sum := int64(0)
			for i := 0; i < nlen; i++ {
				sum += nums[i]
			}

			if sum == int64(nlen)*nums[0] {
				break
			}
			nlen--
		}

	}

	return res
}

func (d *Day9) Q2() int64 {
	var res int64 = 0
	if d.mode == TestMode {
		d.Data = util.ReadByLine("./data/day9_q2_test.txt")
	} else {
		d.Data = util.ReadByLine("./data/day9_input.txt")
	}

	for _, line := range d.Data {
		numStrs := strings.Split(line, " ")
		nums := make([]int64, len(numStrs))
		for i, numStr := range numStrs {
			num, _ := strconv.Atoi(numStr)
			nums[len(numStrs)-i-1] = int64(num)
		}

		res += nums[len(nums)-1]
		nlen := len(nums) - 1

		for nlen > 0 {
			for i := 0; i < nlen; i++ {
				nums[i] = nums[i+1] - nums[i]
			}
			res += nums[nlen-1]
			flag := true
			for i := 0; i < nlen; i++ {
				if nums[i] != 0 {
					flag = false
					break
				}
			}
			if flag {
				break
			}
			nlen--
		}

	}

	return res
}

func main() {
	day := NewDay9(DevMode)
	fmt.Printf("q1: %d\n", day.Q1())
	fmt.Printf("q2: %d\n", day.Q2())
}
