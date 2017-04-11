package main

import (
	"fmt"
	"github.com/ometa/go_training/day_2/codingexamples/integer_translator"
)

func main() {

	t := integer_translator.FizzBuzzTranslator{}

	for i := 1; i < 100; i++ {
		if output, input, converterType, err := t.Translate(i); err == nil {
			fmt.Printf("[%s] %d - %s\n", converterType, input, output)
		} else {
			fmt.Printf("Error translating: %s", err)
		}

	}
}