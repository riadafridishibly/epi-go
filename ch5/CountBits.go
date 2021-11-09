package ch5

func CountBits(x int) int {
	numBits := 0
	for x != 0 {
		numBits += x & 0x1
		x >>= 1
	}

	return numBits
}
