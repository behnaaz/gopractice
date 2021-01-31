package main

type Medable struct {
	Nums    []int
	Beg     int
	End     int
	IsFirst bool
}

func NewMedable(Ar []int, First bool) Medable {
	return Medable{Nums: Ar, IsFirst: First, Beg: 0, End: len(Ar) - 1}
}
