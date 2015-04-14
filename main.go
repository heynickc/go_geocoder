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

type Geocoder struct {
	httpUrl *url.URL
}

func NewGeocoder() *Geocoder {

	u := new(url.URL)

	u.Scheme = "http"
	u.Host = "geodata.md.gov"
	u.Path = "imap/rest/services/GeocodeServices/MD_CompositeLocatorWithZIPCodeCentroids/GeocodeServer/findAddressCandidates"

	v := url.Values{
		"Street":       []string{"507 S Pinehurst Ave"},
		"City":         []string{"Salisbury"},
		"State":        []string{"Maryland"},
		"ZIP":          []string{"21801"},
		"SingleLine":   []string{""},
		"outFields":    []string{""},
		"maxLocations": []string{"United States"},
		"outSR":        []string{"4326"},
		"searchExtent": []string{""},
		"f":            []string{"json"},
	}
	u.RawQuery = v.Encode()

	return &Geocoder{u}
}

func (m *Geocoder) Geocode(url string) {
	if _, err := http.Get(os.Stdout, url); err != nil {
		log.Fatalf("unable to fetch %q: %v", os.Args[1], err)
	}
}

func main() {

	gc := NewGeocoder()

	gc.Geocode(gc.httpUrl.String())
}
