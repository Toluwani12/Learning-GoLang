package functions 

import ("fmt")

func Number(num1, num2 float32) (Prod, Sum, Diff float32) {
	Prod = num1 * num2
	Sum = num1 + num2
	Diff = num1 - num2 
	fmt.Printf("The Prod is %v, Sum is %v and Diff is %v", Prod, Sum, Diff)
	return
}