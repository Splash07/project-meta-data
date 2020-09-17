package main

import (
	"os"

	"github.com/urfave/cli"
	"gitlab.ghn.vn/online/common/config"
	"gitlab.com/Splash07/project-meta-data/script"
	"gitlab.com/Splash07/project-meta-data/web"
)

var cfg = config.GetConfig()

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:    "project-meta-data",
			Aliases: []string{"s"},
			Action: func(c *cli.Context) error {
				return web.MasterData.Start()
			},
		},
		{
			Name:    "ward-set-status-1",
			Aliases: []string{"ward-set-status-1"},
			Action: func(c *cli.Context) error {
				return script.WardStatusRunner.Run()
			},
		},
	}

	if len(os.Args) == 1 {
		os.Args = append(os.Args, "s")
	}

	app.Run(os.Args)
}
