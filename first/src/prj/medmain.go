package main

import "fmt"

func main() {
	m := []int{2, 3, 5, 7, 11, 13}
	a := NewMedable(m, true)
	l := Lessers(a, 2)
	fmt.Print(a)
	fmt.Print(l)
}
