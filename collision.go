package main

func (rectA FloatRect) overlaps(rectB FloatRect) bool {

	if rectA.X+rectA.W >= rectB.X &&
		rectA.X <= rectB.X+rectB.W &&
		rectA.Y+rectA.H >= rectB.Y &&
		rectA.Y <= rectB.Y+rectB.H {
		return true
	}

	return false
}

func checkLineSegmentsIntersection(xA, yA, xB, yB, xC, yC, xD, yD float64) bool {
	d := (xA-xB)*(yC-yD) - (yA-yB)*(xC-xD)
	if d == 0 {
		return false
	}

	t := ((xA-xC)*(yC-yD) - (yA-yC)*(xC-xD)) / d
	u := ((xA-xC)*(yA-yB) - (yA-yC)*(xA-xB)) / d

	return 0 <= t && t <= 1 && 0 <= u && u <= 1
}

func GetLinesIntersectionPoint(lineA, lineB Line) Vec2f {

	xA := lineA.Begin.X
	yA := lineA.Begin.Y
	xB := lineA.End.X
	yB := lineA.End.Y
	xC := lineB.Begin.X
	yC := lineB.Begin.Y
	xD := lineB.End.X
	yD := lineB.End.Y

	d := (xA-xB)*(yC-yD) - (yA-yB)*(xC-xD)

	xP := ((xA*yB-yA*xB)*(xC-xD) - (xA-xB)*(xC*yD-yC*xD)) / d
	yP := ((xA*yB-yA*xB)*(yC-yD) - (yA-yB)*(xC*yD-yC*xD)) / d

	return Vec2f{xP, yP}
}
