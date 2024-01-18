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

type {{.Day}} struct {
	Data []string
	mode mode
}


func New{{.Day}}(m mode) *{{.Day}} {
	return &{{.Day}}{
		mode: m,
	}
}

func (d *{{.Day}}) Q1() int64 {
	if d.mode == TestMode { 
		d.Data = util.ReadByLine("./data/{{.Filename}}_q1_test.txt")
	} else {
		d.Data = util.ReadByLine("./data/{{.Filename}}_input.txt")
	}

	return 0
}

func (d *{{.Day}}) Q2() int64 {
	if d.mode == TestMode {
		d.Data = util.ReadByLine("./data/{{.Filename}}_q2_test.txt")
	} else {
		d.Data = util.ReadByLine("./data/{{.Filename}}_input.txt")
	}

	return 0
}

func main() {
	day := New{{.Day}}(DevMode)
	fmt.Printf("q1: %d\n", day.Q1())
	fmt.Printf("q2: %d\n", day.Q2())
}
