package engotil

import (
	"engo.io/ecs"
	"engo.io/engo/common"
)

//Engo
type IDable interface {
	ID() uint64
}
type BasicFace interface {
	GetBasicEntity() *ecs.BasicEntity
	ID() uint64
}

type AnimateFace interface {
	GetAnimationComponent() *common.AnimationComponent
	Cell() common.Drawable
	NextFrame()
}

type SpaceFace interface {
	GetSpaceComponent() *common.SpaceComponent
}

type RenderFace interface {
	GetRenderComponent() *common.RenderComponent
}

type CollisionFace interface {
	GetCollisionComponent() *common.CollisionComponent
}

//My Basic
type VelocityFace interface {
	GetVelocityComponent() *VelocityComponent
}

//Combi
type Collidable interface {
	BasicFace
	SpaceFace
	CollisionFace
}

type Spaceable interface {
	BasicFace
	SpaceFace
}

type Moveable interface {
	BasicFace
	SpaceFace
	VelocityFace
}

type Animatable interface {
	BasicFace
	AnimateFace
	RenderFace
}
