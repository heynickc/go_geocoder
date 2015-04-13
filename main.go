package main

import (
	"github.com/gorilla/http"
	"log"
	"net/url"
	"os"
)

type Address struct {
	street, city, state, zip string
}

func (a *Address) SingleLine() string {
	return a.street + " " + a.city + "," + a.state + " " + a.zip
}

type MDiMapGeocoder struct {
	url.URL
}

func main() {

	// myAddres := Address{"507 S Pinehurst Ave", "Salisbury", "MD", "21801"}

	var url = "http://geodata.md.gov/imap/rest/services/GeocodeServices/MD_CompositeLocatorWithZIPCodeCentroids/GeocodeServer/findAddressCandidates?Street=507+S+Pinehurst+Ave&City=Salisbury&State=Maryland&ZIP=21801&SingleLine=&outFields=&maxLocations=United+States&outSR=4326&searchExtent=&f=pjson"

	if _, err := http.Get(os.Stdout, url); err != nil {
		log.Fatalf("unable to fetch %q: %v", os.Args[1], err)
	}
}
