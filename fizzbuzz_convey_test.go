package main

// to import this library:
// $ go get github.com/smartystreets/goconvey/convey
import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"fmt"
)

func TestFizzBuzzConveyTestable(t *testing.T) {
	for input, want := range map[int]string{
		1: "1",
		2: "2",
		3: "Fizz",
		4: "4",
		5: "Buzz",
		15: "FizzBuzz",
	} {
		Convey(fmt.Sprintf("A value of %d", input), t, func() {
			Convey(fmt.Sprintf("Will return %s", want), func() {
				r := fizzBuzzTestable(input)
				So(r, ShouldEqual, want)
			})
		})
	}
}
