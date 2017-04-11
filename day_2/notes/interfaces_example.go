package notes

import "io"

// If you have an interface with only a single method, the idiomatic pattern is
// to add "er" to the end of the interface name.  Read(er)
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type ReaderWriter interface{
	Read()
	Write()
}

// Type assertion

func foo2(v interface{}) {
	i := v.(int) // type assert that v is an int. during runtime, if not, panic.
	println(i)

	// to handle a mismatched type assertion:
	if i, ok := v.(int); ok {
		println(i)
	}

	// this doesn't scale if we're handling different types. use a "type switch"

	switch v := v.(type) {
	case int:
		println("int", v)
	case string:
		println("string", v)
	case io.Reader:
		println("io.Reader")
	default:
		println("unknown")
	}
}
