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

	e.Movement.Speed = Vec2f{120, 0}

	heroWidth, heroHeight := 50, 50
	img := ebiten.NewImage(heroWidth, heroHeight)
	img.Fill(color.RGBA{R: 0, G: 0, B: 255, A: 255})

	e.Graphics.Rect = img

	return e, nil
}