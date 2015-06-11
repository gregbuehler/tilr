package cmd

import (
	"fmt"
	"net/http"

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

func runWeb(ctx *cli.Context) {
	if ctx.IsSet("port") {
		setting.Port = ctx.String("port")
	}

	// handle routes
	router := httprouter.New()
	router.GET("/", meta.MetaHandler)
	router.GET("/:tileset", tileset.TilesetHandler)
	router.GET("/:tileset/:z/:y/:x", tileset.TileHandler)

	listenAddr := fmt.Sprintf("%s:%s", setting.Host, setting.Port)
	log.Printf("Listening on %s", listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, router))
}
