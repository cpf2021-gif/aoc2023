package main

import (
	"aoc2023/pkg/util"
	"fmt"
	"math"
)

type mode int

const (
	TestMode mode = iota
	DevMode
)

type Day6 struct {
	Data []string
	mode mode
}

func NewDay6(m mode) *Day6 {
	return &Day6{
		mode: m,
	}
}

func (d *Day6) Q1() int64 {
	var res int64 = 1

	if d.mode == TestMode {
		d.Data = util.ReadByLine("./data/day6_q1_test.txt")
	} else {
		d.Data = util.ReadByLine("./data/day6_input.txt")
	}

	nums := []int64{}
	targets := []int64{}

	// get nums
	numsStr := d.Data[0]
	i := 0
	for i < len(numsStr) {
		if numsStr[i] >= '0' && numsStr[i] <= '9' {
			num := int64(numsStr[i] - '0')
			i++
			for i < len(numsStr) && numsStr[i] >= '0' && numsStr[i] <= '9' {
				num = num*10 + int64(numsStr[i]-'0')
				i++
			}
			nums = append(nums, num)
		} else {
			i++
		}
	}

	// get targets
	targetsStr := d.Data[1]
	i = 0
	for i < len(targetsStr) {
		if targetsStr[i] >= '0' && targetsStr[i] <= '9' {
			num := int64(targetsStr[i] - '0')
			i++
			for i < len(targetsStr) && targetsStr[i] >= '0' && targetsStr[i] <= '9' {
				num = num*10 + int64(targetsStr[i]-'0')
				i++
			}
			targets = append(targets, num)
		} else {
			i++
		}
	}

	for i, t := range targets {
		// (nums[i] - x)x - t = 0
		// -x^2 + nums[i]x - t = 0
		// a = -1, b = nums[i], c = -t
		res *= CountPointsBetweenRoots(-1, float64(nums[i]), float64(-t))
	}

	return res
}

func (d *Day6) Q2() int64 {
	if d.mode == TestMode {
		d.Data = util.ReadByLine("./data/day6_q2_test.txt")
	} else {
		d.Data = util.ReadByLine("./data/day6_input.txt")
	}

	return 0
}

func CountPointsBetweenRoots(a, b, c float64) int64 {
	delta := b*b - 4*a*c
	if delta < 0 {
		return 0
	}

	root1 := (-b + math.Sqrt(delta)) / (2 * a)
	root2 := (-b - math.Sqrt(delta)) / (2 * a)

	if root1 > root2 {
		root1, root2 = root2, root1
	}

	// no include root1 and root2
	root1 += 0.0000000001
	root2 -= 0.0000000001

	return int64(math.Floor(root2) - math.Ceil(root1) + 1)
}

func main() {
	day := NewDay6(DevMode)
	fmt.Printf("q1: %d\n", day.Q1())
	fmt.Printf("q2: %d\n", day.Q2())
}
