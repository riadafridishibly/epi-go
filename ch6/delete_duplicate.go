package ch6

func DeleteDuplicateInt(a []int) int {
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

func DeleteDuplicate(slice []interface{}, equal func(a, b interface{}) bool) (lastIndex int) {
	if len(slice) == 0 {
		return 0
	}

	writeIndex := 1
	for i := 1; i < len(slice); i++ {
		if !equal(slice[writeIndex-1], slice[i]) {
			slice[writeIndex] = slice[i]
			writeIndex++
		}
	}

	return writeIndex
}
