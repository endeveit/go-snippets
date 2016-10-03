package arrays

// Checks if string exists in a slice
// Returns the key for needle if it is found in the slice, -1 otherwise.
func StringInSlice(haystack []string, needle string) (result int) {
	result = -1

	for k, v := range haystack {
		if v == needle {
			result = k
			break
		}
	}

	return
}

// Checks if integer exists in a slice
// Returns the index for needle if it is found in the slice, -1 otherwise.
func IntInSlice(haystack []int, needle int) (index int) {
	index = -1

	for k, v := range haystack {
		if v == needle {
			index = k
			break
		}
	}

	return
}
