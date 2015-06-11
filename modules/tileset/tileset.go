package tileset

import (
	"fmt"
	"net/http"
	"strconv"

	log "github.com/Sirupsen/logrus"

	"github.com/julienschmidt/httprouter"
)

type Tile struct {
	x, y, z   int
	lat, long float32
}

func TilesetHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "tileset: %s\n", ps.ByName("tileset"))
}

func TileHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tileset := ps.ByName("tileset")
	z, err := strconv.Atoi(ps.ByName("z"))
	if err != nil {
		log.Panic(err)
	}

	y, err := strconv.Atoi(ps.ByName("y"))
	if err != nil {
		log.Panic(err)
	}

	x, err := strconv.Atoi(ps.ByName("x"))
	if err != nil {
		log.Panic(err)
	}

	t := Tile{
		z: z,
		y: y,
		x: x,
	}

	log.Printf("%s, %d, %d, %d",
		tileset,
		t.z,
		t.y,
		t.z,
	)

	// i, err := RetrieveTile(tileset, t)
	// if err != nil {
	// 	log.Panic(err)
	// } else {
	// 	png.Encode(w, i)
	// }
}
