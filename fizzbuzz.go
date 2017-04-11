package documentationexample

import (
	"strconv"
)

const (
	fizz = "Fizz"
	buzz = "Buzz"
)

// fizzbuzz that takes a integer slice parameter
func fizzBuzzTestable(i int) string {
	switch {
	case i%3 == 0 && i%5 == 0:
		return "FizzBuzz"
	case i%3 == 0:
		return fizz
	case i%5 == 0:
		return buzz
	default:
		return strconv.Itoa(i)
	}
}

