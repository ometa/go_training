package fizzbuzz2_test

import (
	"testing"
	"github.com/ometa/go_training/day_2/codingexamples/fizzbuzz2"
)

func TestFizz(t *testing.T) {
	for input, want := range map[int]string {
		1: "",
		2: "",
		3: "Fizz",
	} {
		if r:= fizzbuzz2.Fizz(input); r != want {
			t.Error("Fizz(%s) did not equal expected: %s", input, want)
		}
	}
}

func TestBuzz(t *testing.T) {
	for input, want := range map[int]string {
		1: "",
		2: "",
		3: "",
		4: "",
		5: "Buzz",
	} {
		if r:= fizzbuzz2.Buzz(input); r != want {
			t.Error("Buzz(%s) did not equal expected: %s", input, want)
		}
	}
}

func TestNumber(t *testing.T) {
	for input, want := range map[int]string {
		1: "1",
		2: "2",
		3: "",
		4: "4",
		5: "",
	} {
		if r:= fizzbuzz2.Number(input); r != want {
			t.Error("Number(%s) did not equal expected: %s", input, want)
		}
	}
}