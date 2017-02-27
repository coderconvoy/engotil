package engopoint

import "engo.io/engo"

func Neg(a engo.Point) engo.Point {
	return engo.Point{-a.X, -a.Y}
}
func Sub(a, b engo.Point) engo.Point {

	return engo.Point{
		a.X - b.X,
		a.Y - b.Y,
	}
}

func Add(a, b engo.Point) engo.Point {
	return engo.Point{
		a.X + b.X,
		a.Y + b.Y,
	}
}
