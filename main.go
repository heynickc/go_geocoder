package main

import (
	"github.com/codegangsta/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "go_mdimapgeocoder"
	app.Usage = "adds map coordinates to input csv file"
	app.Version = "0.1.0"
	app.Action = func(c *cli.Context) {
		if len(c.Args()) < 2 {
			log.Fatalf("usage: go_mdimapgeocoder.exe infile.csv outfile.csv\n")
		}
		inFileName, outFileName := c.Args()[0], c.Args()[1]
		if inFileName == outFileName {
			log.Fatalln("won't overwrite a file with itself")
		}
		err := GeocodeFile(inFileName, outFileName)
		if err != nil {
			log.Fatalln("No can do.. bub: ", err)
		}
	}
	app.Run(os.Args)
}
