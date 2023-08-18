package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Entity struct {
	Transform *TransformComponent
	Movement  *MovementComponent
	Graphics  *GraphicsComponent
}

type Vec2f struct {
	X, Y float64
}

type TransformComponent struct {
	Position  Vec2f
	Direction Vec2f
}

type MovementComponent struct {
	Speed    Vec2f
	Velocity Vec2f
}

type GraphicsComponent struct {
	Rect *ebiten.Image
}

type FloatRect struct {
	X, Y, W, H float64
}

func NewEntity() Entity {
	rect := ebiten.NewImage(50, 50)
	rect.Fill(color.White)

	return Entity{
		Transform: &TransformComponent{},
		Movement:  &MovementComponent{},
		Graphics: &GraphicsComponent{
			Rect: rect,
		},
	}
}

func (e *Entity) getAABB() FloatRect {
	w, h := float64(e.Graphics.Rect.Bounds().Dx()), float64(e.Graphics.Rect.Bounds().Dy())
	return FloatRect{e.Transform.Position.X, e.Transform.Position.Y, w, h}
}

func (e *Entity) Draw(screen *ebiten.Image) {
	var m ebiten.GeoM
	m.Translate(
		e.Transform.Position.X,
		e.Transform.Position.Y,
	)
	screen.DrawImage(e.Graphics.Rect, &ebiten.DrawImageOptions{
		GeoM: m,
	})
}
