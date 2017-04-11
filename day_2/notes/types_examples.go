package notes

type name string
type username string

func typesExamples() {
	n := name("Devin")
	u := username("dbreen")
	println(n)
	println(u)

	// we cannot do direct equality comparison between two types
	/*
	if u == n {				// mismatched types
		println("equal")
	}
	*/
}

// -----------------------------------------------------
// calc can only accept functions matching the binaryOp pattern

type binaryOp func(int, int) int

func add(i, j int) int 	{ return i+j }
func sub(i, j int) int 	{ return i-j }
func mul(i, j int) int 	{ return i*j }
func div(i, j int) int 	{ return i/j }

func calc(op binaryOp, i, j int) int {
	return op(i, j)
}

// -----------------------------------------------------
// we can limit the types of functions that can be call by consumers:

type foo string

const Bar = foo("Bar")
const Baz = foo("Baz")

func Do(f foo) {}