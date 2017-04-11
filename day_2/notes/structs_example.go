package notes


// -------------------------------------------------------------------------
// a method receiver assigns a method to a struct so you can call the method from the struct.
// method receiver gets a copy of the struct, NOT a pointer to it, so any operations
// you do on the struct inside that method will be GC'd when the method completes.

// Rule of thumb:
// If you don't know if you should be using a pointer receiver or not, use a pointer receiver.

type Value struct {
	input int
	output string
}

// method receiver (precedes the func name)
// this gets a copy of v
//func (v Value) Print() {
func (v *Value) Print() {
	println(v.input, v.output)
}

// since v is a copy, it will be GC'd after this method, so we use a pointer receiver (*Value)
//func (v Value) Output(o string) {
func (v *Value) Output(o string) {
	v.output = o
}

func testIt() {
	v := Value{1, "1"}
	v.Output("FizzBuzz")
	v.Print()
}