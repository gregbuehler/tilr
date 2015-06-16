package tileset

import (
	"encoding/json"
	"image"
	"net/http"
	"os"
	"path"
	"strconv"

	log "github.com/Sirupsen/logrus"

	"github.com/julienschmidt/httprouter"
)

type FilesystemTileset struct {
	Tileset
	TilesetLocation  string
	TilesetImageType string
}

func (t FilesystemTileset) InfoHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	json.NewEncoder(w).Encode(t)
}

func (t FilesystemTileset) GetTile(tile Tile) (Tile, error) {
	z := strconv.Itoa(tile.Z)
	y := strconv.Itoa(tile.Y)
	x := strconv.Itoa(tile.X)

	// get cache path
	tilepath := path.Join(
		t.TilesetLocation,
		t.Name,
		z,
		y,
		x+"."+t.TilesetImageType,
	)

	log.Infof("Filesystem(%s)::GetTile: %s, %s, %s => %s", t.Name, z, y, x, tilepath)

	// check if tile exists
	if _, err := os.Stat(tilepath); err != nil {
		log.Errorf("GetTile: Tile %s does not exist. %v", tilepath, err)
		return EmptyTile(), err
	} else {
		// load tile
		f, err := os.Open(tilepath)
		if err != nil {
			log.Errorf("GetTile: Failed to open %s. %v", tilepath, err)
			return EmptyTile(), err
		}

		defer f.Close()

		tile.Image, _, err = image.Decode(f)
		if err != nil {
			return EmptyTile(), err
		}

		return tile, nil
	}
}
