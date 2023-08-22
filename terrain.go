package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Terrain struct {
	Blocks []*TerrainBlock
}

func makeTerrain() *Terrain {
	return &Terrain{}
}

func (t *Terrain) generateTerrain(numBlocks int) {
	rand.Seed(time.Now().UnixNano())

	lastXEnd := 0.00
	lastElevation := 50.0

	for i := 0; i < numBlocks; i++ {
		width := 100.00
		elevationEnd := 50 + rand.Float64()*(100-50)
		colour := color.RGBA{50, 30, 60, 100}

		tb := makeTerrainBlock(lastXEnd, width, lastElevation, elevationEnd, colour, i)
		t.Blocks = append(t.Blocks, tb)

		lastXEnd += width
		lastElevation = elevationEnd
	}
}

func SurroundingTerrainBlocks(t *Terrain, i, renderRadius int) []*TerrainBlock {
	if i < 0 || i >= len(t.Blocks) {
		return nil // return nil or handle error if index out of range
	}
	start := i - renderRadius
	if start < 0 {
		start = 0
	}
	end := i + renderRadius + 1
	if end > len(t.Blocks) {
		end = len(t.Blocks)
	}
	return t.Blocks[start:end]
}

func (g *Game) updateTerrain() {
	t := g.t
	for _, block := range t.Blocks {
		block.Scroll(*g.hero.Movement)
	}
}

func (t *Terrain) Draw(screen *ebiten.Image, i int) {

	slice := SurroundingTerrainBlocks(t, i, 1)

	for _, block := range slice {
		block.Draw(screen)
	}
}
