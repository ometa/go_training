package notes

type name string
type username string

func typesExamples() {

	n := name("Devin")
	u := username("dbreen")
/*
	if username(u) == n {
		println("equal")
	}
*/

}



type binaryOp func(int, int) int

func add(i, j int) int 	{ return i+j }
func sub(i, j int) int 	{ return i-j }
func mul(i, j int) int 	{ return i*j }
func div(i, j int) int 	{ return i/j }

func calc(op binaryOp, i, j int) int {
	return op(i, j)
}
