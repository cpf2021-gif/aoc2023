package main

import (
	"testing"
)

func TestQ1(t *testing.T) {
	day := NewDay8(TestMode)
	var have int64 = day.Q1()
	var want int64 = 6

	if want != have {
		t.Fatalf("Expected %d got %d", want, have)
	}
}

func TestQ2(t *testing.T) {
	day := NewDay8(TestMode)
	var have int64 = day.Q2()
	var want int64 = 6

	if want != have {
		t.Fatalf("Expected %d got %d", want, have)
	}
}
