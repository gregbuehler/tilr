package tileset

import (
	"image"
	"image/color"
	"image/draw"
)

type Tile struct {
	Z, Y, X   int
	Lat, Long float32
	Image     image.Image
}

func EmptyTile() (t Tile) {
	tile := new(Tile)

	m := image.NewRGBA(image.Rect(0, 0, 256, 256))
	gray := color.RGBA{0, 128, 128, 128}
	draw.Draw(
		m,
		m.Bounds(),
		&image.Uniform{gray},
		image.ZP,
		draw.Src,
	)

	tile.Image = m
	return *tile
}
