package arrays

func Filter(s []int, fn func(int) bool) []int{
	// Create an empty slice to store the filtered elements
	result := make([]int, 0)

	// Iterate over each element of the input slice s
	for _, i := range s {
		// Apply the provided function fn to the current element
		if fn(i) {
			// If the function returns true, add the element to the result slice
			result = append(result, i)
		}

	}
	return result
}

// func filterEven(num int) bool {
//     return num%2 == 0
// }