package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	gameScreenWidth  = 640
	gameScreenHeight = 360

	scale        = 2
	windowWidth  = gameScreenWidth * scale
	windowHeight = gameScreenHeight * scale

	dt = 1 / 60.0
)

type Game struct {
	hero *Hero
}

func main() {

	g := &Game{}
	if err := g.Init(); err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Space Hopper")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Update() error {
	g.updateHero()

	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 360
}

func (g *Game) Init() error {

	if err := g.loadObjects(); err != nil {
		return err
	}
	return nil
}

// Initialising game objects
func (g *Game) loadObjects() error {

	// hero
	hero, err := makeHero()
	if err != nil {
		return err
	}
	g.hero = hero

	return nil
}

func (g *Game) updateHero() {
	t := g.hero.Transform
	m := g.hero.Movement

	t.Position.X += m.Velocity.X * dt

	newPos := Vec2f{50, 50}
	t.Position = newPos
}

// Called every frame by the game loop
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	g.hero.Draw(screen)
}
