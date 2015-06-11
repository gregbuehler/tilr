package tileserver

import (
	"image"

	"github.com/gregbuehler/tilr/modules/tileset"
)

type TileServer interface {
	Read(t tileset.Tile) (i image.Image, err error)
}
