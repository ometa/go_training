package notes

import "fmt"

// this is threadsafe, but also non-deterministic
// <- will block synchronously
func channels() {
	c := make(chan int, 100)
	for i := 0; i < 10; i++ {
		go compute(i, c)
	}
	for i := 0; i < 10; i++ {
		println(i, <-c)
	}
}

func compute(id int, c chan int) {

	c <- id
}


// read from multiple channels at once


// this is blocking as we are waiting for a channel to send something
func multi_channels_blocking() {
	c1 := make(chan int, 100)
	c2 := make(chan int, 100)

	select {
	case v1 := <- c1:
		fmt.Println(v1)
	case v2 := <- c2:
		fmt.Println(v2)
	}
}

func multi_channels_nonblocking() {
	c1 := make(chan int, 100)
	c2 := make(chan int, 100)

	select {
	case v1 := <- c1:
		fmt.Println(v1)
	case v2 := <- c2:
		fmt.Println(v2)
	default:
		// do nothing, don't block
	}
}





