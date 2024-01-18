package main

import (
	"aoc2023/pkg/util"
	"cmp"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type mode int

const (
	TestMode mode = iota
	DevMode
)

type Day7 struct {
	Data []string
	mode mode
}

type Card struct {
	Number string
	Score  int64
	Type   CardType
}

var numberMap = map[rune]int{
	'2': 1,
	'3': 2,
	'4': 3,
	'5': 4,
	'6': 5,
	'7': 6,
	'8': 7,
	'9': 8,
	'T': 9,
	'J': 10,
	'Q': 11,
	'K': 12,
	'A': 13,
}

var numberMap2 = map[rune]int{
	'J': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'Q': 11,
	'K': 12,
	'A': 13,
}

func NewCard1(number string) *Card {
	data := strings.Split(number, " ")

	num, _ := strconv.Atoi(data[1])

	return &Card{
		Number: data[0],
		Score:  int64(num),
		Type:   GetCardType(data[0]),
	}
}

func NewCard2(number string) *Card {
	data := strings.Split(number, " ")

	num, _ := strconv.Atoi(data[1])

	return &Card{
		Number: data[0],
		Score:  int64(num),
		Type:   GetCardType2(data[0]),
	}
}

func GetCardType(Number string) CardType {
	s := strings.Split(Number, "")
	sort.Strings(s)
	Number = strings.Join(s, "")

	sortMap := make(map[rune]int64)
	for _, v := range Number {
		sortMap[v]++
	}

	if len(sortMap) == 5 {
		return HighCard
	} else if len(sortMap) == 4 {
		// 2 1 1 1
		return OnePair
	} else if len(sortMap) == 3 {
		// 3 1 1
		for _, v := range sortMap {
			if v == 3 {
				return ThreeOfAKind
			}
		}
		// 2 2 1
		return TwoPair
	} else if len(sortMap) == 2 {
		// 4 1
		for _, v := range sortMap {
			if v == 4 {
				return FourOfAKind
			}
		}
		// 3 2
		return FullHouse
	} else {
		return FiveOfAKind
	}
}

func GetCardType2(Number string) CardType {
	s := strings.Split(Number, "")
	sort.Strings(s)
	Number = strings.Join(s, "")

	sortMap := make(map[rune]int64)
	for _, v := range Number {
		sortMap[v]++
	}

	// J can be any number
	if len(sortMap) == 5 {
		if sortMap['J'] == 1 {
			return OnePair
		} else {
			return HighCard
		}
	} else if len(sortMap) == 4 {
		switch sortMap['J'] {
		case 0:
			return OnePair
		default:
			return ThreeOfAKind
		}
	} else if len(sortMap) == 3 {
		// 3 1 1
		switch sortMap['J'] {
		case 0:
			// 3 1 1
			for _, v := range sortMap {
				if v == 3 {
					return ThreeOfAKind
				}
			}
			// 2 2 1
			return TwoPair
		case 1:
			// 4 1
			for _, v := range sortMap {
				if v == 3 {
					return FourOfAKind
				}
			}
			// 3 2
			return FullHouse
		default:
			return FourOfAKind
		}
	} else if len(sortMap) == 2 {
		if _, ok := sortMap['J']; ok {
			return FiveOfAKind
		} else {
			// 4 1
			for _, v := range sortMap {
				if v == 4 {
					return FourOfAKind
				}
			}
			// 3 2
			return FullHouse
		}
	} else {
		return FiveOfAKind
	}
}

type CardType int

const (
	HighCard CardType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func NewDay7(m mode) *Day7 {
	return &Day7{
		mode: m,
	}
}

func (d *Day7) Q1() int64 {
	var score int64 = 0

	if d.mode == TestMode {
		d.Data = util.ReadByLine("./data/day7_q1_test.txt")
	} else {
		d.Data = util.ReadByLine("./data/day7_input.txt")
	}

	Cards := make([]*Card, 0)
	for _, v := range d.Data {
		Cards = append(Cards, NewCard1(v))
	}

	var cmpFunc = func(a, b *Card) int {
		if n := cmp.Compare(a.Type, b.Type); n != 0 {
			return n
		} else {
			for i := 0; i < len(a.Number); i++ {
				if n := cmp.Compare(numberMap[rune(a.Number[i])], numberMap[rune(b.Number[i])]); n != 0 {
					return n
				}
			}
		}
		return 0
	}

	slices.SortFunc(Cards, cmpFunc)

	for i, v := range Cards {
		score += v.Score * int64(i+1)
	}

	return score
}

func (d *Day7) Q2() int64 {
	var score int64 = 0

	if d.mode == TestMode {
		d.Data = util.ReadByLine("./data/day7_q2_test.txt")
	} else {
		d.Data = util.ReadByLine("./data/day7_input.txt")
	}

	Cards := make([]*Card, 0)
	for _, v := range d.Data {
		Cards = append(Cards, NewCard2(v))
	}

	var cmpFunc = func(a, b *Card) int {
		if n := cmp.Compare(a.Type, b.Type); n != 0 {
			return n
		} else {
			for i := 0; i < len(a.Number); i++ {
				if n := cmp.Compare(numberMap2[rune(a.Number[i])], numberMap2[rune(b.Number[i])]); n != 0 {
					return n
				}
			}
		}
		return 0
	}

	slices.SortFunc(Cards, cmpFunc)

	for i, v := range Cards {
		score += v.Score * int64(i+1)
	}

	return score
}

func main() {
	day := NewDay7(DevMode)
	fmt.Printf("q1: %d\n", day.Q1())
	fmt.Printf("q2: %d\n", day.Q2())
}
