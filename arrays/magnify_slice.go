package arrays


func MangnifySlice(s []int, factor int) []int {
	magnify := make ([]int, len(s) * factor)
	copy(magnify, s)

	return magnify
}