package main

import (
	"echo/conf"
	"echo/model/cloud"
	"echo/model/example"
	"echo/router"
	"github.com/urfave/cli"

	"os"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	app := cli.NewApp()
	app.Name = "new echo project"
	var config string
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "config",
			Usage:       "set config path",
			Value:       "./conf/config.toml",
			Destination: &config,
		},
	}

	app.Action = func(c *cli.Context) error {
		conf.LoadConfig(config)
		example.LoadExampleSQL()
		cloud.LoadCloudSQL()
		router.Start()
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
