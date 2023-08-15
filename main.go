package main

import (
	"fmt"
	_ "examples.com/packages/functions"
	_"examples.com/packages/arrays"
	_"examples.com/packages/maps"
	"examples.com/packages/structs"
)

// func main() {
// 	//fmt.Print(functions.Sum(3, 2))
// 	//fmt.Print(functions.Number(4,3))
// 	//fmt.Println(arrays.LoopCounter())
// 	fmt.Println(arrays.Fibs())
// }


// func main() {
// 	// slice := []int{1, 2, 3}
// 	// fmt.Println(arrays.MangnifySlice(slice, 5))

// 	//fmt.Println(arrays.InsertSlice())
// 	//fmt.Println(arrays.InsertSliceIndex(3))

// 	//fmt.Println(arrays.SortSlice())
// 	// fmt.Println(arrays.BubbleSort())
// 	// fmt.Println(arrays.ReverseString())
// 	fmt.Println(maps.MapsDays(3))
// }

// Structs
func main(){
	//Anonymous Slice
	// calling := structs.Anonymous{
	// 	Float: 32,
	// 	Int : 32,
	// 	String : "Tolu",
	// }
	// fmt.Println(calling.Float, calling.Int, calling.String)

	//Employee Salary
	Emp := new(structs.Employee)
	Emp.Salary = 100000
	Emp.GiveRaise(0.98)
	fmt.Println(Emp.Salary)

}
