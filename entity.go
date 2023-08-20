package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Entity struct {
	Transform *TransformComponent
	Movement  *MovementComponent
	Graphics  *GraphicsComponent
}

type TransformComponent struct {
	Position  Vec2f
	Direction Vec2f
	Scale     Vec2f
}

type MovementComponent struct {
	Speed    Vec2f
	Velocity Vec2f
}

type GraphicsComponent struct {
	Sprite *Sprite
}

func NewEntity(scale Vec2f) Entity {
	return Entity{
		Transform: &TransformComponent{
			Scale: scale,
		},
		Movement: &MovementComponent{},
		Graphics: &GraphicsComponent{},
	}
}

func (e *Entity) getAABB() FloatRect {
	w := e.Graphics.Sprite.GetSize().X * e.Transform.Scale.X
	h := e.Graphics.Sprite.GetSize().Y * e.Transform.Scale.Y
	return FloatRect{e.Transform.Position.X, e.Transform.Position.Y, w, h}
}

func (e *Entity) collidesWith(other *Entity) bool {
	return e.getAABB().overlaps(other.getAABB())
}

func (e *Entity) onGround(tb TerrainBlock) bool {

	var rect FloatRect = e.getAABB()

	return checkLineSegmentsIntersection(rect.X+rect.W/2, rect.Y+rect.H/2, rect.X+rect.W/2, rect.Y+rect.H,
		float64(tb.xBegin), float64(tb.yBegin(gameScreenHeight)), float64(tb.xEnd()), float64(tb.yEnd(gameScreenHeight)),
	)
}

func (e *Entity) groundCollision(tb TerrainBlock) {

	hLine := e.verticalCollisionLine()
	tbLine := tb.collisionLine()

	var tbIntersectionPoint Vec2f = GetLinesIntersectionPoint(hLine, tbLine)

	if e.onGround(tb) {
		if e.getAABB().Y+e.getAABB().H >= tbIntersectionPoint.Y {
			e.Transform.Position.Y = tbIntersectionPoint.Y - e.getAABB().H
		}
	}
}

func (e *Entity) verticalCollisionLine() Line {
	rect := e.getAABB()
	l := Line{
		Begin: Vec2f{X: rect.X + rect.W/2, Y: rect.Y + rect.H/2},
		End:   Vec2f{X: rect.X + rect.W/2, Y: rect.Y + rect.H},
	}
	return l
}

func (e *Entity) drawDebugLines(screen *ebiten.Image) {
	var rect FloatRect = e.getAABB()
	var clr color.Color = color.RGBA{255, 0, 0, 100}

	vector.StrokeLine(screen, float32(rect.X)+float32(rect.W)/2, float32(rect.Y)+float32(rect.H/2), float32(rect.X)+float32(rect.W)/2, float32(rect.Y)+float32(rect.H), 2, clr, false)
}

func (e *Entity) update(tb TerrainBlock) {
	e.groundCollision(tb)
}

func (e *Entity) Draw(screen *ebiten.Image) {
	var m ebiten.GeoM
	m.Scale(e.Transform.Scale.X, e.Transform.Scale.Y)
	m.Translate(
		e.Transform.Position.X,
		e.Transform.Position.Y,
	)
	e.Graphics.Sprite.Draw(screen, m)
	e.drawDebugLines(screen)
}
