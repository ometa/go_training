package intconv

type Converter struct {
	F func(int) string
	Name string
}

type Result struct {
	In int
	Out string
	Converter
}

func (c Converter) Convert(i int) Result {
	r := Result{In: i, Converter: c	}
	r.Out = c.F(i)
	return r
}