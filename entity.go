package main

import (
	"github.com/hajimehoshi/ebiten/v2"
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

func (e *Entity) Draw(screen *ebiten.Image) {
	var m ebiten.GeoM
	m.Scale(e.Transform.Scale.X, e.Transform.Scale.Y)
	m.Translate(
		e.Transform.Position.X,
		e.Transform.Position.Y,
	)
	e.Graphics.Sprite.Draw(screen, m)
}
