package structs

type Employee struct{
	Salary float32
}

func (e Employee) GiveRaise(percent float32){
	e.Salary += e.Salary * percent
}