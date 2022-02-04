package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	p, n := 0, 0

	return func() int {
		p, n = n, p+n

		if n == 0 {
			n = 1
		}
		return p
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 15; i++ {
		fmt.Println(f())
	}
}
