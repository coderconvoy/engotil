package engotil

func RemoveGCollisionable(sl []GCollisionable, b IDable) []GCollisionable {
	id := b.ID()
	dp := -1
	for i, v := range sl {
		if v.ID() == id {
			dp = i
			break
		}
	}
	if dp >= 0 {
		return append(sl[:dp], sl[dp+1:]...)
	}
	return sl
}
func RemoveCollidable(sl []Collidable, b IDable) []Collidable {
	id := b.ID()
	dp := -1
	for i, v := range sl {
		if v.ID() == id {
			dp = i
			break
		}
	}
	if dp >= 0 {
		return append(sl[:dp], sl[dp+1:]...)
	}
	return sl
}
func RemoveSpaceable(sl []Spaceable, b IDable) []Spaceable {
	id := b.ID()
	dp := -1
	for i, v := range sl {
		if v.ID() == id {
			dp = i
			break
		}
	}
	if dp >= 0 {
		return append(sl[:dp], sl[dp+1:]...)
	}
	return sl
}
func RemoveAnimatable(sl []Animatable, b IDable) []Animatable {
	id := b.ID()
	dp := -1
	for i, v := range sl {
		if v.ID() == id {
			dp = i
			break
		}
	}
	if dp >= 0 {
		return append(sl[:dp], sl[dp+1:]...)
	}
	return sl
}
func RemoveMoveable(sl []Moveable, b IDable) []Moveable {
	id := b.ID()
	dp := -1
	for i, v := range sl {
		if v.ID() == id {
			dp = i
			break
		}
	}
	if dp >= 0 {
		return append(sl[:dp], sl[dp+1:]...)
	}
	return sl
}
