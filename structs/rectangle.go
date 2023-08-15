package structs

type Rectangle struct{
	length, width int64
}

func  (r Rectangle) Area() int64{
	return r.length * r.width
}

func (R Rectangle) Perimeter() int64{
	return 2*(R.length + R.width)
}