package main

import (
	"github.com/gorilla/http"
	"log"
	"os"
)

type Address struct {
	street, city, state, zip string
}

func (a *Address) SingleLine() string {
	return a.street + " " + a.city + "," + a.state + " " + a.zip
}

type MDiMapGeocoder struct {
}

func (m *MDiMapGeocoder) Geocode(url string) {
	if _, err := http.Get(os.Stdout, url); err != nil {
		log.Fatalf("unable to fetch %q: %v", os.Args[1], err)
	}
}

func main() {

	var g = new(MDiMapGeocoder)

	var url = "http://geodata.md.gov/imap/rest/services/GeocodeServices/MD_CompositeLocatorWithZIPCodeCentroids/GeocodeServer/findAddressCandidates?Street=507+S+Pinehurst+Ave&City=Salisbury&State=Maryland&ZIP=21801&SingleLine=&outFields=&maxLocations=United+States&outSR=4326&searchExtent=&f=pjson"

	g.Geocode(url)

}
