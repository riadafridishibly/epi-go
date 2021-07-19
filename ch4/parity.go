package ch4

// ParityNaive explanation
//  XOR:
// 	 x   y   res
//   0   0    0
//   1   0    1
//   0   1    1
//   1   1    0
//
//  Iter 1:
//   x = 0b1101
//   res = 0 ^ 1 = 1
//   x = 0b110
//
//  Iter 2:
//   x = 0b110
//   res = 1 ^ 0 = 1
//   x = 0b11
//
//  Iter 3:
//   x = 0b11
//   res = 1 ^ 1 = 0
//   x = 0b1
//
//  Iter 4:
//   x = 0b1
//   res = 0 ^ 1 = 1
//   x = 0
//
//  res will be fliped n number of times, n is the number of ones
//  starting from 0, if it flips even number of times then it'll be
//  back to 0, and if it flips odd number of times, it'll be 1
//
//  Time Complexity, O(n), n is the int size
func ParityNaive(x int) int {
	res := 0
	for x != 0 {
		res ^= x & 1
		x >>= 1
	}

	return res
}

// ParityRemoveLowestSetBit explanation
//  From ParityNaive we came to learn that we need to flip the res
//  n times, where n is the number of set bit. Here we're doing just
//  that.
//
//  Previously for input like `0b1001` we had to shift bit 4 times,
//  but now it's only two times
//
//  Time complexity, O(k), k is the number of set bit in the int
func ParityRemoveLowestSetBit(x int) int {
	res := 0
	for x != 0 {
		res ^= 1
		x &= x - 1 // This line clears the lowest set bit
	}

	return res
}
