package cmd

import "github.com/codegangsta/cli"

var CmdDump = cli.Command{
	Name:        "dump",
	Usage:       "Dump Tilr cache",
	Description: `Dump compresses cached files for backup or archive`,
	Action:      runDump,
	Flags: []cli.Flag{
		cli.StringFlag{"config, c", "custom/conf/app.yaml", "Custom configuration file path", ""},
		cli.BoolFlag{"verbose, v", "show process details", ""},
	},
}

func runDump(ctx *cli.Context) {

}
