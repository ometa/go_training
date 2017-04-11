package documentationexample

import "testing"

// MUST prefix tests with Test
func TestFizzBuzzTestable(t *testing.T) {
	for input, want := range map[int]string{
		1: "1",
		2: "2",
		3: "Fizz",
		4: "4",
		5: "Buzz",
		15: "FizzBuzz",
	} {
		if r:= fizzBuzzTestable(input); r != want {
			t.Error("fizzbuzz(%s) did not equal expected: %s", input, want)
		}
	}
}
