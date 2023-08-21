package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Terrain struct {
	Blocks []TerrainBlock
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

		tb := makeTerrainBlock(lastXEnd, width, lastElevation, elevationEnd, colour)
		t.Blocks = append(t.Blocks, *tb)

		lastXEnd += width
		lastElevation = elevationEnd
	}
}

func (t *Terrain) Draw(screen *ebiten.Image) {
	for _, block := range t.Blocks {
		block.Draw(screen)
	}
}
