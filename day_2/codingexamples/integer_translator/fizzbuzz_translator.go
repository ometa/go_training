package integer_translator

import (
	"strconv"
	"fmt"
)

const (
	cFizz = "Fizz"
	cBuzz = "Buzz"
)

type FizzBuzzTranslator struct{}

// TODO: return false if there's an error
func (translator FizzBuzzTranslator) Translate(i int) (output string, input int, conversionType string, err error) {
	output = fmt.Sprintf("%s%s%s",
		fizz.translate(i),
		buzz.translate(i),
		number.translate(i),
	)

	return output, i, "FizzBuzzTranslator", nil
}

type fizzBuzzTranslator func(i int) string

func (op fizzBuzzTranslator) translate(i int) string {
	return op(i)
}

var fizz fizzBuzzTranslator = func(i int) string {
	if i % 3 == 0 {
		return cFizz
	}
	return ""
}

var buzz fizzBuzzTranslator = func(i int) string {
	if i % 5 == 0 {
		return cBuzz
	}
	return ""
}

var number fizzBuzzTranslator = func(i int) string {
	if  i % 3 == 0 || i % 5 == 0 {
		return ""
	}
	return strconv.Itoa(i)
}