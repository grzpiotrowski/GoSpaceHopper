package main

import "github.com/hajimehoshi/ebiten/v2"

func GetStickState() Vec2f {
	stick := Vec2f{}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		stick.X += -1.
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		stick.X += 1.
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		stick.Y += -1.
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		stick.Y += 1.
	}

	return stick
}
