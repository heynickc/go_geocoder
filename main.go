package main

import (
	"github.com/codegangsta/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "go_geocoder"
	app.Usage = "adds map coordinates to input csv file\nOSX: ./go_geocoder <infile> <outfile>\nWindows: go_geocoder.exe <infile> <outfile>"
	app.Version = "0.1.0"
	app.Action = func(c *cli.Context) {
		if len(c.Args()) < 2 {
			log.Fatalf("usage: go_geocoder infile.csv outfile.csv\n")
		}
		inFileName, outFileName := c.Args()[0], c.Args()[1]
		if inFileName == outFileName {
			log.Fatalln("Can't overwrite a file with itself")
		}
		err := GeocodeFile(inFileName, outFileName)
		if err != nil {
			log.Fatalln("Couldn't geocode file: ", err)
		}
	}
	app.Run(os.Args)
}
