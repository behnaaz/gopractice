package main

func Lessers(M Medable, F float32) int {
	return lessers(M, F, 0, len(M.Nums)-1)
}

func lessers(M Medable, F float32, B int, E int) int {
	if B < 0 || E < 0 {
		return 0
	}

	a := M.Nums
	if E >= len(a) {
		return len(a)
	}

	if B == E {
		t := float32(a[B])
		if F < t || (M.IsFirst && F == t) {
			return B
		}
		return B + 1
	}

	idx := (B + E) / 2
	t := float32(a[idx])
	if t > F {
		return lessers(M, F, B, idx-1)
	}

	if t < F && (idx == len(a)-1 || (float32(a[idx+1]) == F && !M.IsFirst) || float32(a[idx+1]) > F) {
		return idx + 1
	}
	if t < F {
		return lessers(M, F, idx+1, E)
	}
	if t == F {
		if M.IsFirst {
			return lessers(M, F, idx+1, E)
		}
		return lessers(M, F, B, idx-1)
	}
	return -99999999
}
