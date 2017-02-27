package engopoint

import "engo.io/engo"

func Abs(n float32) float32 {
	if n < 0 {
		return -n
	}
	return n
}

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

func Addn(pts ...engo.Point) engo.Point {
	res := &engo.Point{}
	for _, p := range pts {
		res.X += p.X
		res.Y += p.Y
	}
	return *res
}

func Angle8(p engo.Point) int {
	if Abs(p.X) > 2*Abs(p.Y) {
		if p.X > 0 {
			return 2
		}
		return 6
	}

	if Abs(p.Y) > 2*Abs(p.X) {
		if p.Y > 0 {
			return 4
		}
		return 0
	}
	if p.X < 0 {
		if p.Y < 0 {
			return 7
		}
		return 5
	}
	if p.Y < 0 {
		return 1
	}
	return 3

}
