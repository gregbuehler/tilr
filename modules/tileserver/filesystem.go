package tileserver

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"path"
	"strconv"
)

func (t FilesystemTileServer) Read(t Tile) (i image.Image, err error) {

}

func RetrieveTile(t Tile) (i image.Image, e error) {
	fileDirectory := path.Join(
		TileCachePath,
		t.tileset,
		strconv.Itoa(t.z),
		strconv.Itoa(t.y),
	)

	filename := strconv.Itoa(t.x) + ".png"
	file := path.Join(
		TileCachePath,
		t.tileset,
		strconv.Itoa(t.z),
		strconv.Itoa(t.y),
		filename,
	)

	if _, err := os.Stat(file); err == nil {
		r, err := os.Open(file)
		if err != nil {
			return nil, err
		} else {
			defer r.Close()
		}

		i, _, err := image.Decode(r)
		if err != nil {
			return nil, err
		}

		log.Printf("CH\tpath: %s, tileset: %s, z: %d, y: %d, x: %d\n",
			file,
			t.tileset,
			t.z,
			t.y,
			t.x,
		)

		return i, nil
	}

	log.Printf("CM\tpath: %s, tileset: %s, z: %d, y: %d, x: %d\n",
		file,
		t.tileset,
		t.z,
		t.y,
		t.x,
	)

	j, err := RenderTile(t)
	if err != nil {
		return nil, err
	}

	os.MkdirAll(fileDirectory, os.ModeDir)
	out, err := os.Create(file)
	if err != nil {
		return nil, err
	}

	png.Encode(out, j)
	return j, nil
}

func RenderTile(t Tile) (i image.Image, err error) {

	log.Printf("RT\ttileset: %s, z: %d, y: %d, x: %d\n",
		t.tileset,
		t.z,
		t.y,
		t.x,
	)

	m := image.NewRGBA(image.Rect(0, 0, TileDimension, TileDimension))

	background := color.RGBA{0, 255, 0, 255}

	// checkerboard pattern
	// if (t.x+t.y)%2 == 0 {
	// 	background = color.RGBA{0, 0, 255, 255}
	// }

	draw.Draw(m, m.Bounds(), &image.Uniform{background}, image.ZP, draw.Src)

	return m, nil
}
