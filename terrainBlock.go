package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type TerrainBlock struct {
	xBegin                       float32
	width                        float32
	elevationBegin, elevationEnd float32
	colour                       color.Color
}

func makeTerrainBlock(xBegin, width, elevationBegin, elevationEnd float32, colour color.Color) *TerrainBlock {
	return &TerrainBlock{
		xBegin:         xBegin,
		width:          width,
		elevationBegin: elevationBegin,
		elevationEnd:   elevationEnd,
		colour:         colour,
	}
}

func (tb *TerrainBlock) xEnd() float32 {
	return tb.xBegin + tb.width
}

func (tb *TerrainBlock) yBegin(screenHeight float32) float32 {
	return screenHeight - tb.elevationBegin
}

func (tb *TerrainBlock) yEnd(screenHeight float32) float32 {
	return screenHeight - tb.elevationEnd
}

func createSolidImage(colour color.Color) *ebiten.Image {
	img := ebiten.NewImage(1, 1)
	img.Fill(colour)
	return img
}

func (tb *TerrainBlock) Draw(screen *ebiten.Image) {

	options := &ebiten.DrawTrianglesOptions{}

	colourImage := createSolidImage(tb.colour)

	v := []ebiten.Vertex{
		{
			DstX:   tb.xBegin,
			DstY:   float32(screen.Bounds().Dy()),
			SrcX:   0,
			SrcY:   0,
			ColorR: 1,
			ColorG: 1,
			ColorB: 1,
			ColorA: 1,
		},
		{
			DstX:   tb.xBegin,
			DstY:   tb.yBegin(float32(screen.Bounds().Dy())),
			SrcX:   0,
			SrcY:   0,
			ColorR: 1,
			ColorG: 1,
			ColorB: 1,
			ColorA: 1,
		},
		{
			DstX:   tb.xEnd(),
			DstY:   tb.yEnd(float32(screen.Bounds().Dy())),
			SrcX:   0,
			SrcY:   0,
			ColorR: 1,
			ColorG: 1,
			ColorB: 1,
			ColorA: 1,
		},
		{
			DstX:   tb.xEnd(),
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
