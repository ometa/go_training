package main

import (
	"fmt"
	"github.com/ometa/go_training/day_2/codingexamples/fizzbuzz2"
)

func main() {
	for i := 1; i < 100; i++ {
		fmt.Printf("%d - %s%s%s\n", i,
			fizzbuzz2.Fizz.Translate(i),
			fizzbuzz2.Buzz.Translate(i),
			fizzbuzz2.Number.Translate(i),
		)
	}
}