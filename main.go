package documentationexample

import (
	"fmt"
	"strconv"
	"os"

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

// synactic sugar: ...int is sugar for passing a slice
func variaticArgumentFizzBuzz(vals ...int) {
	fizzBuzz(vals)
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

func pointerExample1() {
	a := &[]int{1,2,3}	// create me a pointer to a slice of ints  & = create something
	b := a
	println(b)
	//b[2] = 5 // this is an error, cannot set the 3rd element of a pointer!
}

func pointerExample2() {
	a := []*int{}	// create a slice of pointers to ints
	b := a
	println(b)
}

func pointerExample3() {
	i := 1
	a := []*int{&i}
	b := a
	fmt.Println(a)
	fmt.Println(b)
}

// -----------------------------------------------------------

// this function receives a copy of the value of i, not the pointer to i
func demonstrateCopy2(i int) {
	i++
	println(i)
}

// prints:
// 2
// 1
func demonstrateCopy() {
	i := 1
	demonstrateCopy2(i)
	println(i)
}


// -----------------------------------------------------------

type foo struct {
	i int
}

func (f foo) foo() {
	f.i++
}

// this function receives a copy of the value of i, not the pointer to i
func demonstratePassByReference() {
	f := foo{3}
	f.foo()
	fmt.Println(f)
}

// -----------------------------------------------------------
// maps
//
func mapExample() {
	// var m map[string]int  	// nil!

	var m = make(map[string]int, 100) 	// use make when we know how big our map will be
	m["a"] = 1

	x := map[string]int{
		"a": 1,
	}

	println(x["b"]) 	// prints 0 since 0 is the default empty value for int.

	// since 0 is the default for int, we don't know if it's actually contained in our map, so add a
	// second paraemter which is a bool that says if the key is in the map or not.

/* ask him how to get the bool for if a value is present in the map
	y := map[string]int{
		"a": 1,
	}

	println(y["b"], present) 	// prints 0 since 0 is the default empty value for int.
*/

	// map iteration in Go is deliberately random.

/* ask him what's wrong with this.
	m := map[int]string{
		1: "a", 2: "b", 3: "c"}

	for k := range m {
		if k % 2 == 0 {
			delete(m, k)
		}
	}
*/
}

func fizzBuzzMap(vals ...int) map[int]string {
	//result := map[int]string{}			// less efficient
	result := make(map[int]string, len(vals))	// more efficient since we know the # of entries

	for _, i := range(vals) {
		switch {
		case i%3 == 0 && i%5 == 0:
			result[i] = "FizzBuzz"
		case i%3 == 0:
			result[i] = "Fizz"
		case i%5 == 0:
			result[i] = "Buzz"
		default:
			result[i] = strconv.Itoa(i)
		}
	}
	return result
}

// -----------------------------------------------------------

/*
import "errors"

var (
	TryAgain = errors.New("try again")
	SlowDown = errors.New("slow down")
)

func tryOutErrors() {
	if TryAgain {
		// do something
	}
}
*/

// -----------------------------------------------------------

// defer is like "finally"
// defer is a stack, arguments to functions that are deferred are evaluated when defer is called
func deferSomething() {

	f, err := os.Open("/etc/passwd")
	if err != nil {
		println(err)
	}
	defer f.Close()
}

// -----------------------------------------------------------

// since defer is a stack, this function will output the results in the reverse order
func deferredFizzBuzz(vals []int) {
	for _, i := range(vals) {
		switch {
		case i%3 == 0 && i%5 == 0:
			defer fmt.Println("FizzBuzz")
		case i%3 == 0:
			defer fmt.Println("Fizz")
		case i%5 == 0:
			defer fmt.Println("Buzz")
		default:
			defer fmt.Println(i)
		}
	}
}

// -----------------------------------------------------------

// a function can be a parameter
func apply(f func(int) int, i int) int {
	return f(i)
}

func double(i int) int { return i*2 }
func square(i int) int { return i*i }

func testFunctionAsParameter(x int) {
	i := apply(double, apply(square, x))
	println(i)
}

// -----------------------------------------------------------
// panic

func makePanic() {
	panic("kaboom")
}

func testPanic() {
	makePanic()
	println("foo")
}

func testPanicWithRecovery() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("recovered: %v\n", x)
		}
	}()

	panic([]int{1,2}) 	// whatever you pass to panic is what recover returns.
}

// -----------------------------------------------------------
// -----------------------------------------------------------
// -----------------------------------------------------------
// -----------------------------------------------------------
// -----------------------------------------------------------
// -----------------------------------------------------------

/*
func main() {
	//fizzbuzz()
	//print_slice_details()
	//key_val()
	//string_ascii_gotcha()
	//create_new_slice()
	//partial_fizzbuzz()
	//fizzBuzz([]int {1,2,3,4,5,6,7,8,9,10,42,55,99,33,60})
	//variaticArgumentFizzBuzz(1,2,3,4,5,6,7,8,9,10,42,55,99,33,60)
	//demonstratePassByReference()
	//fmt.Println(fizzBuzzMap(1,2,3,4,5,6,7,8,9,10,42,55,99,33,60))
	//testFunctionAsParameter(2)
}

*/


func main() {
	fmt.Println(To(199))
}