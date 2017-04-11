package main

/* State machine behavior:
a
0 -> a
1 -> b

b
0 -> c
1 -> a

c
0 -> b
1 -> a
 */

type state struct {
	name string
	zero *state
	one *state
}

func (s *state) listen(zero chan struct{}, one chan struct{}, state chan *state) {
	select {
	case <-zero:
		state <- s.zero
	case <-one:
		state <- s.one
	}
}
type stateMachine struct {
	states 	[]*state
	current *state
	zero  	chan struct{}
	one   	chan struct{}
	s    	chan *state
}


func (sm *stateMachine) send(i int) {
	switch i {
	case 0:
		sm.zero <- struct{}{}
	case 1:
		sm.one <- struct{}{}
	default:
		panic("send only supports 0 or 1")

	}

	sm.current = <- sm.s
	go sm.current.listen(sm.zero, sm.one, sm.s)
}

func (sm *stateMachine) state() string {
	//close(sm.zero)
	//close(sm.one)
	//close(sm.s)
	return sm.current.name
}

func newStateMachine() *stateMachine {
	a := &state{name: "A"}
	b := &state{name: "B"}
	c := &state{name: "C"}

	a.zero = a
	a.one = b
	b.zero = c
	b.one = a
	c.zero = b
	c.one = a

	zero 		:= make(chan struct{})
	one 		:= make(chan struct{})
	statechan 	:= make(chan *state)

	go a.listen(zero, one, statechan)

	sm := &stateMachine{
		states: []*state { a, b, c },
		current: a,
		zero: zero,
		one: one,
		s: statechan,
	}
	return sm
}

func main() {
	sm := newStateMachine()
	sm.send(1)          // "state A + 1 => state B"
	//println(sm.state()) // "state B"
	sm.send(0)          // "state B + 0 => state C"
	println(sm.state()) // "state C"
}