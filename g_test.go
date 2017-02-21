package engotil

import (
	"image/color"
	"testing"

	"engo.io/ecs"
	"engo.io/engo/common"
)

type Anime struct {
	Basic
}

func Test_Add(t *testing.T) {

	k := Anime{
		BasicEntity:        BasicEntity{ecs.NewBasic()},
		AnimationComponent: common.NewAnimationComponent([]Drawable{}, 0.1),
		RenderComponent: RenderComponent{
			Drawable: Triangle{},
			Color:    color.Black,
		},
	}
	a := AnimationSystem{}

	a.AddByInterface(k)

}
