package arrays

func InsertSlice() []string {
    slice := []string {"M", "N", "O", "P", "Q", "R"}
	insert := []string {"A", "B","C"}
    slice = append (slice, insert...)
    return slice
}

//To use in a certain index

func InsertSliceIndex(index int) []string {
	slice := []string {"M", "N", "O", "P", "Q", "R"}
	insertSlice := []string {"A", "B","C"}

	// Create a new slice with length equal to the combined length of both slices
	newslice := make ([]string, len(slice) + len(insertSlice))

	// Copy the elements before the insertion index from the original slice to the new slice in another variable
	result := copy(newslice, slice[:index])

	// Copy the elements from the insert slice to the new slice and add to the result
	result += copy (newslice[index:], insertSlice )

	// Copy the remaining elements from the original slice to the new slice
	copy (newslice[result:], slice[index:])

	return newslice
}