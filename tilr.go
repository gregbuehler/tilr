package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

const (
	TilrVersion = "0.0.1"
	TilrBind    = ":5555"
)

const (
	TileCachePath = "/var/tilr/cache"
	TileDimension = 256
	TileQuality   = 80
)

type Tile struct {
	tileset string
	x       int
	y       int
	z       int
}

func RetrieveTile(t Tile) (i image.Image, e error) {
	fileDirectory := path.Join(
		TileCachePath,
		t.tileset,
		strconv.Itoa(t.z),
	)

	file := path.Join(
		TileCachePath,
		t.tileset,
		strconv.Itoa(t.z),
		strings.Join([]string{strconv.Itoa(t.y), strconv.Itoa(t.x)}, "_"),
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

	var opt jpeg.Options
	opt.Quality = TileQuality

	jpeg.Encode(out, j, &opt)
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

	background := color.RGBA{0, 255, 255, 255}

	if (t.x+t.y)%2 == 0 {
		background = color.RGBA{0, 0, 255, 255}
	}

	draw.Draw(m, m.Bounds(), &image.Uniform{background}, image.ZP, draw.Src)

	return m, nil
}

func MetaHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Tilr v%s", TilrVersion)
}

func TilesetHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "tileset: %s\n", ps.ByName("tileset"))
}

func TileHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
		tileset: ps.ByName("tileset"),
		z:       z,
		y:       y,
		x:       x,
	}

	i, err := RetrieveTile(t)
	if err != nil {
		log.Panic(err)
	} else {
		var opt jpeg.Options
		opt.Quality = TileQuality

		jpeg.Encode(w, i, &opt)
	}
}

func main() {
	log.Println("Starting Tilr v%s", TilrVersion)

	// handle routes
	router := httprouter.New()
	router.GET("/", MetaHandler)
	router.GET("/:tileset", TilesetHandler)
	router.GET("/:tileset/:z/:y/:x", TileHandler)

	log.Println("Listening on %s", TilrBind)
	log.Fatal(http.ListenAndServe(TilrBind, router))
}
