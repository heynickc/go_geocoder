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
		if len(c.Args()) != 3 {
			log.Fatalf("usage: %s infile.ext outfile.ext\n")
		}
		inFileName, outFileName := c.Args()[1], c.Args()[2]
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
