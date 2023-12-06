package main

import (
	"testing"
)

func TestQ1(t *testing.T) {
	day1 := NewDay2(TestMode)
	var have int64 = day1.Q1()
	var want int64 = 8

	if want != have {
		t.Fatalf("Expected %d got %d", want, have)
	}
}

func TestQ2(t *testing.T) {
	day1 := NewDay2(TestMode)
	var have int64 = day1.Q2()
	var want int64 = 2286

	if want != have {
		t.Fatalf("Expected %d got %d", want, have)
	}
}
