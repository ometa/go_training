package main

import (
"fmt"
	"strconv"
)


func static_fizzbuzz() {
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

/*
func array_example() {
	a := [2]int{1,2} // 2 elements
	b := [4]int{1,2} // initialize first 2 elements
	c := [...]int{1,2} // count them for me, I'm lazy

	println(a)
	println(b)
	println(c)
}
*/

// a slice is an array without a set size
func slice_example() {
	a := []int{123, 456, 789}

	fmt.Printf("%v\n", a[0])
	a[1] = 555
	fmt.Printf("%v\n", a[1:])  // from element 1 to end
	fmt.Printf("%v\n", a[:1])  // from beginning to element 1, non-inclusive
	fmt.Printf("%v\n", a[1:2]) // from 1st element to the 2nd element
}


func two_dimensional_slices() {
	_ = [][]int{
		{1, 2, 3},
		{4, 5},
	}
}


func print_slice_details() {
	a := []int{1,2,3}
	println(a)
}

func create_new_slice() {
	a := []int{1,2,3}
	b := a
	b = append(b, 7)
	b[2] = 5
	fmt.Println(a)
	fmt.Println(b)
}
// use a range to specify which values to iterate over
func partial_fizzbuzz () {

	for _, i := range []int{1,2,3,4,5,6,7,10,21,42} {

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

// fizzbuzz that takes a integer slice parameter
func fizzBuzz(vals []int) {
	for _, i := range(vals) {
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

func main() {
	//fizzbuzz()
	//print_slice_details()
	//key_val()
	//string_ascii_gotcha()
	//create_new_slice()
	//partial_fizzbuzz()
	fizzBuzz([]int {1,2,3,4,5,6,7,8,9,10,42,55,99,33,60})
}

