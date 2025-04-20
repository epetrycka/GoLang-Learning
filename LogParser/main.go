package main

import (
	"LogParser/Logic"
	"log"
	"os"
	"github.com/urfave/cli/v2"
)

func main(){
	app := &cli.App{
		Name: "logparser",
		Usage: "Parse and filter logs",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "file",
				Aliases: []string{"f"},
				Usage: "Path to file with logs",
				Required: true,
			},
			&cli.StringFlag{
				Name: "logLevel",
				Aliases: []string{"l"},
				Usage: "Filter by log's level",
			},
			&cli.StringFlag{
				Name: "service",
				Aliases: []string{"s"},
				Usage: "Filter by log's service",
			},
		},
		Action: Logic.Parse,
	}

	err := app.Run(os.Args)
	if err != nil{
		log.Fatal(err)
	}
}