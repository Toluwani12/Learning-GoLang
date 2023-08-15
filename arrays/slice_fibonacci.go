package arrays

func FibArray(term int) []int{
	slice := make([]int, term)
	slice [0] = 1
	slice [1] = 1

	for i:= 2; i < term; i++ {
		slice[i] = slice[i-1] + slice[i-2]	
	}
    return slice
}