package engotil

import (
	"engo.io/engo"
	"engo.io/engo/common"
)

func SpaceCenter(sc common.SpaceComponent) engo.Point {
	return engo.Point{sc.Position.X + sc.Width/2, sc.Position.Y + sc.Height/2}
}
