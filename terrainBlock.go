package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type TerrainBlock struct {
	xBegin                       float64
	width                        float64
	elevationBegin, elevationEnd float64
	colour                       color.Color
	index                        int
}

func makeTerrainBlock(xBegin, width, elevationBegin, elevationEnd float64, colour color.Color, index int) *TerrainBlock {
	return &TerrainBlock{
		xBegin:         xBegin,
		width:          width,
		elevationBegin: elevationBegin,
		elevationEnd:   elevationEnd,
		colour:         colour,
		index:          index,
	}
}

func (tb *TerrainBlock) xEnd() float64 {
	return tb.xBegin + tb.width
}

func (tb *TerrainBlock) yBegin() float64 {
	return gameScreenHeight - tb.elevationBegin
}

func (tb *TerrainBlock) yEnd() float64 {
	return gameScreenHeight - tb.elevationEnd
}

func (tb *TerrainBlock) collisionLine() Line {
	return Line{
		Begin: Vec2f{X: tb.xBegin, Y: tb.yBegin()},
		End:   Vec2f{X: tb.xEnd(), Y: tb.yEnd()},
	}
}

func createSolidImage(colour color.Color) *ebiten.Image {
	img := ebiten.NewImage(1, 1)
	img.Fill(colour)
	return img
}

func (tb *TerrainBlock) Scroll(m MovementComponent) {
	tb.xBegin -= m.Velocity.X * dt
}

func (tb *TerrainBlock) Draw(screen *ebiten.Image) {

	options := &ebiten.DrawTrianglesOptions{}

	colourImage := createSolidImage(tb.colour)

	v := []ebiten.Vertex{
		{
			DstX:   float32(tb.xBegin),
			DstY:   float32(screen.Bounds().Dy()),
			SrcX:   0,
			SrcY:   0,
			ColorR: 1,
			ColorG: 1,
			ColorB: 1,
			ColorA: 1,
		},
		{
			DstX:   float32(tb.xBegin),
			DstY:   float32(tb.yBegin()),
			SrcX:   0,
			SrcY:   0,
			ColorR: 1,
			ColorG: 1,
			ColorB: 1,
			ColorA: 1,
		},
		{
			DstX:   float32(tb.xEnd()),
			DstY:   float32(tb.yEnd()),
			SrcX:   0,
			SrcY:   0,
			ColorR: 1,
			ColorG: 1,
			ColorB: 1,
			ColorA: 1,
		},
		{
			DstX:   float32(tb.xEnd()),
			DstY:   float32(screen.Bounds().Dy()),
			SrcX:   0,
			SrcY:   0,
			ColorR: 1,
			ColorG: 1,
			ColorB: 1,
			ColorA: 1,
		},
	}
	i := []uint16{0, 1, 2, 0, 2, 3}

	screen.DrawTriangles(v, i, colourImage, options)
}
