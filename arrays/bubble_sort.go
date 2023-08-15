package arrays

import ("sort")

//if it's an array
func SortSlice() [] int{
	Slice := []int {4,3,6,7,1,0,4}

	sort.Ints(Slice)

	return Slice
}

//lt's use the bubble sort format
func BubbleSort () []int{
	numbers := []int {3,4,7,9,0,2,3,4}
	n := len(numbers)

	for i := 0; i < n; i ++{
		// Each iteration moves the largest element to the end
		for j:=0; j< n-1; j++{
			// Compare adjacent elements and swap if necessary
			if (numbers[j] > numbers[j+1]){
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
	return numbers
}