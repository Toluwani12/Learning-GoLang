package interfaces

type Simpler interface{
	Get() int
	Set(t int) int
}

type Simple struct{
	length int
}

func (r *Simple) Get() int{
	return r.length
}

func (r *Simple) Set(t int){
	r.length = t
}

func CallBoth(them Simpler) int{
	them.Set(2)
	return them.Get()
}