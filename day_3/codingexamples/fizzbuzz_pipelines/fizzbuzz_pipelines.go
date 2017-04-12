package main

import (
	"fmt"
	"strconv"
)

// define an "interface" for our converters
type converter func(in <-chan fizzbuzz, out chan<- fizzbuzz)

type fizzbuzz struct {
	in  int
	out string
}

// define a slice of converters
var converters = []converter{fizz, buzz, number}

// define the converter functions
func fizz(in <-chan fizzbuzz, out chan<- fizzbuzz) {
	for i := range in {
		if i.in%3 == 0 {
			i.out = "Fizz"
		}
		out <- i
	}
}

func buzz(in <-chan fizzbuzz, out chan<- fizzbuzz) {
	for i := range in {
		if i.in%5 == 0 {
			i.out += "Buzz"
		}
		out <- i
	}
}

func number(in <-chan fizzbuzz, out chan<- fizzbuzz) {
	for i := range in {
		if i.in%3 != 0 && i.in%5 != 0 {
			i.out = strconv.Itoa(i.in)
		}
		out <- i
	}
}

func main() {
	// create a slice of channels that accept a fizzbuzz object
	chans := make([]chan fizzbuzz, 4)

	// instantiate the channels, and close them when the work is done
	for i := range chans {
		chans[i] = make(chan fizzbuzz)
		defer close(chans[i])			// close when function is done.
	}

	// boot up the channels
	for i := 0; i < len(chans)-1; i++ {
		go converters[i](chans[i], chans[i+1])
	}

	// iterate over our input set
	for i := 1; i < 100; i++ {
		chans[0] <- fizzbuzz{in: i}
		r := <-chans[len(chans)-1]
		fmt.Printf("%d - %s\n", r.in, r.out)
	}
}