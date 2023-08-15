package interfaces

type Simplers interface{
	Geto() int
	Seto(t int)
}

type Simples struct{
	length int
}

func (r *Simples) Geto() int{
	return r.length
}

func (r *Simples) Seto(t int){
	r.length = t
}

type RSimple struct{
	num1 int
	num2 int
}

func (r RSimple) Geto() int{
	return r.num1
}

func (r RSimple) Seto(t int){
	r.num2 = t
}

func FI (it Simplers) int{
	switch it.(type) {
	case *Simples:
		it.Seto(5)
		return it.Geto()

	case *RSimple:
		it.Seto(50)
		return it.Geto()

	default:
		return 99
	}
	return 0
}

func CallBoths(them Simplers) int{
	them.Seto(2)
	return them.Geto()
}