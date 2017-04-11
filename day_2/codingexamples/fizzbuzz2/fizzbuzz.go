package fizzbuzz2

import "strconv"

const (
	fizz = "Fizz"
	buzz = "Buzz"
)

type translator func(i int) string

func (op translator) Translate(i int) string {
	return op(i)
}

var Fizz translator = func(i int) string {
	if i % 3 == 0 {
		return fizz
	}
	return ""
}

var Buzz translator = func(i int) string {
	if i % 5 == 0 {
		return buzz
	}
	return ""
}

var Number translator = func(i int) string {
	if  i % 3 == 0 || i % 5 == 0 {
		return ""
	}
	return strconv.Itoa(i)
}