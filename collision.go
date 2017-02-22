package engotil

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type Collision2Message struct {
	Main  Collidable
	Buddy Collidable
}

func (Collision2Message) Type() string { return "Collision2Message" }

type CollisionSystem struct {
	entities []Collidable
}

func (c *CollisionSystem) Add(ob Collidable) {
	c.entities = append(c.entities, ob)
}

func (c *CollisionSystem) Remove(basic ecs.BasicEntity) {
	c.entities = RemoveCollidable(c.entities, basic)
}

func (cs *CollisionSystem) Update(dt float32) {
	for i1, e1 := range cs.entities {
		cc1 := e1.GetCollisionComponent()
		if !cc1.Main {
			continue // with other entities
		}
		sc1 := e1.GetSpaceComponent()

		entityAABB := sc1.AABB()
		offset := engo.Point{cc1.Extra.X / 2, cc1.Extra.Y / 2}
		entityAABB.Min.X -= offset.X
		entityAABB.Min.Y -= offset.Y
		entityAABB.Max.X += offset.X
		entityAABB.Max.Y += offset.Y

		var collided bool

		for i2, e2 := range cs.entities {
			if i1 == i2 {
				continue // with other entities, because we won't collide with ourselves
			}
			cc2 := e2.GetCollisionComponent()
			sc2 := e2.GetSpaceComponent()

			otherAABB := sc2.AABB()
			offset = engo.Point{cc2.Extra.X / 2, cc2.Extra.Y / 2}
			otherAABB.Min.X -= offset.X
			otherAABB.Min.Y -= offset.Y
			otherAABB.Max.X += offset.X
			otherAABB.Max.Y += offset.Y

			if common.IsIntersecting(entityAABB, otherAABB) {
				if cc1.Solid && cc2.Solid {
					mtd := common.MinimumTranslation(entityAABB, otherAABB)
					sc1.Position.X += mtd.X
					sc1.Position.Y += mtd.Y
				}

				collided = true
				engo.Mailbox.Dispatch(Collision2Message{Main: e1, Buddy: e2})
			}
		}

		cc1.Collides = collided
	}
}
