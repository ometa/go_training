package fizzbuzz

import (
	"strconv"
	"github.com/ometa/go_training/day_2/codingexamples/integer_translator/intconv"
)

var FizzBuzzer = intconv.Converter{
	F: f,
	Name: "FizzBuzzer",
}

var f = func(i int) string {
	switch {
	case i % 3 == 0 && i % 5 == 0:
		return "FizzBuzz"
	case i % 3 == 0:
		return "Fizz"
	case i % 5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(i)
	}
}