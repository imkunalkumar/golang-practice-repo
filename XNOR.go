// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	fmt.Println(xnor(false, false))
	fmt.Println(xnor(false, true))
	fmt.Println(xnor(true, false))
	fmt.Println(xnor(true, true))
}

func xnor(a, b bool) bool {
	p1 := and(a, b)
	p2 := and(!a, !b)

	return or(p1, p2)
}

func and(a, b bool) bool {
	if a == b {
		return a
	}
	return false
}
func or(a, b bool) bool {
	if a == b {
		return a
	}
	return true
}
