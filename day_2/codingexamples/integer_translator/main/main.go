package main

import (
	"fmt"
	"github.com/ometa/go_training/day_2/codingexamples/integer_translator/fizzbuzz"
	"github.com/ometa/go_training/day_2/codingexamples/integer_translator/rn"
	"github.com/ometa/go_training/day_2/codingexamples/integer_translator/intconv"
)

func runConverter(c intconv.Converter) {
	for i := 1; i < 100; i++ {
		result := c.Convert(i)
		fmt.Printf("[%s] %d - %s\n", result.Converter.Name, result.In, result.Out)
	}
}
func main() {
	runConverter(fizzbuzz.FizzBuzzer)
	runConverter(rn.RNer)
}