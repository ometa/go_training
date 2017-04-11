package fizzbuzz2

import "testing"

// setup a dummy translator, then run a test to ensure the dummy translator can call Translate.

var dummyTranslator translator = func(i int) string { return "works" }

func TestTranslate(t *testing.T) {

	if r := dummyTranslator.Translate(5); r != "works" {
		t.Error("Translate method receiver did not work")
	}
}