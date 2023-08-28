package interfaces

type Square struct {
	Side float32
}

type Triangle struct {	// implement this struct
	Base float32
	Height float32
}

type AreaInterface interface {
	Area() float32
}

type PeriInterface interface { // implement this interface 
	Perimeter() float32
}

func (sq *Square) Area() float32 {
	return sq.Side * sq.Side
}

func (sq *Square) Perimeter() float32 { // implement method called on square to calculate its perimeter
	return 4 * sq.Side
}

func (tr *Triangle) Area() float32 { // implement method called on triangle to calculate its area
	return 0.5 * tr.Base * tr.Height
}