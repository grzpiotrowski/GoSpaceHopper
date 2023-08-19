package main

func getRectRectCollision(rectA, rectB FloatRect) bool {

	if rectA.X+rectA.W >= rectB.X &&
		rectA.X <= rectB.X+rectB.W &&
		rectA.Y+rectA.H >= rectB.Y &&
		rectA.Y <= rectB.Y+rectB.H {
		return true
	}

	return false
}
