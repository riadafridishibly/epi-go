package ch5

func DeleteDuplicate(a []int) int {
	if len(a) == 0 {
		return 0
	}

	writeIndex := 1
	for i := 1; i < len(a); i++ {
		if a[writeIndex-1] != a[i] {
			a[writeIndex] = a[i]
			writeIndex++
		}
	}

	return writeIndex
}
