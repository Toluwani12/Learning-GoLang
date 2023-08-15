package functions

func Sum(list ...int) (sum int){
	for _,v := range list{
		sum = sum + v
	}
	return
}