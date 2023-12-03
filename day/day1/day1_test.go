package main

import (
	"testing"
)

func TestQ1(t *testing.T) {
	day1 := NewDay1(TestMode)
	var have int64 = day1.Q1()
	var want int64 = 142


	if want != have {
		t.Fatalf("Expected %d got %d", want, have)
	}
}

func TestQ2(t *testing.T) {
	day1 := NewDay1(TestMode)
	var have int64 = day1.Q2()
	var want int64 = 281


	if want != have {
		t.Fatalf("Expected %d got %d", want, have)
	}
}