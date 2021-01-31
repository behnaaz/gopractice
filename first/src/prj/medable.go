package main

type Medable struct {
	Nums    []int
	Beg     int
	End     int
	IsFirst bool
}

type Median struct {
	Val  float32
	Less int
}

func NewMedable(Ar []int, First bool) Medable {
	return Medable{Nums: Ar, IsFirst: First, Beg: 0, End: len(Ar) - 1}
}

func midex(m Medable) int {
	return (m.Beg + m.End) / 2
}

func ShiftBeg(m Medable) {
	m.Beg = midex(m) - 1
}

func ShiftEnd(m Medable) {
	m.End = midex(m) + 1
}

func IsDone(m Medable) bool {
	return (m.Beg > m.End)
}

func CalcMedian(m Medable) Median {
	a := m.Nums
	v := float32(a[midex(m)])
	if len(a)%2 == 0 {
		v = float32((a[midex(m)] + a[midex(m)+1]) / 2)
	}
	return Median{Val: v, Less: midex(m)}
}
