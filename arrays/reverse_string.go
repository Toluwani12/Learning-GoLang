package arrays

func ReverseString () string {
	original:= "Google"
	reversed := ""
    for i := len(original) - 1; i >= 0; i-- {
        reversed += string(original[i])
    }
    return reversed 
}