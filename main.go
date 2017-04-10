package main

import (
"fmt"
	"strconv"
)


func fizzbuzz() {
	for i := 1; i <= 100; i++ {

		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

func key_val() {
	for k, v := range []int{1,1,2,3,5,8} {
		println(k, v)
	}
}

func string_ascii_gotcha() {
	println(string(47))
}

func correct_string_conversion() {
	println(strconv.Itoa(47))
}


func array_example() {
	a := [2]int{1,2} // 2 elements
	b := [4]int{1,2} // initialize first 2 elements
	c := [...]int{1,2} // count them for me, I'm lazy
}

// a slice is an array without a set size
func slice_example() {
	a := []int{123, 456, 789}

	fmt.Printf("%v\n", a[0])
	a[1] = 555
	fmt.Printf("%v\n", a[1:])  // from element 1 to end
	fmt.Printf("%v\n", a[:1])  // from beginning to element 1, non-inclusive
	fmt.Printf("%v\n", a[1:2]) // from 1st element to the 2nd element
}



func main() {
	//fizzbuzz()
	key_val()
	string_ascii_gotcha()
}

