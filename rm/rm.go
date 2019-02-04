package rm

import "math/rand"

const DefaultTrials = 6

func bits(n int64) []int64 {
	var bits []int64
	for ; n > 0; n = n >> 1 {
		bits = append(bits, n%2)
	}
	return bits
}

func mrCompositeWitness(a, n int64) bool {
	theBits := bits(n - 1)
	var rem int64 = 1
	for i := len(theBits) - 1; i >= 0; i-- {
		x := rem
		rem = (rem * rem) % n

		if rem == 1 && x != 1 && x != n-1 {
			return true
		}

		if theBits[i] == 1 {
			rem = (rem * a) % n
		}

	}
	if rem != 1 {
		return true
	}
	return false
}

func Prime(n, trials int64) bool {
	if n < 2 {
		return false
	}

	var i int64
	for ; i < trials; i++ {
		r := rand.Intn(int(n)-1) + 1
		if mrCompositeWitness(int64(r), n) {
			return false
		}
	}
	return true
}
