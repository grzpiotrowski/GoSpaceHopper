package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Hero struct {
	Entity
}

func makeHero() (*Hero, error) {
	e := &Hero{
		Entity: NewEntity(),
	}

	e.Movement.Speed = Vec2f{120, 120}

	heroWidth, heroHeight := 50, 50
	img := ebiten.NewImage(heroWidth, heroHeight)
	img.Fill(color.RGBA{R: 0, G: 0, B: 255, A: 255})

	e.Graphics.Rect = img

	return e, nil
}

func (g *Game) updateHero() {
	t := g.hero.Transform
	m := g.hero.Movement

	t.Position.X += m.Velocity.X * dt
	t.Position.Y += m.Velocity.Y * dt
}
