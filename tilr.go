package main

import (
	"os"
	"runtime"

	"github.com/codegangsta/cli"

	"github.com/gregbuehler/tilr/cmd"
	"github.com/gregbuehler/tilr/modules/setting"
)

const APP_VER = "0.0.2"

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	setting.AppVer = APP_VER
}

func main() {
	app := cli.NewApp()
	app.Name = "Tilr"
	app.Usage = "A Tile server/render"
	app.Version = APP_VER
	app.Commands = []cli.Command{
		cmd.CmdWeb,
		cmd.CmdDump,
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
}
