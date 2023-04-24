package cmd

import (
	"context"

	"github.com/LyricTian/gin-admin/v10/internal/bootstrap"
	"github.com/urfave/cli/v2"
)

func StartCmd() *cli.Command {
	return &cli.Command{
		Name:  "start",
		Usage: "Start server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "dir",
				Aliases: []string{"d"},
				Usage:   "Configurations directory",
				Value:   "configs",
			},
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Configuration file",
				Value:   "config.toml",
			},
			&cli.StringFlag{
				Name:    "static",
				Aliases: []string{"s"},
				Usage:   "Static files directory",
			},
			&cli.BoolFlag{
				Name:    "daemon",
				Aliases: []string{"d"},
				Usage:   "Run as a daemon",
			},
		},
		Action: func(c *cli.Context) error {
			err := bootstrap.Run(context.Background(), bootstrap.RunConfig{
				ConfigDir:  c.String("dir"),
				ConfigFile: c.String("config"),
				StaticDir:  c.String("static"),
				Daemon:     c.Bool("daemon"),
			})
			if err != nil {
				panic(err)
			}
			return nil
		},
	}
}
