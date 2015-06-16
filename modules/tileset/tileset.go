package tileset

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"strconv"

	"github.com/gregbuehler/tilr/modules/setting"
)

type TilesetProvider interface {
	PutCache(Tile) (Tile, error)
	GetCache(Tile) (Tile, error)
	GetTile(Tile) (Tile, error)
}

type Tileset struct {
	Name string
	Type string
}

// func (t *Tileset) Info() string {
// 	// i, err := json.Marshal(t)
// 	// if err != nil {
// 	// 	log.Panic("error:", err)
// 	// }
// 	//
// 	// return string(i[:])
// }

func (t Tileset) PutCache(tile Tile) (Tile, error) {
	z := strconv.Itoa(tile.Z)
	y := strconv.Itoa(tile.Y)
	x := strconv.Itoa(tile.X)

	// get cache path
	tilepath := path.Join(
		setting.CacheLocation,
		t.Name,
		z,
		y,
		x+"."+setting.CacheFiletype,
	)

	// check if cache directories exist and create them
	// TODO: figure out proper permissions
	err := os.MkdirAll(path.Dir(tilepath), 0777)
	if err != nil {
		return EmptyTile(), err
	}

	f, err := os.Open(tilepath)
	if err != nil {
		return EmptyTile(), err
	}

	defer f.Close()

	switch setting.CacheFiletype {
	case "png":
		png.Encode(f, tile.Image)
	case "jpg", "jpeg":
		jpeg.Encode(f, tile.Image, &jpeg.Options{jpeg.DefaultQuality})
	default:
		return EmptyTile(), errors.New("invalid cache filetype")
	}

	return tile, nil
}

func (t Tileset) GetCache(tile Tile) (Tile, error) {
	z := strconv.Itoa(tile.Z)
	y := strconv.Itoa(tile.Y)
	x := strconv.Itoa(tile.X)

	// get cache path
	tilepath := path.Join(
		setting.CacheLocation,
		t.Name,
		z,
		y,
		x+"."+setting.CacheFiletype,
	)

	// check if tile exists
	if _, err := os.Stat(tilepath); err != nil {
		// load tile
		f, err := os.Open(tilepath)
		if err != nil {
			return EmptyTile(), err
		} else {
			defer f.Close()
		}

		tile.Image, _, err = image.Decode(f)
		if err != nil {
			return EmptyTile(), err
		} else {
			return tile, nil
		}
	}

	return EmptyTile(), nil
}

func (t Tileset) GetTile(tile Tile) (Tile, error) {
	// This is only a stub. It should always have an implementation override
	return EmptyTile(), nil
}
