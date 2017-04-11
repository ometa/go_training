package documentationexample

import (
	"strconv"
	"fmt"
)

const (
	cFizz = "Fizz"
	cBuzz = "Buzz"
	cFizzBuzz = "FizzBuzz"
)

func isFizzbuzz(i int) 	bool 	{ return i%3 == 0 && i%5 == 0 }
func isFizz(i int) 	bool	{ return i%3 == 0 }
func isBuzz(i int) 	bool 	{ return i%5 == 0 }
//func isNumber(i int) 	bool 	{ fmt.Println(strconv.Itoa(i)) }

// fizzbuzz1 that takes a integer slice parameter
func fizzBuzzClosure(i int) string {
	switch {
	case isFizzbuzz(i):
		return cFizzBuzz
	case isFizz(i):
		return cFizz
	case isBuzz(i):
		return cBuzz
	default:
		return strconv.Itoa(i)
	}
}

func fizzBuzzClosureMain() {
	for _, i := range []int{1,2,3,4,5,6,7,8,9,10,15} {
		fmt.Println(fmt.Sprintf("%d: %s", i, fizzBuzzClosure(i)))
	}
}
