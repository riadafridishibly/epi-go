package algo

// Merge2SortedArray merge two sorted arrays
func Merge2SortedArray(a, b []int) []int {
	if len(a) == 0 {
		return b
	}

	if len(b) == 0 {
		return a
	}

	newArray := make([]int, len(a)+len(b))

	aIdx := 0
	bIdx := 0
	curr := 0

	for aIdx < len(a) && bIdx < len(b) {
		if a[aIdx] < b[bIdx] {
			newArray[curr] = a[aIdx]
			aIdx++
		} else {
			newArray[curr] = b[bIdx]
			bIdx++
		}

		curr++
	}

	if aIdx < len(a) {
		copy(newArray[curr:], a[aIdx:])
	}

	if bIdx < len(b) {
		copy(newArray[curr:], b[bIdx:])
	}

	return newArray
}

// Merge2SortedArrayN merge two sorted arrays upto n elements
func Merge2SortedArrayN(a, b []int, n int) []int {
	if n == 0 {
		return nil
	}

	if len(a) == 0 {
		if len(b) > n {
			return b[:n]
		}

		return b
	}

	if len(b) == 0 {
		if len(a) > n {
			return a[:n]
		}

		return a
	}

	newArray := make([]int, n)

	aIdx := 0
	bIdx := 0
	curr := 0

	for curr < n && aIdx < len(a) && bIdx < len(b) {
		if a[aIdx] < b[bIdx] {
			newArray[curr] = a[aIdx]
			aIdx++
		} else {
			newArray[curr] = b[bIdx]
			bIdx++
		}

		curr++
	}

	if curr < n && aIdx < len(a) {
		copy(newArray[curr:], a[aIdx:n])
	}

	if curr < n && bIdx < len(b) {
		copy(newArray[curr:], b[bIdx:n])
	}

	return newArray
}
