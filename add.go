package engotil

import "engo.io/engo/common"

//Animation
type AnimationComponent common.AnimationComponent

func (ac *AnimationComponent) GetAnimationComponent() *common.AnimationComponent {
	return ac.AnimationComponent
}

type AnimationSystem common.AnimationSystem

func (a *AnimationSystem) AddByInterface(c Animatable) {
	cd.Add(c.GetBasicEntity(), c.GetAnimationComponent(), c.GetRenderComponent())
}

//Collision
type CollisionSystem common.CollisionSystem

func (cs *CollisionSystem) AddByInterface(c Collidable) {
	cs.Add(c.GetBasicEntity(), c.GetCollisionComponent(), c.GetSpaceComponent())
}
