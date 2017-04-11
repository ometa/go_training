package fizzbuzz1

import "strconv"

const (
	fizz = "Fizz"
	buzz = "Buzz"
)

type fizzBuzzOp func(i int) string

func FizzBuzz(op fizzBuzzOp, i int) string {
	return op(i)
}

func Fizz(i int) string {
	if i % 3 == 0 {
		return fizz
	}
	return ""
}

func Buzz(i int) string {
	if i % 5 == 0 {
		return buzz
	}
	return ""
}

func Number(i int) string {
	if  i % 3 == 0 || i % 5 == 0 {
		return ""
	}
	return strconv.Itoa(i)
}