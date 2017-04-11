package main

import (
	"fmt"
	"github.com/ometa/go_training/day_2/codingexamples/integer_translator/fizzbuzz"
)

func main() {

	t := fizzbuzz.FizzBuzzer


	for i := 1; i < 100; i++ {
		result := t.Convert(i)
		fmt.Printf("[%s] %d - %s\n", result.Converter.Name, result.In, result.Out)
	}
}