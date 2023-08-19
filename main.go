package main

import (
	"embed"
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed data
var data embed.FS

const (
	gameScreenWidth  = 640
	gameScreenHeight = 360

	scale        = 2
	windowWidth  = gameScreenWidth * scale
	windowHeight = gameScreenHeight * scale

	dt = 1 / 60.0
)

type Game struct {
	hero    *Hero
	monster *Monster
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

func (g *Game) processInput() {
	stickState := GetStickState()
	h := g.hero
	h.Movement.Velocity.X = stickState.X * h.Movement.Speed.X
	h.Movement.Velocity.Y = stickState.Y * h.Movement.Speed.Y
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

	// monster
	monster, err := makeMonster()
	if err != nil {
		return err
	}
	g.monster = monster

	return nil
}

// Called every frame by the game loop
func (g *Game) Update() error {
	g.processInput()
	g.updateMonster()
	g.updateHero()

	var hit bool = g.hero.collidesWith(&g.monster.Entity)

	fmt.Println(fmt.Sprint(hit))

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	g.monster.Draw(screen)
	g.hero.Draw(screen)
	// Debug text
	//fmt.Println(g.hero.getAABB().String())

}
