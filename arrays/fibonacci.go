package arrays

func Fibs() [15]int64{
	var fib [15]int64
	fib[0] = 1		
	fib[1] = 1	
	
	for i:= 2; i <15; i++ {
		fib[i] = fib[i-1] + fib[i-2]	
	}
    return fib
}