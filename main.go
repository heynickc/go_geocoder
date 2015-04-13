package main

import (
	"log"
	"os"
	"github.com/gorilla/http"
)

func main() {

	var url = "http://geodata.md.gov/imap/rest/services/GeocodeServices/MD_CompositeLocatorWithZIPCodeCentroids/GeocodeServer/findAddressCandidates?Street=507+S+Pinehurst+Ave&City=Salisbury&State=Maryland&ZIP=21801&SingleLine=&outFields=&maxLocations=United+States&outSR=4326&searchExtent=&f=pjson"

	if _, err := http.Get(os.Stdout, url); err != nil {
		log.Fatalf("unable to fetch %q: %v", os.Args[1], err)
	}
}