package struct_01

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Shape interface {
	Area() float64
}

func (r Rectangle) Area() float64 {
	return 2 * (r.Width + r.Height)
}

func (c Circle) Area() float64 {
	const pi = 3.14
	return 2 * pi * c.Radius
}

func Area(s Shape) float64 {
	return s.Area()
}
