package tileset

import (
	"encoding/json"
	"fmt"
	"image"
	"net/http"

	log "github.com/Sirupsen/logrus"

	"github.com/julienschmidt/httprouter"
)

type ProxyTileset struct {
	Tileset
	ProxyUrl       string
	ProxyType      string
	ProxyExtension string
}

func (t ProxyTileset) InfoHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	json.NewEncoder(w).Encode(t)
}

func (t ProxyTileset) GetTile(tile Tile) (Tile, error) {

	// initialize url to default value
	url := fmt.Sprintf("%s/%d/%d/%d", t.ProxyUrl, tile.Z, tile.Y, tile.X)

	switch t.ProxyType {
	case "zyx", "ZYX":
		url = fmt.Sprintf("%s/%d/%d/%d", t.ProxyUrl, tile.Z, tile.Y, tile.X)
	case "xyz", "XYZ":
		url = fmt.Sprintf("%s/%d/%d/%d", t.ProxyUrl, tile.X, tile.Y, tile.Z)
	default: // zyx
		url = fmt.Sprintf("%s/%d/%d/%d", t.ProxyUrl, tile.Z, tile.Y, tile.X)
	}

	if t.ProxyExtension != "" {
		url = fmt.Sprintf("%s.%s", url, t.ProxyExtension)
	}
	log.Infof("ProxyTileset::GetTile => requesting %s", url)
	response, err := http.Get(url)
	if err != nil {
		return EmptyTile(), err
	}

	defer response.Body.Close()

	m, _, err := image.Decode(response.Body)
	if err != nil {
		return EmptyTile(), err
	}

	tile.Image = m
	return tile, nil
}
