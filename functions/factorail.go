package functions

func Factorial (num int64) (num1 int64){
	if (num < 1){
		return 1
	}
	num1 = num * Factorial(num - 1)
	return
}