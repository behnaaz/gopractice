package main

import "fmt"

func main() {
	m := []int{2, 3, 5, 7, 11, 13}
	a := NewMedable(m, true)
	l := Lessers(a, 3.3)
	fmt.Print("hi %d", l)
}
