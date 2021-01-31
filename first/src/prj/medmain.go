package main

import "fmt"

func main() {
	m := []int{2, 3, 5, 7, 11, 13, 13, 13, 13, 13, 13, 17}
	a := NewMedable(m, false)
	l := Lessers(a, 13.4)
	fmt.Print(a)
	fmt.Print(l)
}
