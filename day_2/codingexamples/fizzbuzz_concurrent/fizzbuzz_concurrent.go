package main

import (
	"strconv"
	"fmt"
)

type element struct{
	input int
	output string
}

var fizzbuzz = func(i int) string {
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

func compute(i int, c chan element) {
	c <- element{input: i, output: fizzbuzz(i)}
}


// concurrency is non-derministic (unordered)
func main() {
	iterations := 15
	c := make(chan element, 100)
	for i := 1; i <= iterations; i++ {
		go compute(i, c)
	}

	for i := 1; i <= iterations; i++ {
		elem := <- c
		fmt.Println(fmt.Sprintf("%d: %s", elem.input, elem.output))
	}
}


