package cmd

import (
	"encoding/json"
	"fmt"
	"image/jpeg"
	"net"
	"net/http"
	"strconv"

	log "github.com/Sirupsen/logrus"

	"github.com/codegangsta/cli"

	"github.com/gregbuehler/tilr/modules/meta"
	"github.com/gregbuehler/tilr/modules/setting"
	"github.com/gregbuehler/tilr/modules/tileset"

	"github.com/julienschmidt/httprouter"
)

var CmdWeb = cli.Command{
	Name:        "web",
	Usage:       "Start Tilr server",
	Description: `Tilr server does everything for you`,
	Action:      runWeb,
	Flags: []cli.Flag{
		cli.StringFlag{"port, p", "3000", "Temporary port number to prevent conflict", ""},
		cli.StringFlag{"config, c", "custom/conf/app.yaml", "Custom configuration file path", ""},
	},
}

var tilesets = make(map[string]tileset.TilesetProvider)

func InfoHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tileset := ps.ByName("tileset")
	log.Infof("InfoHandler: %s	%v", r.Host, ps)
	json.NewEncoder(w).Encode(tilesets[tileset])
}

func TileHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	name := ps.ByName("tileset")
	z, _ := strconv.Atoi(ps.ByName("z"))
	y, _ := strconv.Atoi(ps.ByName("y"))
	x, _ := strconv.Atoi(ps.ByName("x"))

	log.Infof("TileHandler: %s => %d, %d, %d", name, z, y, x)

	ts := tilesets[name]
	t, err := ts.GetTile(tileset.Tile{
		Z: z,
		Y: y,
		X: x,
	})
	if err != nil {
		log.Errorf("TileHandler: Failed to get tile(%v). %v", t, err)
	}

	log.Infof("TileHandler: tileset(%v), tile(%v)", ts, t)

	switch setting.TileFileType {
	case "jpg", "jpeg":
		log.Infof("Sending jpeg response")
		jpeg.Encode(w, t.Image, &jpeg.Options{jpeg.DefaultQuality})
	// case "png":
	// 	log.Infof("Sending png response")
	// 	png.Encode(w, t.Image)
	default:
		log.Debugf("Sending default(jpeg) response")
		jpeg.Encode(w, t.Image, &jpeg.Options{jpeg.DefaultQuality})
	}
}

func runWeb(ctx *cli.Context) {
	if ctx.IsSet("config") {
		setting.Load(ctx.String("config"))
	}

	if ctx.IsSet("port") {
		setting.Port = ctx.String("port")
	}

	// tilesets := map[string]tileset.TilesetProvider{
	// 	"foo": tileset.FilesystemTileset{
	// 		TilesetLocation:  "/var/tilr/foo",
	// 		TilesetImageType: "png",
	// 	},
	// 	"bar": tileset.ProxyTileset{
	// 		ProxyUrl:  "https://b.tile.openstreetmap.org/",
	// 		ProxyType: "zyx",
	// 	},
	// }

	tilesets["fs"] = tileset.FilesystemTileset{
		Tileset: tileset.Tileset{
			Name: "fs",
			Type: "bar",
		},
		TilesetLocation:  "/var/tilr",
		TilesetImageType: "png",
	}

	tilesets["proxy"] = tileset.ProxyTileset{
		Tileset: tileset.Tileset{
			Name: "proxy",
			Type: "bar",
		},
		ProxyUrl:       "https://c.tile.openstreetmap.org",
		ProxyType:      "ZYX",
		ProxyExtension: "png",
	}

	// tilesets["foo"] = new(tileset.FilesystemTileset)
	// tilesets["foo"].TilesetLocation = "/var/tilr/foo"
	// tilesets["foo"].TilesetImageType = "png"

	// handle routes
	router := httprouter.New()
	router.GET("/", meta.MetaHandler)
	router.GET("/:tileset/:z/:y/:x", TileHandler)
	router.GET("/:tileset", InfoHandler)

	listenAddr := fmt.Sprintf("%s:%s", setting.Host, setting.Port)
	l, err := net.Listen("tcp4", listenAddr)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Listening on %s", listenAddr)
	}
	log.Fatal(http.Serve(l, router))
}
